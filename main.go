package main

import (
	"context"
	"log"
	"mypackage/invoicer"
	"net"

	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverReg := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverReg, service)

	err = serverReg.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
