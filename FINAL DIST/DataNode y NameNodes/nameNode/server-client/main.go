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
	"math/rand"
	"time"
)

type server struct {
	pb.UnimplementedEntradaMensajeServer
}

//MODIFICAR AQUI LOS VALORES DE LAS MAQUINAS
//IP M1
var ipDataNode1 = "localhost:50053"
//IP M2
var ipDataNode2 = "localhost:50054"
//IP M3
var ipDataNode3 = "localhost:50055"

//Creacion de archivo data
var (
	newfile *os.File
	err	error
)

//var res1 string

//Extremo entre Lider y NameNode (para que namenode reciba solicitudes de Lider) [POR DEFECTO PORT: 50052]
func (s *server ) Intercambio (ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	fmt.Println("NameNode recibió el siguiente mensaje: "+ req.Body)
	var res1 string
	conn1, err1 := grpc.Dial(ipDataNode1, grpc.WithInsecure())
	if err1 != nil {
		panic("No se puede conectar al Data Node 1 "+ err.Error())
	}
	serviceClient1 := pb.NewEntradaMensajeClient(conn1)
	
	conn2, err2 := grpc.Dial(ipDataNode2, grpc.WithInsecure())
	if err2 != nil {
		panic("No se puede conectar al Data Node 2 "+ err.Error())
	}
	serviceClient2 := pb.NewEntradaMensajeClient(conn2)

	conn3, err3 := grpc.Dial(ipDataNode3, grpc.WithInsecure())
	if err3 != nil {
		panic("No se puede conectar al Data Node 3 "+ err.Error())
	}
	serviceClient3 := pb.NewEntradaMensajeClient(conn3)

	// FORMATO DE SOLICITUD DESDE LIDER: "DATA,Jugador_<numero_jugador>,Ronda_<numero_ronda>"
	// FORMATO DE JUGADA DESDE LIDER: "JUGA,<numero_jugador>,<numero_ronda>,<jugada>"
	request := strings.Split(string(req.Body),",")
	fmt.Println(request)
	if len(request) == 3 {
		file, err := os.Open("data/data.txt")
		if err != nil{
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var text []string
		for scanner.Scan() {
			text = strings.Split(scanner.Text(), " ")
			if (text[0] == request[1]) && (text[1] == request[2]){
				jugador := strings.Split(text[0],"_")[1]
				solicitud := "ARCHIVO,jugador_"+jugador+"__ronda_"+strings.Split(request[2],"_")[1]+".txt"
				if ipDataNode1 == text[2]{
					res, err := serviceClient1.Intercambio(context.Background(), &pb.Mensaje{
						Body: solicitud,
					})

					if err != nil {
						panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
					}
					res1 = res.Body
				} else if ipDataNode2 == text[2]{
					res, err := serviceClient2.Intercambio(context.Background(), &pb.Mensaje{
						Body: solicitud,
					})

					if err != nil {
						panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
					}
					res1 = res.Body

				}else{
					res, err := serviceClient3.Intercambio(context.Background(), &pb.Mensaje{
						Body: solicitud,
					})

					if err != nil {
						panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
					}
					res1 = res.Body
				}
			}
		} 
		if res1 == ""{
			res1 = "No se tiene data del jugador en la "+request[2]
		} 
	} else {
		file, err := os.Open("data/data.txt")
		if err != nil{
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var text []string
		flag := 0
		for scanner.Scan() {
			text = strings.Split(scanner.Text(), " ")
			if (text[0] == "Jugador_"+request[1]) && (text[1] == "Ronda_"+request[2]){
				flag = 1
				solicitud:="AGREGA,jugador_"+request[1]+"__ronda_"+request[2]+".txt,"+request[3]
				if ipDataNode1 == text[2]{
					res, err := serviceClient1.Intercambio(context.Background(), &pb.Mensaje{
						Body: solicitud,
					})

					if err != nil {
						panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
					}
					res1 = res.Body
				} else if ipDataNode2 == text[2]{
					res, err := serviceClient2.Intercambio(context.Background(), &pb.Mensaje{
						Body: solicitud,
					})

					if err != nil {
						panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
					}
					res1 = res.Body
				}else{
					res, err := serviceClient3.Intercambio(context.Background(), &pb.Mensaje{
						Body: solicitud,
					})

					if err != nil {
						panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
					}
					res1 = res.Body
				}
			}
		}
		file.Close()
		if flag == 0 {
			rand.Seed(time.Now().UnixNano())
			nodo := (rand.Intn(3)+1)
			solicitud := "AGREGA,jugador_"+request[1]+"__ronda_"+request[2]+".txt,"+request[3]
			if nodo == 1{
				res, err := serviceClient1.Intercambio(context.Background(), &pb.Mensaje{
					Body: solicitud,
				})

				if err != nil {
					panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
				}
				res1 = res.Body
				f, err := os.OpenFile("data/data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
    				panic(err)
				}

				defer f.Close()

				if _, err = f.WriteString("Jugador_"+request[1]+" Ronda_"+request[2]+" "+ipDataNode1+"\n"); err != nil {
    				panic(err)
				}
			} else if nodo == 2{
				res, err := serviceClient2.Intercambio(context.Background(), &pb.Mensaje{
					Body: solicitud,
				})

				if err != nil {
					panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
				}
				res1 = res.Body
				f, err := os.OpenFile("data/data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
    				panic(err)
				}

				defer f.Close()

				if _, err = f.WriteString("Jugador_"+request[1]+" Ronda_"+request[2]+" "+ipDataNode2+"\n"); err != nil {
    				panic(err)
				}

			}else{
				res, err := serviceClient3.Intercambio(context.Background(), &pb.Mensaje{
					Body: solicitud,
				})

				if err != nil {
					panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
				}
				res1 = res.Body
				f, err := os.OpenFile("data/data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
    				panic(err)
				}

				defer f.Close()

				if _, err = f.WriteString("Jugador_"+request[1]+" Ronda_"+request[2]+" "+ipDataNode3+"\n"); err != nil {
    				panic(err)
				}
			}
		}
	}
	return &pb.Mensaje{Body: res1}, nil 
}

func main() {
	/*
	newFile, err = os.Create("data/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
	*/
	
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic("No se puede crear la conexión tcp: "+ err.Error())
	}
	
	serv := grpc.NewServer()
	pb.RegisterEntradaMensajeServer(serv, &server{})
	if err = serv.Serve(listener); err != nil {
		panic("No se ha podido inicializar el servidor gRPC: "+ err.Error())
	}
}