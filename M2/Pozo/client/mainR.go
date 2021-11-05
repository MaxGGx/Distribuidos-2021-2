package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

//Este es el codigo para el cliente de la aplicacion, se conecta a la cola de rabbit para postear un mensaje o solicitud, luego el servidor hara pull a los mensajes para procesarlos

func main() {
	fmt.Println("Test RabbitMQ")

	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial("amqp://guest:guest@dist33:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Si lees esto es porque se ejecuto bien la conexion a la instancia de RabbitMQ")


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

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("J1 DEAD 3"),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Si lees este otro mensaje, quiere decir que se publico un mensaje exitosamente en la cola TestQueue")
}