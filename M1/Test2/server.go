package main

import (
	"log"
	"net"

	"github.com/MaxGGx/Distribuidos-2021-2/tree/main/M1/Test2/chat"
	"google.golang.org/grpc"
	
)

//Comando para inicializar el go.mod y go.sum, se le indica al comando donde estará alojado en el repo la carpeta con los archivos (o algo asi)
//De esa forma los crea para que cuando se haga el commit, los vaya a ver allí también
//go mod init github.com/MaxGGx/Distribuidos-2021-2/tree/main/M1/Test2

func main() {
	lis, err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatalf("No se puede escuchar en el puerto 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo en habilitar servidor gRPC en el puerto 9000: %v", err)
	}

	
}
