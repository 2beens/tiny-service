package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/2beens/tiny-service/internal"
	tseProto "github.com/2beens/tiny-stock-exchange-proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func main() {
	log.Println("starting tiny stock exchange grpc server ...")
	flag.Parse()
	host := flag.String("host", "localhost", "tiny stock exchange server host")
	port := flag.Int("port", 9002, "tiny stock exchange server port")
	mdbHost := flag.String("mdbhost", "localhost", "mongo db host")

	chOsInterrupt := make(chan os.Signal, 1)
	signal.Notify(chOsInterrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	ctx, cancel := context.WithCancel(context.Background())

	// connect to mongo db
	log.Println("creating mongodb client ...")
	mdbConnStr := fmt.Sprintf("mongodb://root:root@%s:27017/?maxPoolSize=20&w=majority", *mdbHost)
	opts := options.Client()
	opts.SetConnectTimeout(5 * time.Second)
	opts.ApplyURI(mdbConnStr)

	mongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %s", err)
	}

	log.Printf("ping mongodb client [%s] ...", mdbConnStr)
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer timeoutCancel()
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("failed to ping mongodb: %s", err)
	}
	log.Println("mongodb client ping ok")

	grpcServer := grpc.NewServer()
	tinyStockExchange := internal.NewTinyStockExchange(mongoClient)
	go func() {
		tseProto.RegisterTinyStockExchangeServer(grpcServer, tinyStockExchange)

		tcpListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Printf("listening on tcp %s:%d", *host, *port)
		grpcServer.Serve(tcpListener)
	}()

	receivedSig := <-chOsInterrupt
	log.Printf("signal [%s] received, killing everything ...", receivedSig)
	cancel()

	grpcServer.GracefulStop()

	timeoutCtx, timeoutCancel = context.WithTimeout(context.Background(), time.Second*5)
	defer timeoutCancel()
	if err = mongoClient.Disconnect(timeoutCtx); err != nil {
		log.Fatalf("failed to disconnect from mongodb: %s", err)
	}
}
