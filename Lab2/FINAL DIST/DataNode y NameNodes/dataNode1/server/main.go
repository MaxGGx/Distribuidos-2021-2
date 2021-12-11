package main

import (
	"context"
	"fmt"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
	"net"
	"strings"
	"log"
	"os"
	"bufio"
	"errors"
)

type server struct {
	pb.UnimplementedEntradaMensajeServer
}

var res string

func (s *server ) Intercambio (ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	fmt.Println("NameNode recibió el siguiente mensaje: "+ req.Body)
	request := strings.Split(string(req.Body),",")
	fmt.Println(request)
	if request[0] == "ARCHIVO"{
		if _,err := os.Stat("../data/"+request[1]); err == nil {
			file, err := os.Open("../data/"+request[1])

			if err != nil{
				log.Fatalf("Fallo en abrir archivo")
			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string

			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()
			res = strings.Join(text,"\n")
		} else if errors.Is(err, os.ErrNotExist) {
			res = "Jugador no tiene datos en DataNode"
		}
	} else {
		f, err := os.OpenFile("../data/"+request[1], os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
    		panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(request[2]+"\n"); err != nil {
    		panic(err)
		}
		res = "Jugada agregada exitosamente"
	}

	return &pb.Mensaje{Body: res}, nil 
}

func main() {
	//M1
	listener, err := net.Listen("tcp", ":50053")

	if err != nil {
		panic("No se puede crear la conexión tcp: "+ err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterEntradaMensajeServer(serv, &server{})
	if err = serv.Serve(listener); err != nil {
		panic("No se ha podido inicializar el servidor: "+ err.Error())
	}
}