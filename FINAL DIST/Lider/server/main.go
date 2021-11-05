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
	"math"
	"github.com/streadway/amqp"

)

type server struct {
	pb.UnimplementedEntradaMensajeServer
}

var flagListo = 0
var solicitudes [17]string
var respuestas [17]string
var status = [17]int {0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
//var conn1 *grpc.ClientConn

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

//Es como un len de python, cuenta los valores 1 de un array.
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
	//fmt.Println("Mensajes de jugadores recibidos:")
	//fmt.Println(solicitudes)
	if tamanio(solicitudes) == tamanio2(status){
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
	
	//AGREGAR IP DE CONEXION A SERVER GRPC POZO
	
	var res string
	//fmt.Println("Se recibió el siguiente mensaje: "+ req.Body)
	//Mientras espera mensaje jugador debe escribir "<N jugador> Listo?"
	if strings.Contains(req.Body, "Listo?"){
		if (flagListo == 1) && (PendingRequest(req.Body) == 0) {
			res = Delivery(req.Body)
		}else {
			res = "[*] Processing..."
		}
		
	//Para pedir monto de pozo jugador debe hacer "<N jugador> POZO"
	}else if strings.Contains(req.Body, "POZO"){

		conn1, err1 := grpc.Dial("localhost:50056", grpc.WithInsecure())
		if err1 != nil {
			panic("No se puede conectar al Data Node 1 "+ err1.Error())
		}
		serviceClient1 := pb.NewEntradaMensajeClient(conn1)

		solicitud := "POZO"
	 	res1, err := serviceClient1.Intercambio(context.Background(), &pb.Mensaje{
			Body: solicitud,
		})
		fmt.Println("ALSDJASLDJKASLDASKDLKASL UWU")
		fmt.Println(res1.Body)
		if err != nil {
			panic("Mensaje no pudo ser creado ni enviado: "+ err.Error())
		}
		res = res1.Body
		if res == ""{
			res = "SIN DATOS"
		}
		defer conn1.Close()
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

	go initServer()
	
	//CONEXION A SERVIDOR RABBIT MQ AGREGAR IP AQUI
	
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(q)



	//CONEXION A NAME NODE IP: M3:
	conn3, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic("No se puede conectar al servidor "+ err.Error())
	}

	serviceClient := pb.NewEntradaMensajeClient(conn3)

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
			fmt.Println(res.Body)
		}

		//Procesado de respuesta
		resR1,_ := strconv.Atoi(solicitudes[i])
		if resR1 == ronda1valor{
			fmt.Println("Jugador "+stringjugador+" Ha MUERTO")
			status[i] = 0
			respuestas[i] = "MUERTO 1"
			err = ch.Publish(
				"",
				"TestQueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body: []byte("MUERTE,"+stringjugador+",1"),
				},
			)

			if err != nil {
				fmt.Println(err)
				panic(err)
			}
		} else {
			juego1sumas[i] = resR1
			respuestas[i] = "VIVO 1"
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
		if status[i] == 1{
			//Previo a procesar la respuesta se registra en el log de NameNode:
			stringjugador = strconv.Itoa(i)
			
			res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "JUGA,"+stringjugador+",1,"+solicitudes[i],
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+err.Error())	
			}else{
				fmt.Println(res.Body)
			}

			//Procesado de respuesta
			resR2,_ := strconv.Atoi(solicitudes[i])
			if resR2 == ronda2valor{
				fmt.Println("Jugador "+stringjugador+" Ha MUERTO")
				status[i] = 0
				respuestas[i] = "MUERTO 1"
				err = ch.Publish(
				"",
				"TestQueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body: []byte("MUERTE,"+stringjugador+",1"),
				},
				)

				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			} else {
				juego1sumas[i] += resR2
				respuestas[i] = "VIVO 1"
			}
		}
	}
	VaciarSolicitudes()
	flagListo = 1

	//RECEPCION DE SOLICITUDES POR PARTE DE LOS USUARIOS PARA RONDA 3
	//INICIALIZACION VARIABLES RONDA 3
	rand.Seed(time.Now().UnixNano())
	ronda3valor := (rand.Intn(4)+6)

	fmt.Println("[*] Esperando respuestas Ronda 3 Juego 1...")
	for tamanio(solicitudes) < tamanio2(status){
		
	}
	flagListo = 0
	VaciarRespuestas()

	for i:=1; i<17; i++{
		if status[i] == 1{
			//Previo a procesar la respuesta se registra en el log de NameNode:
			stringjugador := strconv.Itoa(i)
			
			res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "JUGA,"+stringjugador+",1,"+solicitudes[i],
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+ err.Error())	
			}else{
				fmt.Println(res.Body)
			}

			//Procesado de respuesta
			resR3,_ := strconv.Atoi(solicitudes[i])
			if resR3 == ronda3valor{
				fmt.Println("Jugador "+stringjugador+" Ha MUERTO")
				status[i] = 0
				respuestas[i] = "MUERTO 1"
				err = ch.Publish(
				"",
				"TestQueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body: []byte("MUERTE,"+stringjugador+",1"),
				},
				)

				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			} else {
				juego1sumas[i] += resR3
				respuestas[i] = "VIVO 1"
			}
		}
	}
	VaciarSolicitudes()
	flagListo = 1

	//RECEPCION DE SOLICITUDES POR PARTE DE LOS USUARIOS PARA RONDA 4
	//INICIALIZACION VARIABLES RONDA 4
	rand.Seed(time.Now().UnixNano())
	ronda4valor := (rand.Intn(4)+6)

	fmt.Println("[*] Esperando respuestas Ronda 4 Juego 1...")
	for tamanio(solicitudes) < tamanio2(status){
		
	}
	flagListo = 0
	VaciarRespuestas()

	for i:=1; i<17; i++{
		if status[i] == 1{
			if juego1sumas[i] < 21{
				//Previo a procesar la respuesta se registra en el log de NameNode:
				stringjugador := strconv.Itoa(i)
				
				res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
				Body: "JUGA,"+stringjugador+",1,"+solicitudes[i],
				})
				if err != nil {
					panic("Error con la solicitud para ir registrando el historial de un jugador: "+ err.Error())	
				}else{
					fmt.Println(res.Body)
				}

				//Procesado de respuesta
				resR4,_ := strconv.Atoi(solicitudes[i])
				if resR4 == ronda4valor{
					fmt.Println("Jugador "+stringjugador+" Ha MUERTO")
					status[i] = 0
					respuestas[i] = "MUERTO FIN"
					err = ch.Publish(
						"",
						"TestQueue",
						false,
						false,
						amqp.Publishing{
							ContentType: "text/plain",
							Body: []byte("MUERTE,"+stringjugador+",1"),
						},
						)

						if err != nil {
							fmt.Println(err)
							panic(err)
						}
				} else {
					juego1sumas[i] += resR4
					respuestas[i] = "VIVO FIN"
				}
			}else{
				respuestas[i] = "VIVO FIN"
			}
		}
	}
	VaciarSolicitudes()
	flagListo = 1

	//PREPARACION JUEGO 2:
	fmt.Println("[*] Preparando todo para Juego 2...")
	for tamanio(solicitudes) < tamanio2(status){
		
	}
	flagListo = 0
	VaciarRespuestas()
	if tamanio2(status)%2 != 0{
		actual := tamanio2(status)
		rand.Seed(time.Now().UnixNano())
		amatar := (rand.Intn(15)+1)
		for tamanio2(status) == actual{
			if status[amatar] == 1{
				status[amatar] = 0
				fmt.Println("Jugador "+strconv.Itoa(amatar)+" Ha MUERTO")
				err = ch.Publish(
				"",
				"TestQueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body: []byte("MUERTE,"+strconv.Itoa(amatar)+",2"),
				},
				)

				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			} else{
				amatar = (rand.Intn(15)+1)
			}
		}
	}
	for i:=1; i<17; i++{
		if status[i] == 1{
			respuestas[i] = "VIVO"
		} else {
			respuestas[i] = "MUERTO"
		}
	}
	//Armado de equipos:
	var equipo1 []int
	var Nequipo1 = 0
	var Nequipo2 = 0
	aequipo := (rand.Intn(15)+1)
	for len(equipo1) != (tamanio2(status)/2){
		if (status[aequipo] == 1){
			equipo1 = append(equipo1,aequipo)
			aequipo = (rand.Intn(15)+1)
		}else{
			aequipo = (rand.Intn(15)+1)
		}
	}
	VaciarSolicitudes()
	flagListo = 1

	// ###############################################################################
	//JUEGO 2 
	// ###############################################################################
	for tamanio(solicitudes) < tamanio2(status){
		
	}
	flagListo = 0
	VaciarRespuestas()
	//SE ASUME QUE A ESTE PUNTO SE RECIBEN TODAS LAS SOLICITUDES.
	//PROCESAMIENTO JUEGO 1:
	PromptLider = -1
	for PromptLider != 0{
		fmt.Println("Todo listo para el Juego 2!")
		fmt.Println("Seleccione una opción para continuar escribiendo un número:\n")
		fmt.Println("0) Dar inicio al juego 1 de Squid Game <コ:彡")
		fmt.Println("Ingrese un numero del 1 al 16 para consultar el historial de un jugador")
		fmt.Scanln(&PromptLider)
		if PromptLider != 0{
			res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "DATA,Jugador_"+strconv.Itoa(PromptLider)+",Ronda_1",
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+err.Error())	
			}else{
				fmt.Println(res.Body)
			}
			PromptLider = -1
		}
	}
	fmt.Println("[*] Esperando respuestas Juego 2...")
	NJuego2 := (rand.Intn(15)+1)
	for i:=1; i<17; i++{
		banderita := 0
		if status[i] == 1{
			ingresado,_ := strconv.Atoi(solicitudes[i])
			res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "JUGA,"+strconv.Itoa(i)+",2,"+solicitudes[i],
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+ err.Error())	
			}else{
				fmt.Println(res.Body)
			}
			for _,b := range equipo1{
				if b == i{
					banderita = 1
					Nequipo1+=ingresado
				}	
			}
			if banderita != 1{
				Nequipo2+=ingresado
			}
		}
	}
	//Si las paridades son iguales pero distintas a las del lider.
	flagmatanza := 0
	if (Nequipo1%2 == Nequipo2%2) && (Nequipo1%2 != NJuego2%2) {
		//SE MATA UN EQUIPO AL AZAR
		flagmatanza = 1
		equipoamatar := (rand.Intn(1)+1)
		if equipoamatar == 1{
			for i,_ := range status{
				for _,b := range equipo1{
					if i == b{
						respuestas[i] = "MUERTO"
						status[i] = 0
						fmt.Println("Jugador "+strconv.Itoa(i)+"ha MUERTO")
						err = ch.Publish(
							"",
							"TestQueue",
							false,
							false,
							amqp.Publishing{
								ContentType: "text/plain",
								Body: []byte("MUERTE,"+strconv.Itoa(i)+",2"),
							},
							)

							if err != nil {
								fmt.Println(err)
								panic(err)
							}
					}
				}
				if respuestas[i] != "MUERTO"{
					respuestas[i] = "VIVO"
				}
			} 
		}else{
			for i,_ := range status{
				for _,b := range equipo1{
					if i == b{
						respuestas[i] = "VIVO"
					}
				}
				if respuestas[i] != "VIVO"{
					respuestas[i] = "MUERTO"
					status[i] = 0
					fmt.Println("Jugador "+strconv.Itoa(i)+"ha MUERTO")
					err = ch.Publish(
						"",
						"TestQueue",
						false,
						false,
						amqp.Publishing{
							ContentType: "text/plain",
							Body: []byte("MUERTE,"+strconv.Itoa(i)+",2"),
						},
						)

						if err != nil {
							fmt.Println(err)
							panic(err)
						}
				}
			} 
		}
	}else if (Nequipo1%2 == NJuego2%2){
		flagmatanza = 1
		for i,_ := range status{
			for _,b := range equipo1{
				if i == b{
					respuestas[i] = "VIVO"
				}
			}
			if respuestas[i] != "VIVO"{
				respuestas[i] = "MUERTO"
				status[i] = 0
				fmt.Println("Jugador "+strconv.Itoa(i)+"ha MUERTO")
				err = ch.Publish(
					"",
					"TestQueue",
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body: []byte("MUERTE,"+strconv.Itoa(i)+",2"),
					},
					)

					if err != nil {
						fmt.Println(err)
						panic(err)
					}
			}
		}
	}else{
		flagmatanza = 1
		for i,_ := range status{
			for _,b := range equipo1{
				if i == b{
					respuestas[i] = "MUERTO"
					status[i] = 0
					fmt.Println("Jugador "+strconv.Itoa(i)+"ha MUERTO")
					err = ch.Publish(
					"",
					"TestQueue",
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body: []byte("MUERTE,"+strconv.Itoa(i)+",2"),
					},
					)

					if err != nil {
						fmt.Println(err)
						panic(err)
					}
				}
			}
			if respuestas[i] != "MUERTO"{
				respuestas[i] = "VIVO"
			}
		}
	}
	if flagmatanza == 0{
		for i:=1 ; i<17; i++{
			respuestas[i] = "VIVO"
		}
	}
	VaciarSolicitudes()
	flagListo = 1
	// ###############################################################################
	//JUEGO 3 
	// ###############################################################################
	if tamanio2(status)%2 != 0{
		actual := tamanio2(status)
		rand.Seed(time.Now().UnixNano())
		amatar := (rand.Intn(15)+1)
		for tamanio2(status) == actual{
			if status[amatar] == 1{
				status[amatar] = 0
				fmt.Println("Jugador "+strconv.Itoa(amatar)+"ha MUERTO")
				err = ch.Publish(
					"",
					"TestQueue",
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body: []byte("MUERTE,"+strconv.Itoa(amatar)+",3"),
					},
					)

					if err != nil {
						fmt.Println(err)
						panic(err)
					}
			} else{
				amatar = (rand.Intn(15)+1)
			}
		}
	}
	for i:=1; i<17; i++{
		if status[i] == 1{
			respuestas[i] = "VIVO"
		} else {
			respuestas[i] = "MUERTO"
		}
	}
	//Armado de equipos:
	var equipos [][]int
	var tempequi []int
	for i,v := range status{
		if v == 1{
			tempequi = append(tempequi,i)
			if len(tempequi) == 2{
				equipos = append(equipos, tempequi)
				tempequi = nil
			}
		}
	}
	for tamanio(solicitudes) < tamanio2(status){
		
	}
	flagListo = 0
	VaciarRespuestas()
	PromptLider = -1
	for PromptLider != 0{
		fmt.Println("Todo listo para el Juego 3! RONDA FINAL")
		fmt.Println("Seleccione una opción para continuar escribiendo un número:\n")
		fmt.Println("0) Dar inicio al juego 1 de Squid Game <コ:彡")
		fmt.Println("Ingrese un numero del 1 al 16 para consultar el historial de un jugador")
		fmt.Scanln(&PromptLider)
		if PromptLider != 0{
			res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "DATA,Jugador_"+strconv.Itoa(PromptLider)+",Ronda_1",
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+err.Error())	
			}else{
				fmt.Println(res.Body)
			}
			res, err = serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "DATA,Jugador_"+strconv.Itoa(PromptLider)+",Ronda_2",
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+err.Error())	
			}else{
				fmt.Println(res.Body)
			}
			PromptLider = -1
		}
	}
	fmt.Println("[*] Esperando respuestas Juego 3...")
	NJuego3 := (rand.Intn(9)+1)
	for _,v := range equipos{
		jugador1 := v[0] 
		vjugador1,_ := strconv.Atoi(solicitudes[jugador1])
		jugador2 := v[1]
		vjugador2,_:= strconv.Atoi(solicitudes[jugador2])
		res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "JUGA,"+strconv.Itoa(jugador1)+",3,"+solicitudes[jugador1],
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+ err.Error())	
			}else{
				fmt.Println(res.Body)
			}
		res, err = serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "JUGA,"+strconv.Itoa(jugador2)+",3,"+solicitudes[jugador2],
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+ err.Error())	
			}else{
				fmt.Println(res.Body)
			}
		if math.Abs(float64(vjugador1 - NJuego3)) < math.Abs(float64(vjugador2 - NJuego3)){
			respuestas[jugador1] = "VIVO"
			respuestas[jugador2] = "MUERTO"
			status[jugador2] = 0
			fmt.Println("Jugador "+strconv.Itoa(jugador2)+"ha MUERTO")
			err = ch.Publish(
				"",
				"TestQueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body: []byte("MUERTE,"+strconv.Itoa(jugador2)+",3"),
				},
				)

				if err != nil {
					fmt.Println(err)
					panic(err)
				}
		} else if math.Abs(float64(vjugador1 - NJuego3)) > math.Abs(float64(vjugador2 - NJuego3)){
			respuestas[jugador2] = "VIVO"
			respuestas[jugador1] = "MUERTO"
			status[jugador1] = 0
			fmt.Println("Jugador "+strconv.Itoa(jugador1)+"ha MUERTO")
			err = ch.Publish(
				"",
				"TestQueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body: []byte("MUERTE,"+strconv.Itoa(jugador1)+",3"),
				},
				)

				if err != nil {
					fmt.Println(err)
					panic(err)
				}
		} else {
			respuestas[jugador1] = "VIVO"
			respuestas[jugador2] = "VIVO"
		}
	}
	VaciarSolicitudes()
	flagListo = 1
	fmt.Println("Squid Game TERMINADO")
	fmt.Println("GANADORES:")
	for i,v := range status{
		if v == 1{
			fmt.Println("Jugador "+strconv.Itoa(i))
		}
	}
	fmt.Println("Antes de finalizar, puede hacer lo siguiente:")
	PromptLider = -1
	for PromptLider != 0{
		fmt.Println("Seleccione una opción para continuar escribiendo un número:\n")
		fmt.Println("0) Dar término al Squid Game <コ:彡")
		fmt.Println("Ingrese un numero del 1 al 16 para consultar el historial de un jugador")
		fmt.Scanln(&PromptLider)
		if PromptLider != 0{
			res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "DATA,Jugador_"+strconv.Itoa(PromptLider)+",Ronda_1",
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+err.Error())	
			}else{
				fmt.Println(res.Body)
			}
			res, err = serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "DATA,Jugador_"+strconv.Itoa(PromptLider)+",Ronda_2",
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+err.Error())	
			}else{
				fmt.Println(res.Body)
			}
			res, err = serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "DATA,Jugador_"+strconv.Itoa(PromptLider)+",Ronda_3",
			})
			if err != nil {
				panic("Error con la solicitud para ir registrando el historial de un jugador: "+err.Error())	
			}else{
				fmt.Println(res.Body)
			}
			PromptLider = -1
		}
	}
	fmt.Println("Gracias por jugar...")
}