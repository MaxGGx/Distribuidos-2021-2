package main

import (
	"context"
	"fmt"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
	"reflect"
	"strconv"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("No se puede conectar al servidor "+ err.Error())
	}

	serviceClient := pb.NewEntradaMensajeClient(conn)
	fmt.Println(reflect.TypeOf(serviceClient))

	for i:=1 ; i<17;i++{
		valor := strconv.Itoa(i)
		res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: valor+" Listo?",
	})
	if err != nil {
		panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
	}
	fmt.Println(res.Body)
	}
	

	

}