package main

import(
	"fmt"
	"github.com/streadway/amqp"
	"strings"
	"os"
	"strconv"
	"context"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
	"net"
)
/*
FORMATO PARA JUGADOR ELIMINADO: ("MUERTE",JUGADOR,RONDA)
FORMATO PARA POZO ("POZO")
*/
var pozo = 0

type server struct {
	pb.UnimplementedEntradaMensajeServer
}

func (s *server ) Intercambio (ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	res := ""
	fmt.Println("Se recibió el siguiente mensaje: "+ req.Body)
	response := strings.Split(string(req.Body)," ")
	return &pb.Mensaje{Body: string(res)}, nil 
}

//Toma los mjes de llegada y los procesa, mostrando por pantalla los mensajes acutales.
func Recepcion(mensaje string) (resmje string){
	temp := strings.Split(mensaje, " ")
	valor, _ := strconv.Atoi(temp[0])
	fmt.Println("Mensajes recibidos:")
	fmt.Println(temp)
	resmje = "[*] Respuesta Recibida"
	return string(valor)
}

//Retorno de respuesta a jugador que la solicitó
func Delivery(mensaje string) (respuesta string){
	temp := strings.Split(mensaje, " ")
	valor,_ := strconv.Atoi(temp[0])
	//respuesta = respuestas[valor]
	return string(valor)

}

func main() {
	f, err := os.Create("jugadoresEliminados.txt")
	defer f.Close()
	
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
	
	go func(){
		fmt.Println("ANDANDO")
		listener, err := net.Listen("tcp", ":50060")

		if err != nil {
			panic("No se puede crear la conexión tcp: "+ err.Error())
		}
		fmt.Println(listener)
		serv := grpc.NewServer()
		pb.RegisterEntradaMensajeServer(serv, &server{})
		if err = serv.Serve(listener); err != nil {
			panic("No se ha podido inicializar el servidor: "+ err.Error())
		}
	}()



	forever := make(chan bool)
	go func() {
		for d := range msgs{
			fmt.Printf("Mensaje recibido: %s\n", d.Body)
			response := strings.Split(string(d.Body),",")
			if response[1] == "DEAD"{
				pozo += 100000000
				data := []byte("Jugador_"+string(response[0])+" Ronda_"+string(response[2])+" "+strconv.Itoa(pozo)+"\n")
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