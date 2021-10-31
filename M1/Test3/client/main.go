package main

import (
	"context"
	"fmt"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
)

func msg(direccion string){
	//fmt.Printf("%s\n","localhost:"+direccion)
	conn, err := grpc.Dial("localhost:"+direccion, grpc.WithInsecure())

	if err != nil {
		panic("No se puede conectar al servidor "+ err.Error())
	}

	serviceClient := pb.NewEntradaMensajeClient(conn)

	res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: "POZO",
	})

	if err != nil {
		panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
	}

	fmt.Println(res.Body)
}

func main() {
	//var srvrs = []string{"50051","50052","50053","50054","50055","50056","50057","50058","50059","50060","50061","50062","50063","50064","50065","50066"}
	msg("50051")
	msg("50052")
	msg("50060")
	//for i := 0; i < 15; i++ {
	//	msg(srvrs[i])
	//}
	/*
	conn, err := grpc.Dial("localhost"+":50051", grpc.WithInsecure())

	if err != nil {
		panic("No se puede conectar al servidor "+ err.Error())
	}

	serviceClient := pb.NewEntradaMensajeClient(conn)

	res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: "Enviando prueba...",
	})

	if err != nil {
		panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
	}

	fmt.Println(res.Body)
	*/
}