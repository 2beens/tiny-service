package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/2beens/tiny-service/internal"
	tseProto "github.com/2beens/tiny-stock-exchange-proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting tiny stock exchange grpc server ...")

	flag.Parse()
	host := flag.String("host", "localhost", "tiny stock exchange server host")
	port := flag.Int("port", 9002, "tiny stock exchange server port")

	tcpListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("will be listening on tcp %s:%d", *host, *port)

	grpcServer := grpc.NewServer()
	tinyStockExchange := internal.NewTinyStockExchange()
	tseProto.RegisterTinyStockExchangeServer(grpcServer, tinyStockExchange)
	grpcServer.Serve(tcpListener)
}
