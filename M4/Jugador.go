package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
)

func Solicitud(serviceClient pb.EntradaMensajeClient, msg string) string {
	res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: msg,
	})
	if err != nil {
		panic("Mensaje no pudo ser creado ni enviado: " + err.Error())
	}
	fmt.Println(res.Body)
	return res.Body
}

func Jugada(limit int) int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(limit) + 1
	return n

func Opciones(serviceClient pb.EntradaMensajeClient) {
	var jugada int
	flag := true
	for flag {
		fmt.Println("Que desea hacer a continuacion?")
		fmt.Println("1. Ir a la siguiente ronda")
		fmt.Println("2. Ver el monto del Pozo")
		fmt.Scanln(&jugada)

		if jugada == 1 {
			fmt.Println("Continundo a la siguiente etapa")
			flag = false

		} else if jugada == 2 {
			fmt.Println("Solicitando el monto del Pozo")
			res := Solicitud(serviceClient, "POZO")
			fmt.Println("El monto total del POZO es: ", res)

		} else {
			flag = true
			fmt.Println("ingrese una opcion correcta")
		}
	}
}

func IA(Jugador int, Channel chan int) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("No se puede conectar al servidor " + err.Error())
	}
	serviceClient := pb.NewEntradaMensajeClient(conn)

	autorizado := true // Autorizacion de parte del lider
	flag := false

	Solicitud(serviceClient, strconv.Itoa(Jugador)+" Sol1")
	flag = true
	fmt.Println()

	for flag {
		if Solicitud(serviceClient, strconv.Itoa(Jugador)+" Listo?") != "[*] Processing..." {
			flag = false
			autorizado = true
		}
	}

	if autorizado {

		var jugada int

		Vivo := true //Condicion de si puede seguir jugando o no, verificada con el lider

		////////////////////////Etapa 1/////////////////////////////
		Fin := false //Se llego a la ronda 4
		ronda := 1
		Total := 0

		for Vivo && !Fin {

			if Total <= 21 {
				//fmt.Printf("----------Ronda %d----------\n\n", ronda)
				//fmt.Println("Ingrese su numero entre 1 y 10")
				jugada = Jugada(10)
				Solicitud(serviceClient, strconv.Itoa(Jugador)+" "+strconv.Itoa(jugada))
			} else {
				Solicitud(serviceClient, strconv.Itoa(Jugador)+" nul")
			}

			flag = true

			for flag {
				res := Solicitud(serviceClient, strconv.Itoa(Jugador)+" Listo?")
				if res != "[*] Processing..." {
					flag = false
					l := strings.Split(res, " ")

					if l[0] != "VIVO" {
						Vivo = false
					}

					if l[1] == "FIN" {
						Fin = true
					}
				}
			}

			Total += jugada
			//fmt.Printf("Total del jugador: %d\n", Total)

			if !Vivo {
				//fmt.Println("Has perdido, estas muerto")
				Channel <- 1
				return 
			}
			ronda++
		}

		////////////////////////Etapa 2/////////////////////////////

		//Verificar si el jugador no fue eliminado por azar
		Solicitud(serviceClient, strconv.Itoa(Jugador)+" Sol2")
		for flag {
			res := Solicitud(serviceClient, strconv.Itoa(Jugador)+" Listo?")
			if res != "[*] Processing..." {
				if res == "MUERTO" {
					//fmt.Println("Has sido eliminado")
					Channel <- 1
					return 
				}
			}
		}
		flag = true

		fmt.Printf("\n----------Etapa 2----------\n\n")
		//fmt.Printf("Tirar la cuerda\n\n")
		//fmt.Println("Reglas:")
		//fmt.Println("- Elegir un numero entre 1 y 4 para igualar la paridad del\n numero elegido por el lider")
		//fmt.Println()

		//fmt.Println("Ingrese su numero entre 1 y 4")

		jugada = Jugada(4)
		Solicitud(serviceClient, strconv.Itoa(Jugador)+" "+strconv.Itoa(jugada))

		flag = true
		for flag {
			res := Solicitud(serviceClient, strconv.Itoa(Jugador)+" Listo?")
			if res != "[*] Processing..." {
				if res == "MUERTO" {
					//fmt.Println("Has sido eliminado")
					Channel <- 1
					return
				}
			}
		}
		////////////////////////Etapa 3/////////////////////////////
		ronda = 1
		//Ganador := false

		//for Vivo && !Ganador {
		//fmt.Printf("\n----------Etapa 3----------\n\n")
		//fmt.Printf("Todo o Nada\n\n")
		//fmt.Println("Reglas:")
		//fmt.Println("- Elegir un numero entre 1 y 10 ")
		//fmt.Println()

		//fmt.Printf("Eleccion del numero para la Etapa 3, ronda %d\n", ronda)

		jugada = Jugada(10)
		Solicitud(serviceClient, strconv.Itoa(Jugador)+" "+strconv.Itoa(jugada))

		flag = true
		for flag {
			res := Solicitud(serviceClient, strconv.Itoa(Jugador)+" Listo?")
			if res != "[*] Processing..." {
				flag = false

				if res != "VIVO" {
					//fmt.Println("Has sido eliminado")
					Channel <- 1
					return
				}
			}
		}
		//ronda += 1
		//}
		//fmt.Println("Has ganado el Juego del Calamar")
		Channel <- 1
	}
}

func Jugador(Channel chan int){

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("No se puede conectar al servidor " + err.Error())
	}
	serviceClient := pb.NewEntradaMensajeClient(conn)

	var unirse int
	autorizado := true // Autorizacion de parte del lider
	flag := false

	fmt.Println("Desea unirse al Juego?")
	fmt.Println("1. Si")
	fmt.Println("2. No")

	for flag {
		fmt.Scanln(&unirse)
		if unirse == 1 {
			fmt.Println("Enviar solicitud de unirse")
			flag = false
		} else if unirse == 2 {
			return 1
		} else {
			flag = true
		}
	}

	Solicitud(serviceClient, "16 Sol1")
	flag = true
	fmt.Println()
	for flag {
		if Solicitud(serviceClient, "16 Listo?") != "[*] Processing..." {
			flag = false
			autorizado = true
		}

	}

	if autorizado {

		var jugada int

		Vivo := true //Condicion de si puede seguir jugando o no, verificada con el lider

		////////////////////////Etapa 1/////////////////////////////
		Fin := false //Se llego a la ronda 4
		ronda := 1
		Total := 0

		Opciones(serviceClient)

		fmt.Printf("\n----------Etapa 1----------\n\n")
		fmt.Printf("Luz Roja, Luz Verde\n\n")
		fmt.Println("Reglas:")
		fmt.Println("- Elegir un numero entre 1 y 10, para sumar 21")
		fmt.Println()

		for Vivo && !Fin {

			if Total <= 21 {
				fmt.Printf("----------Ronda %d----------\n\n", ronda)
				fmt.Println("Ingrese su numero entre 1 y 10")
				fmt.Scanln(&jugada)
				Solicitud(serviceClient, "16 "+strconv.Itoa(jugada))
			} else {
				Solicitud(serviceClient, "16 nul")
			}

			flag = true

			for flag {
				res := Solicitud(serviceClient, "16 Listo?")
				if res != "[*] Processing..." {
					flag = false
					l := strings.Split(res, " ")

					if l[0] != "VIVO" {
						Vivo = false
					}

					if l[1] == "FIN" {
						Fin = true
					}
				}
			}

			Total += jugada
			fmt.Printf("Total del jugador: %d\n", Total)

			if !Vivo {
				fmt.Println("Has perdido, estas muerto")
				Channel <- 1
				return
			}
			ronda++
		}

		////////////////////////Etapa 2/////////////////////////////

		//Verificar si el jugador no fue eliminado por azar
		Solicitud(serviceClient, "16 Sol2")
		for flag {
			res := Solicitud(serviceClient, "16 Listo?")
			if res != "[*] Processing..." {
				if res == "MUERTO" {
					fmt.Println("Has sido eliminado")
					Channel <- 1
					return
				}
			}
		}
		flag = true
		Opciones(serviceClient)

		fmt.Printf("\n----------Etapa 2----------\n\n")
		fmt.Printf("Tirar la cuerda\n\n")
		fmt.Println("Reglas:")
		fmt.Println("- Elegir un numero entre 1 y 4 para igualar la paridad del\n numero elegido por el lider")
		fmt.Println()

		fmt.Println("Ingrese su numero entre 1 y 4")
		fmt.Scanln(&jugada)

		Solicitud(serviceClient, "16 "+strconv.Itoa(jugada))
		flag = true
		for flag {
			res := Solicitud(serviceClient, "16 Listo?")
			if res != "[*] Processing..." {
				if res == "MUERTO" {
					fmt.Println("Has sido eliminado")
					Channel <- 1
					return
				}
			}
		}
		////////////////////////Etapa 3/////////////////////////////
		Solicitud(serviceClient, "16 Sol3")
		for flag {
			res := Solicitud(serviceClient, "16 Listo?")
			if res != "[*] Processing..." {
				if res == "MUERTO" {
					fmt.Println("Has sido eliminado")
					Channel <- 1
					return
				}
			}
		}
		Opciones(serviceClient)
		ronda = 1
		//Ganador := false

		//for Vivo && !Ganador {
		fmt.Printf("\n----------Etapa 3----------\n\n")
		fmt.Printf("Todo o Nada\n\n")
		fmt.Println("Reglas:")
		fmt.Println("- Elegir un numero entre 1 y 10 ")
		fmt.Println()

		fmt.Printf("Eleccion del numero para la Etapa 3, ronda %d\n", ronda)

		fmt.Scanln(&jugada)
		Solicitud(serviceClient, "16 "+strconv.Itoa(jugada))

		flag = true
		for flag {
			res := Solicitud(serviceClient, "16 Listo?")
			if res != "[*] Processing..." {
				flag = false
				if res != "VIVO" {
					fmt.Println("Has sido eliminado")
					Channel <- 1
					return
				}
			}
		}
		//	ronda += 1
		//}
		fmt.Println("Has ganado el Juego del Calamar")
		Channel <- 1
		return
	}
}

func main (){
	nJugadores := 15
	Channel := make(chan int, nJugadores)
	
	for i := 0; i < nJugadores; i++ {
		go IA(i, Channel)
	}

	go Jugador(Channel)

	count := 1 
	for count < nJugadores+1 {
		<-Channel
		count++
	}
}