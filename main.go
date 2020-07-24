package main

import (
	proto "github.com/kiranparajuli589/currency-converter/protos/currency"
	"github.com/kiranparajuli589/currency-converter/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main()  {
	gRPCServer := grpc.NewServer()
	currency := server.NewCurrency()

	proto.RegisterCurrencyServer(gRPCServer, currency)

	reflection.Register(gRPCServer)

	listen, err := net.Listen("tcp", ":9200")
	if err != nil {
		panic(err)
	}
	if err = gRPCServer.Serve(listen); err != nil {
		panic(err)
	}
}
