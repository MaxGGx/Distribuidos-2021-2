package main

import (
	"fmt"

	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
)

func Opciones() {
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

		} else {
			flag = true
			fmt.Println("ingrese una opcion correcta")
		}
	}

}

/*
func Solicitud(serviceClient *pb.EntradaMensajeClient) {
	res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
		Body: "ARCHIVO,Jugador_9,Ronda_3",
	})

	if err != nil {
		panic("Mensaje no pudo ser creado ni enviado: " + err.Error())
	}

	fmt.Println(res.Body)

	return res.body
}*/
func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic("No se puede conectar al servidor " + err.Error())
	}
	serviceClient := pb.NewEntradaMensajeClient(conn)
	/*
		res, err := serviceClient.Intercambio(context.Background(), &pb.Mensaje{
			Body: "ARCHIVO,Jugador_9,Ronda_3",
		})

		if err != nil {
			panic("Mensaje no pudo ser creado ni enviado: " + err.Error())
		}

		fmt.Println(res.Body)
	*/
	var unirse int
	fmt.Println("Desea unirse al Juego?")
	fmt.Println("1. Si")
	fmt.Println("2. No")

	flag := true

	for flag {
		fmt.Scanln(&unirse)
		if unirse == 1 {
			fmt.Println("Enviar solicitud de unirse")
			flag = false
		} else if unirse == 2 {
			return
		} else {
			flag = true
		}
	}

	autorizado := true // Autorizacion de parte del lider

	if autorizado {
		var jugada int

		Vivo := true //Condicion de si puede seguir jugando o no, verificada con el lider
		ronda := 1
		Total := 0

		Opciones()

		fmt.Printf("\n----------Etapa 1----------\n\n")
		fmt.Printf("Luz Roja, Luz Verde\n\n")
		fmt.Println("Reglas:")
		fmt.Println("- Elegir un numero entre 1 y 10, para sumar 21")
		fmt.Println()

		for Vivo && ronda <= 4 {

			fmt.Printf("----------Ronda %d----------\n\n", ronda)
			fmt.Println("Ingrese su numero entre 1 y 10")
			fmt.Scanln(&jugada)

			//quiza esto va en el lider
			if jugada < 1 || 10 < jugada {
				Vivo = false
			}

			//for jugada < 0 || 10 < jugada{}

			Total += jugada
			fmt.Printf("Total del jugador: %d\n", Total)

			if Vivo {
				fmt.Println("Estado del Jugador: Vivo")
			} else {
				fmt.Println("Estado del Jugador: Morido")
				return
			}
			ronda++
		}

		Opciones()

		fmt.Printf("\n----------Etapa 2----------\n\n")
		fmt.Printf("Tirar la cuerda\n\n")
		fmt.Println("Reglas:")
		fmt.Println("- Elegir un numero entre 1 y 4 para igualar la paridad del\n numero elegido por el lider")
		fmt.Println()

		//Concultar si el jugador sigue vivo despues de imparidad

		if Vivo {
			//fmt.Printf("----------Ronda %d----------\n\n", ronda)
			fmt.Println("Ingrese su numero entre 1 y 4")
			fmt.Scanln(&jugada)
			fmt.Println("Eleccion del numero para la Etapa 2:", jugada)

			//fmt.Printf("Total de los jugadores: %d\n", respuesta)
		} else {
			fmt.Println("Estado del Jugador: Morido")
			return
		}

		Opciones()
		ronda = 1

		for Vivo {
			fmt.Printf("\n----------Etapa 3----------\n\n")
			fmt.Printf("Todo o Nada\n\n")
			fmt.Println("Reglas:")
			fmt.Println("- Elegir un numero entre 1 y 10 ")
			fmt.Println()

			fmt.Scanln(&jugada)
			fmt.Printf("Eleccion del numero para la Etapa 3, ronda %d: %d\n", ronda, jugada)

		}
	}
}
