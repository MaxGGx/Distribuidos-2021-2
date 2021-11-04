package main

import (
	//"context"
	"fmt"
	"math/rand"
	"time"
	//pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	//"google.golang.org/grpc"
)

func Jugada(limit int) int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(limit) + 1
	return n
}

func IA(Jugador int) {
	autorizado := true // Autorizacion de parte del lider
	if autorizado {
		Vivo := true //Condicion de si puede seguir jugando o no, verificada con el lider
		ronda := 1

		fmt.Printf("\n----------Etapa 1----------\n\n")

		for Vivo && ronda <= 4 {
			n := Jugada(10)
			fmt.Printf("Jugador %d: Eleccion del numero para la Etapa 1, ronda %d: %d\n", Jugador, ronda, n)
			ronda++
		}

		if Vivo {
			fmt.Printf("\n----------Etapa 2----------\n\n")

			n := Jugada(4)
			fmt.Printf("Jugador %d: Eleccion del numero para la Etapa 2: %d", Jugador, n)
		}

		ronda = 1

		for Vivo {
			fmt.Printf("\n----------Etapa 3----------\n\n")
			n := Jugada(10)
			fmt.Printf("Jugador %d: Eleccion del numero para la Etapa 3, ronda %d: %d\n", Jugador, ronda, n)

			if Jugada(3) == 1 {
				Vivo = false
			}
		}
	}

}

func main() {
	nJugadores := 3

	for i := 1; i < nJugadores; i++ {
		fmt.Println(i)
		go IA(i)
	}
	x := 1
	for true {
		x++
	}
}
