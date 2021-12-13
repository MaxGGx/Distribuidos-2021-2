package main

import (
	"context"
	"fmt"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
	"net"
	"strings"
	"strconv"
	"math/rand"
	"time"
)
//Variables constantes
var ipFulcrum1 = "localhost:50002"
var ipFulcrum2 = "localhost:50003"
var ipFulcrum3 = "localhost:50004"


type server struct {
	pb.UnimplementedEntradaMensajeServer
}

func (s *server ) Intercambio (ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	ans := ""
	fmt.Println("Broker recibió el siguiente mensaje: "+ req.Body)
	if(strings.Split(req.Body, " ")[0] == "GetNumberRebelds"){
		ans = getInfo(req.Body)
	} else {
		ans = obtainIP(req.Body)
	}
	return &pb.Mensaje{Body: ans}, nil 
}

//Funcion ejecutada por gRPC para enviar el mensaje
func Solicitud(serviceClient pb.EntradaMensajeClient, msg string) string{
	res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: msg,
	})
	if err != nil {
		panic("Mensaje no pudo ser creado ni enviado: " + err.Error())
	}
	//fmt.Println(res.Body)
	return res.Body
}

//Funcion que toma la ip a la que se desea enviar, se conecta y realiza el envio del mensaje. Retorna la respuesta
func enviarMsg(ip string, msg string)(answer string){
	conn, err := grpc.Dial(ip, grpc.WithInsecure())

	if err != nil {
		panic("No se puede conectar al servidor "+ err.Error())
	}

	serviceClient := pb.NewEntradaMensajeClient(conn)

	answer = Solicitud(serviceClient, msg)
	
	defer conn.Close()
	
	return
}

func randIP()(ip string){
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	valor := (r1.Intn(2)+1)
	if(valor == 1){
		ip = ipFulcrum1
	} else if (valor == 2){
		ip = ipFulcrum2
	} else {
		ip = ipFulcrum3
	}
	return
}

//Funcion que consulta a los fulcrum para saber si puede redirigir aleatoriamente
func obtainIP(planeta string)(ipfulcrum string){
	//Comando CLK <nombre_planeta>
	//Se consulta a cada fulcrum para comparar luego los relojes.
	consulta := "CLK "+planeta
	data1 := enviarMsg(ipFulcrum1, consulta)
	data2 := enviarMsg(ipFulcrum2, consulta)
	data3 := enviarMsg(ipFulcrum3, consulta)
	if ((data1 == "Err404")&&(data2 == "Err404")&&(data3 == "Err404")){
		//Ninguno lo tiene, puedo retornar uno al azar.
		ipfulcrum = randIP()

	} else if ((data1 != "Err404") && (data2 == "Err404") && (data3 == "Err404")){
		//Solo el fulcrum 1 lo tiene, retorno ese.
		ipfulcrum = ipFulcrum1

	} else if ((data1 == "Err404") && (data2 != "Err404") && (data3 == "Err404")){
		//Solo el fulcrum 2 lo tiene, retorno ese.
		ipfulcrum = ipFulcrum2

	} else if ((data1 == "Err404") && (data2 == "Err404") && (data3 != "Err404")){
		//Solo el fulcrum 3 lo tiene, retorno ese.
		ipfulcrum = ipFulcrum3

	} else if ((data1 == data2) && (data1 == data3)){
		//Son todos iguales, puedo retornar uno al azar.
		ipfulcrum = randIP()
	} else if ((data1 != "Err404") && (data2 != "Err404") && (data3 == "Err404")){
		//Solo data1 y data2 lo tienen (asumiendo que aun a data 3 no le propagan la info)
		//Procedo a comparar los valores de cada uno de los relojes.
		flag1:=0
		datas1 := strings.Split(data1, " ")
		datas2 := strings.Split(data2, " ")
		datap1 := []int {-1,-1,-1}
		datap2 := []int {-1,-1,-1}
		for i:= range datas1 {
			//Se convierten variables a int para ser procesadas.
			datap1[i],_ = strconv.Atoi(datas1[i])
			datap2[i],_ = strconv.Atoi(datas2[i])
		}
		for i:= range datap1{
			if((datap1[i] > datap2[i])){
				ipfulcrum = ipFulcrum1
				flag1 = 1
			} else if ((datap1[i] < datap2[i])) {
				ipfulcrum = ipFulcrum2
				flag1 = 1
			}
		}
		if (flag1 == 0){
			ipfulcrum = ipFulcrum1
		}
	} else if ((data2 != "Err404") && (data3 != "Err404") && (data1 == "Err404")){
		//Solo data2 y data3 lo tienen (asumiendo que aun a data 1 no le propagan la info)
		//Procedo a comparar los valores de cada uno de los relojes.
		flag1:=0
		datas2 := strings.Split(data2, " ")
		datas3 := strings.Split(data3, " ")
		datap2 := []int {-1,-1,-1}
		datap3 := []int {-1,-1,-1}
		for i:= range datas2 {
			//Se convierten variables a int para ser procesadas.
			datap2[i],_ = strconv.Atoi(datas2[i])
			datap3[i],_ = strconv.Atoi(datas3[i])
		}
		for i:= range datap2{
			if((datap2[i] > datap3[i])){
				ipfulcrum = ipFulcrum2
				flag1 = 1
			} else if ((datap2[i] < datap3[i])) {
				ipfulcrum = ipFulcrum3
				flag1 = 1
			}
		}
		if (flag1 == 0){
			ipfulcrum = ipFulcrum2
		}
	} else if ((data1 != "Err404") && (data3 != "Err404") && (data2 == "Err404")){
		//Solo data1 y data3 lo tienen (asumiendo que aun a data 2 no le propagan la info)
		//Procedo a comparar los valores de cada uno de los relojes.
		flag1:=0
		datas1 := strings.Split(data1, " ")
		datas3 := strings.Split(data3, " ")
		datap1 := []int {-1,-1,-1}
		datap3 := []int {-1,-1,-1}
		for i:= range datas1 {
			//Se convierten variables a int para ser procesadas.
			datap1[i],_ = strconv.Atoi(datas1[i])
			datap3[i],_ = strconv.Atoi(datas3[i])
		}
		for i:= range datap1{
			if((datap1[i] > datap3[i])){
				ipfulcrum = ipFulcrum1
				flag1 = 1
			} else if ((datap1[i] < datap3[i])) {
				ipfulcrum = ipFulcrum3
				flag1 = 1
			}
		}
		if (flag1 == 0){
			ipfulcrum = ipFulcrum3
		}
	} else {
		//Procedo a comparar los valores de cada uno de los relojes.
		datas1 := strings.Split(data1, " ")
		datas2 := strings.Split(data2, " ")
		datas3 := strings.Split(data3, " ")
		datap1 := []int {-1,-1,-1}
		datap2 := []int {-1,-1,-1}
		datap3 := []int {-1,-1,-1}
		for i:= range datas1 {
			//Se convierten variables a int para ser procesadas.
			datap1[i],_ = strconv.Atoi(datas1[i])
			datap2[i],_ = strconv.Atoi(datas2[i])
			datap3[i],_ = strconv.Atoi(datas3[i])
		}
		for i:= range datap1{
			if ((datap1[i] > datap2[i]) && (datap1[i] > datap3[i])){
				//fulcrum 1 tiene el más actualizado, retorno ese
				ipfulcrum = ipFulcrum1

			} else if ((datap2[i] > datap1[i]) && (datap2[i] > datap3[i])){
				//fulcrum 2 tiene el más actualizado, retorno ese
				ipfulcrum = ipFulcrum2

			} else if ((datap3[i] > datap1[i]) && (datap3[i] > datap2[i])){
				//fulcrum 3 tiene el más actualizado, retorno ese
				ipfulcrum = ipFulcrum3

			}
		}
	}
	return
}

//Manejara los comandos provenientes de Leia (GetNumberRebelds)
func getInfo(command string)(respuesta string){
	comando := strings.Split(command, " ")
	//Obtengo el planeta, de ahi puedo utilizarlo en getIP, para saber el más actualizado, cosa de retornar ese.
	ip := obtainIP(comando[1])
	respuesta = enviarMsg(ip, command)
	return
}

func main() {
	//fmt.Println(enviarMsg(ipFulcrum1, "Hola"))
	//Solo inicializo server, Funciones se encargan del resto 
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