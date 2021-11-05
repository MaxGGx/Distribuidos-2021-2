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

func (s *server ) Intercambio (ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	fmt.Println("Se recibió el siguiente mensaje: "+ req.Body)
	return &pb.Mensaje{Body: "Mensaje recibido desde servidor"}, nil 
}

func crearServer(server *grpc.Server, direccion string){
	listener, err := net.Listen("tcp", direccion)
	if err != nil {
		panic("No se puede crear la conexión tcp: "+ err.Error())
	}
	if err = server.Serve(listener); err != nil {
		panic("No se ha podido inicializar el servidor: "+ err.Error())
	}
	
}

func main() {
	s := &server{}
	serv1 := grpc.NewServer()
	//serv2 := grpc.NewServer()
	//serv3 := grpc.NewServer()
	//serv4 := grpc.NewServer()
	//serv5 := grpc.NewServer()
	//serv6 := grpc.NewServer()
	//serv7 := grpc.NewServer()
	//serv8 := grpc.NewServer()
	//serv9 := grpc.NewServer()
	//serv10 := grpc.NewServer()
	//serv11 := grpc.NewServer()
	//serv12 := grpc.NewServer()
	//serv13 := grpc.NewServer()
	//serv14 := grpc.NewServer()
	//serv15 := grpc.NewServer()
	//serv16 := grpc.NewServer()

	pb.RegisterEntradaMensajeServer(serv1, s)
	//pb.RegisterEntradaMensajeServer(serv2, s)
	//pb.RegisterEntradaMensajeServer(serv3, s)
	//pb.RegisterEntradaMensajeServer(serv4, s)
	//pb.RegisterEntradaMensajeServer(serv5, s)
	//pb.RegisterEntradaMensajeServer(serv6, s)
	//pb.RegisterEntradaMensajeServer(serv7, s)
	//pb.RegisterEntradaMensajeServer(serv8, s)
	//pb.RegisterEntradaMensajeServer(serv9, s)
	//pb.RegisterEntradaMensajeServer(serv10, s)
	//pb.RegisterEntradaMensajeServer(serv11, s)
	//pb.RegisterEntradaMensajeServer(serv12, s)
	//pb.RegisterEntradaMensajeServer(serv13, s)
	//pb.RegisterEntradaMensajeServer(serv14, s)
	//pb.RegisterEntradaMensajeServer(serv15, s)
	//pb.RegisterEntradaMensajeServer(serv16, s)

	go crearServer(serv1,":50051")
	//go crearServer(serv2,":50052")
	//crearServer(serv3,":50053")
	//go crearServer(serv4,":50054")
	//go crearServer(serv5,":50055")
	//go crearServer(serv6,":50056")
	//go crearServer(serv7,":50057")
	//go crearServer(serv8,":50058")
	//go crearServer(serv9,":50059")
	//go crearServer(serv10,":50060")
	//go crearServer(serv11,":50061")
	//go crearServer(serv12,":50062")
	//go crearServer(serv13,":50063")
	//go crearServer(serv14,":50064")
	//go crearServer(serv15,":50065")
	//go crearServer(serv16,":50066")
	/*
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("No se puede crear la conexión tcp: "+ err.Error())
	}
	if err = serv1.Serve(listener); err != nil {
		panic("No se ha podido inicializar el servidor: "+ err.Error())
	}
	*/
}