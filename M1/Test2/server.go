package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatalf("No se puede escuchar en el puerto 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo en habilitar servidor gRPC en el puerto 9000: %v", err)
	}

	
}