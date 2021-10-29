package main

import (
	"context"
	"fmt"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("No se puede conectar al servidor "+ err.Error())
	}

	serviceClient := pb.NewEntradaMensajeClient(conn)

	res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: "Mensaje de prueba desde cliente",
	})

	if err != nil {
		panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
	}

	fmt.Println(res.Body)
}