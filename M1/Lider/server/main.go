package main

import (
	"context"
	"fmt"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
	"net"
	"time"
	"strconv"
	"math/rand"
	"strings"

)

type server struct {
	pb.UnimplementedEntradaMensajeServer
}

var flagListo = 0
var solicitudes [17]string
var respuestas [17]string

//Es como un len de python, cuenta strings no vacios de un array de tamaño 17.
func tamanio(lista [17]string) (cont int) {
	cont = 0
	for _,s := range lista{
		if s != ""{
			cont++
		}
	}
	return
}

//Es como un len de python, cuenta strings no vacios de un array de tamaño 17.
func tamanio2(lista [17]int) (cont int) {
	cont = 0
	for _,s := range lista{
		if s != 0{
			cont++
		}
	}
	return
}



//Limpia el array de solicitudes para no generar conflictos en la recepcion.
func VaciarSolicitudes(){
	for i:=0 ; i<17 ; i++{
		solicitudes[i] = ""
	}
}

//Limpia el array de respuestas para no generar conflictos de envio de respuestas
func VaciarRespuestas(){
	for i:=0 ; i<17 ; i++{
		respuestas[i] = ""
	}
}

//Toma los mjes de llegada y los procesa, mostrando por pantalla los mensajes acutales.
func Recepcion(mensaje string) (resmje string){
	temp := strings.Split(mensaje, " ")
	valor, _ := strconv.Atoi(temp[0])
	solicitudes[valor] = temp[1]
	fmt.Println("Mensajes de jugadores recibidos:")
	fmt.Println(solicitudes)
	if tamanio(solicitudes) == 16{
		flagListo = 0
	}
	resmje = "[*] Respuesta Recibida"
	return
}

//Retorno de respuesta a jugador que la solicitó
func Delivery(mensaje string) (respuesta string){
	temp := strings.Split(mensaje, " ")
	valor,_ := strconv.Atoi(temp[0])
	respuesta = respuestas[valor]
	return
}

//Verificador secundario para mejorar la coordinación de los nodos. 
func PendingRequest(mensaje string) (val int){
	val = 0
	temp := strings.Split(mensaje, " ")
	valor,_ := strconv.Atoi(temp[0])
	if solicitudes[valor] != ""{
		val = 1
	}else{
		val = 0
	}
	return
}

func (s *server ) Intercambio (ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	var res string
	fmt.Println("Se recibió el siguiente mensaje: "+ req.Body)
	//Mientras espera mensaje jugador debe escribir "<N jugador> Listo?"
	if strings.Contains(req.Body, "Listo?"){
		if (flagListo == 1) && (PendingRequest(req.Body) == 0) {
			res = Delivery(req.Body)
		}else {
			res = "[*] Processing..."
		}
		
	//Para pedir monto de pozo jugador debe hacer "<N jugador> POZO"
	}else if strings.Contains(req.Body, "POZO"){
	 	//ARMAR CONEXION A BASE DE DATOS DE LUCHO PARA OBTENER MONTO
	} else {
		res = Recepcion(req.Body)
		
	}
	return &pb.Mensaje{Body: res}, nil 
}

func initServer(){
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

func main() {
	status := [17]int{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	go initServer()

	//CONEXION A NAME NODE IP: M3:
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		print("PANIK")
		panic("No se puede conectar al servidor "+ err.Error())
	}

	serviceClient := pb.NewEntradaMensajeClient(conn)

	//RECEPCION DE SOLICITUDES POR PARTE DE LOS USUARIOS
	fmt.Println("[*] Esperando solicitudes...")
	for tamanio(solicitudes) < 16{
	}
	//SE ASUME QUE A ESTE PUNTO SE RECIBEN TODAS LAS SOLICITUDES.
	//PROCESAMIENTO JUEGO 1:
	PromptLider := -1
	for PromptLider != 0{
		fmt.Println("Todas las solicitudes recibidas!")
		fmt.Println("Seleccione una opción para continuar escribiendo un número:\n")
		fmt.Println("0) Dar inicio al juego 1 de Squid Game <コ:彡")
		fmt.Println("Aún no hay info sobre las jugadas, debe comenzar el juego primero")
		//fmt.Println("Ingrese un numero del 1 al 16 para consultar el historial de un jugador")
		fmt.Scanln(&PromptLider)
	}
	for i:=1 ; i<17 ; i++{
		respuestas[i] = "OK"
	}
	VaciarSolicitudes()
	flagListo = 1

	//Se autorizó la entrada a los jugadores, se procede a tomar las respuestas de la primera ronda.
	//Inicialiazción valores de ronda 1 de 4:
	juego1sumas := [17]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	
	rand.Seed(time.Now().UnixNano())
	ronda1valor := (rand.Intn(4)+6)

	fmt.Println("[*] Esperando respuestas de Ronda 1 - Juego 1...")

	for tamanio(solicitudes) < 16{
	}

	flagListo = 0
	VaciarRespuestas()
	var stringjugador string 
	for i:=1; i<17; i++{
		//Previo a procesar la respuesta se registra en el log de NameNode:
		stringjugador = strconv.Itoa(i)
		
		res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: "JUGA,"+stringjugador+",1,"+solicitudes[i],
		})
		if err != nil {
			panic("Error con la solicitud para ir registrando el historial de un jugador: "+err.Error())	
		}else{
			fmt.Println(res)
		}

		//Procesado de respuesta
		resR1,_ := strconv.Atoi(solicitudes[i])
		if resR1 == ronda1valor{
			fmt.Println("Jugador "+stringjugador+" Ha MUERTO")
			status[i] = 0
			respuestas[i] = "MUERTO"
		} else {
			juego1sumas[i] = resR1
			respuestas[i] = "OK"
		}
	}
	VaciarSolicitudes()
	flagListo = 1

	//RECEPCION DE SOLICITUDES POR PARTE DE LOS USUARIOS PARA RONDA 2
	//INICIALIZACION VARIABLES RONDA 2
	rand.Seed(time.Now().UnixNano())
	ronda2valor := (rand.Intn(4)+6)

	fmt.Println("[*] Esperando respuestas Ronda 2 Juego 1...")
	for tamanio(solicitudes) < tamanio2(status){
		
	}
	flagListo = 0
	VaciarRespuestas()

	for i:=1; i<17; i++{
		//Previo a procesar la respuesta se registra en el log de NameNode:
		stringjugador = strconv.Itoa(i)
		
		res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: "JUGA,"+stringjugador+",1,"+solicitudes[i],
		})
		if err != nil {
			panic("Error con la solicitud para ir registrando el historial de un jugador: "+err.Error())	
		}else{
			fmt.Println(res)
		}

		//Procesado de respuesta
		resR2,_ := strconv.Atoi(solicitudes[i])
		if resR2 == ronda2valor{
			fmt.Println("Jugador "+stringjugador+" Ha MUERTO")
			status[i] = 0
			respuestas[i] = "MUERTO"
		} else {
			juego1sumas[i] += resR2
			respuestas[i] = "OK"
		}
	}
	VaciarSolicitudes()
	flagListo = 1

	//RECEPCION DE SOLICITUDES POR PARTE DE LOS USUARIOS PARA RONDA 3
	//INICIALIZACION VARIABLES RONDA 3
	rand.Seed(time.Now().UnixNano())
	ronda3valor := (rand.Intn(4)+6)

	fmt.Println("[*] Esperando respuestas Ronda 2 Juego 1...")
	for tamanio(solicitudes) < 16{
		
	}
	flagListo = 0
	VaciarRespuestas()

	for i:=1; i<17; i++{
		//Previo a procesar la respuesta se registra en el log de NameNode:
		stringjugador := strconv.Itoa(i)
		
		res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: "JUGA,"+stringjugador+",1,"+solicitudes[i],
		})
		if err != nil {
			panic("Error con la solicitud para ir registrando el historial de un jugador: "+ err.Error())	
		}else{
			fmt.Println(res)
		}

		//Procesado de respuesta
		resR3,_ := strconv.Atoi(solicitudes[i])
		if resR3 == ronda3valor{
			fmt.Println("Jugador "+stringjugador+" Ha MUERTO")
			status[i] = 0
			respuestas[i] = "MUERTO"
		} else {
			juego1sumas[i] += resR3
			respuestas[i] = "OK"
		}
	}



	
	


}