package main

import(
	"fmt"
	"github.com/streadway/amqp"
	"strings"
	"os"
	"strconv"
)
/*
FORMATO PARA JUGADOR ELIMINADO: ("MUERTE",JUGADOR,RONDA)
*/

func main() {
	f, err := os.Create("jugadoresEliminados.txt")
	defer f.Close()
	pozo := 11
	fmt.Println("Aplicacion consumidor (Server)")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch,err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err :=  ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs{
			fmt.Printf("Mensaje recibido: %s\n", d.Body)
			response := strings.Split(string(d.Body),",")
			if response[0] == "MUERTE"{
				pozo += 100000000
				data := []byte(string(response[1])+" "+string(response[2])+" "+strconv.Itoa(pozo)+"\n")
				_, err2 := f.Write(data)
				if err2 != nil {
			        fmt.Println(err2)
					panic(err2)
			    }
			}

			fmt.Println(pozo)
			//implementar muerte jugador en el .txt
		}

	}()

	fmt.Println("Conectado a la instancia de RabbitMQ")
	fmt.Println("[*] - Esperando mensajes")
	
	<-forever
}