package main

import(
	"fmt"
	"github.com/streadway/amqp"
)
//J1 DEAD
func main() {
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
		}
	}()

	fmt.Println("Conectado a la instancia de RabbitMQ")
	fmt.Println("[*] - Esperando mensajes")

	<-forever
}