package main

import ( 
	"fmt"
	"context"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
	"strconv"
	"strings"
	"bufio"
	"os"
)

//Struct para poder hacer un objeto de planeta dependiendo si es que esta creado.
type planeta struct {
	nombre string //Nombre del planeta manejado (registro)
	relojx int //Dimension X del reloj de vector
	relojy int //Dimension Y del reloj de vector
	relojz int //Dimension Z del reloj de vector
	lastfulcrum string //ip del ultimo fulcrum consultado para este planeta
}

//Lista de structs que almacenará de manera eficiente los planetas.
var planetas []planeta
var direccionBroker = "dist33:50051"
var direccionFulcrum= ""

//Constructor para el planeta, cosa de poder almacenar en memoria la info de los planetas manejados por la consola del informante.
func Cplaneta(name string, x int, y int, z int, ip string)(planet planeta){
	planet = planeta{
		nombre: name,
		relojx: x,
		relojy: y,
		relojz: z,
		lastfulcrum: ip,
	}
	return
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

//Procesa los comandos del usuario (Consulta a broker, luego a Fulcrum).
func processMsg(command string){
	//Comando = ["AddCity planeta0 ciudad0 10"]
	var comando = strings.Split(command, " ")
	
	//Se recibe la ip para el fulcrum
	respuesta := enviarMsg(direccionBroker, command)
	fmt.Println("[*] Ip recibida desde el Broker:\n")
	fmt.Println(respuesta)
	direccionFulcrum = respuesta

	//Se consulta al Fulcrum
	fmt.Println("[*] Ejecutando consulta al servidor fulcrum...")
	respuesta = enviarMsg(respuesta, command)
	fmt.Println("[*] Respuesta recibida!, datos:\n")
	fmt.Println(respuesta)
	
	//Se analiza si no hay error
	data := strings.Split(respuesta, " ")
	if(len(data)==3){
		//Se recibieron los valores del reloj, se verifica consistencia y se actualiza data en struct del planeta.
		dataX,_ := strconv.Atoi(data[0])
		dataY,_ := strconv.Atoi(data[1])
		dataZ,_ := strconv.Atoi(data[2])
		flag := 1
		for i:= range planetas {
			if planetas[i].nombre == comando[1]{
				if ((dataX >= planetas[i].relojx) && (dataY >= planetas[i].relojy) && (dataZ >= planetas[i].relojz)){
					planetas[i].relojx = dataX
					planetas[i].relojy = dataY
					planetas[i].relojz = dataZ
					planetas[i].lastfulcrum = direccionFulcrum
					flag = 0
					fmt.Println("\n[*] Sin Error de consistencia! \n")
					break 
				} else {
					fmt.Println("[*] Error de consistencia!")
					flag = 0
					break
				}
			}
		}
		if(flag == 1){
			//Quiere decir que no se maneja info del planeta y el archivo fue creado.
			planetas = append(planetas, Cplaneta(comando[1], dataX, dataY, dataZ, direccionFulcrum))
		}
	} else {
		//Error, no se hace nada
	}
}

func scanMsg()(mensaje string){
	scanner := bufio.NewScanner(os.Stdin)
	var PromptC = ""
	fmt.Println("Escriba el comando a ejecutar (0 para cerrar programa)")
	fmt.Println("Recuerde ser consistente con mayúsculas y minúsculas para los comandos\n")
	if scanner.Scan() {
		PromptC = scanner.Text()
	}
	mensaje = PromptC
	return
}

func main() {
	/*
	===== Pruebas para la creacion de un planeta =====
	e := Cplaneta("TESTTTT")
	e.relojx++
	fmt.Println("Nombre del planeta: "+e.nombre)
	fmt.Println("Int en X: ")
	fmt.Println(e.relojx)
	fmt.Println(e.relojy)
	fmt.Println(e.relojz)
	*/
	
	/*
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("No se puede conectar al servidor "+ err.Error())
	}

	serviceClient := pb.NewEntradaMensajeClient(conn)
	*/
	
	/* 
	======Para enviar un mje======		
	res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
	Body: valor+" 2",
	})
	
	if err != nil {
		panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
	}
	
	fmt.Println(res.Body)
	*/
	mensaje:="-1"
	for(mensaje != "0"){
		mensaje:=scanMsg()
		if(mensaje == "0"){
			break
		}
		processMsg(mensaje)
	}

}