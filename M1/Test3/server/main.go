package main

import (
	"context"
	"fmt"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedEntradaMensajeServer
}

var cont = 0
var respuestas [16]string

func Recepcion(mensaje string){
	respuestas[cont] = mensaje
	cont++
	fmt.Println(respuestas)
}

func (s *server ) Intercambio (ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	fmt.Println("Se recibió el siguiente mensaje: "+ req.Body)
	Recepcion(req.Body)
	return &pb.Mensaje{Body: "Mensaje recibido desde servidor"}, nil 
}

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("No se puede crear la conexión tcp: "+ err.Error())
	}

	fmt.Println(listener)
	serv := grpc.NewServer()
	pb.RegisterEntradaMensajeServer(serv, &server{})
	if err = serv.Serve(listener); err != nil {
		panic("No se ha podido inicializar el servidor: "+ err.Error())
	}

}