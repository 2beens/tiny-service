package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/2beens/tiny-service/internal"
	tseProto "github.com/2beens/tiny-stock-exchange-proto"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func init() {
	// logger setup
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

func main() {
	log.Println("starting tiny stock exchange grpc server ...")
	host := flag.String("host", "localhost", "tiny stock exchange server host")
	port := flag.String("port", "9002", "tiny stock exchange server port")
	instanceName := flag.String("name", "anon-instance", "name of this tiny api instance")
	mdbHost := flag.String("mdbhost", "localhost", "mongo db host")
	mdbPort := flag.String("mdbport", "27017", "mongo db port")
	tseDBName := flag.String("tsedb", "tiny-stock-exchange", "mongo db tiny stock exchange db name")
	flag.Parse()

	if envVarHost := os.Getenv("TINY_SERVICE_HOST"); envVarHost != "" {
		*host = envVarHost
		log.Debugf("host [%s] present in env. var, will use it instead", envVarHost)
	}
	if envVarPort := os.Getenv("TINY_SERVICE_PORT"); envVarPort != "" {
		*port = envVarPort
		log.Debugf("port [%s] present in env. var, will use it instead", envVarPort)
	}
	if envVarMdbHost := os.Getenv("TINY_SERVICE_MONGO_HOST"); envVarMdbHost != "" {
		*mdbHost = envVarMdbHost
		log.Debugf("mongodb host [%s] present in env. var, will use it instead", envVarMdbHost)
	}
	if envVarMdbPort := os.Getenv("TINY_SERVICE_MONGO_PORT"); envVarMdbPort != "" {
		*mdbPort = envVarMdbPort
		log.Debugf("port [%s] present in env. var, will use it instead", envVarMdbPort)
	}
	if envVarInstanceName := os.Getenv("TINY_SERVICE_INSTANCE_NAME"); envVarInstanceName != "" {
		*instanceName = envVarInstanceName
		log.Debugf("instance name [%s] present in env. var, will use it instead", envVarInstanceName)
	}

	log.Debugf("instance %s: tiny service starting ...", *instanceName)

	chOsInterrupt := make(chan os.Signal, 1)
	signal.Notify(chOsInterrupt, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	// connect to mongo db
	log.Println("creating mongodb client ...")
	mdbConnStr := fmt.Sprintf("mongodb://root:root@%s:%s/?maxPoolSize=20&w=majority", *mdbHost, *mdbPort)
	opts := options.Client()
	opts.SetConnectTimeout(5 * time.Second)
	opts.ApplyURI(mdbConnStr)

	mongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %s", err)
	}

	log.Printf("ping mongodb client [%s] ...", mdbConnStr)
	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, time.Second*5)
	defer timeoutCancel()
	if err := mongoClient.Ping(timeoutCtx, readpref.Primary()); err != nil {
		log.Fatalf("failed to ping mongodb: %s", err)
	}
	log.Println("mongodb client ping ok")

	grpcServer := grpc.NewServer()
	tinyStockExchange, err := internal.NewTinyStockExchange(*instanceName, *tseDBName, mongoClient)
	if err != nil {
		log.Fatalf("create tiny stock exchange: %s", err)
	}

	go func() {
		tseProto.RegisterTinyStockExchangeServer(grpcServer, tinyStockExchange)

		tcpListener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", *host, *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Debugf("listening on tcp %s:%s", *host, *port)
		if err := grpcServer.Serve(tcpListener); err != nil {
			log.Errorf("grpc server serve: %s", err)
		}
	}()

	receivedSig := <-chOsInterrupt
	log.Warnf("signal [%s] received, killing everything ...", receivedSig)
	cancel()

	grpcServer.GracefulStop()

	timeoutCtx, timeoutCancel = context.WithTimeout(context.Background(), time.Second*5)
	defer timeoutCancel()
	if err = mongoClient.Disconnect(timeoutCtx); err != nil {
		log.Fatalf("failed to disconnect from mongodb: %s", err)
	}
}
