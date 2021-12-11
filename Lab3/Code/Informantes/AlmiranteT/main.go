package main

import ( 
	"fmt"
	"context"
	pb "github.com/MaxGGx/Distribuidos-2021-2/tree/inaki/Lab3/Code/gRPC/proto"
	"google.golang.org/grpc"
)

//Struct para poder hacer un objeto de planeta dependiendo si es que esta creado.
type planeta struct {
	nombre string
	relojx int
	relojy int 
	relojz int 
}

//Constructor para el planeta, cosa de poder almacenar en memoria la info de los planetas manejados por la consola del informante.
func Cplaneta(name string)(planet planeta){
	planet = planeta{
		nombre: name,
		relojx: 0,
		relojy: 0,
		relojz: 0,
	}
	return
}

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
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
	if err != nil {
		panic("No se puede conectar al servidor " + err.Error())
	}
	serviceClient := pb.NewEntradaMensajeClient(conn)
	Solicitud(serviceClient, "HOLA")
	
}