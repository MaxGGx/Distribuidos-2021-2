package main

import (
	"context"
	"fmt"
	"net"

	pb "../proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEntradaMensajeServer
}

func (s *server) Intercambio(ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	fmt.Println("Se recibió el siguiente mensaje: " + req.Body)
	return &pb.Mensaje{Body: "Mensaje recibido desde servidor"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("No se puede crear la conexión tcp: " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterEntradaMensajeServer(serv, &server{})
	if err = serv.Serve(listener); err != nil {
		panic("No se ha podido inicializar el servidor: " + err.Error())
	}
}
