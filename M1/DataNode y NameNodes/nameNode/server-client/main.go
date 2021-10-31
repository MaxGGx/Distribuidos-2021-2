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

//Extremo entre Lider y NameNode (para que namenode reciba solicitudes de Lider) [POR DEFECTO PORT: 50051]
func (s *server ) Intercambio (ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	fmt.Println("NameNode recibió el siguiente mensaje: "+ req.Body)

	return &pb.Mensaje{Body: "Mensaje recibido desde servidor"}, nil 
}

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("No se puede crear la conexión tcp: "+ err.Error())
	}
	serv := grpc.NewServer()
	pb.RegisterEntradaMensajeServer(serv, &server{})
	if err = serv.Serve(listener); err != nil {
		panic("No se ha podido inicializar el servidor: "+ err.Error())
	}
}