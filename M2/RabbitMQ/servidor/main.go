package main

import(
	"fmt"
	"github.com/streadway/amqp"
	"strings"
)
/*
FORMATO PARA JUGADOR ELIMINADO: ("MUERTE",JUGADOR,RONDA)
*/

func main() {
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
			}
			fmt.Println(pozo)
			//implementar muerte jugador en el .txt
		}

	}()

	fmt.Println("Conectado a la instancia de RabbitMQ")
	fmt.Println("[*] - Esperando mensajes")
	
	<-forever
}