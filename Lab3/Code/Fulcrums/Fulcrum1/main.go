package main

import (
	"fmt"
	"time"
	"path/filepath"
	"strings"
	"context"
	pb "github.com/MaxGGx/Distribuidos-2021-2/M1/Test3/proto"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"os"
	"bufio"
	"log"
	"errors"
	"io/ioutil"
)

//Variables constantes
var ipFulcrum1 = "dist34:50002"
var ipFulcrum2 = "dist35:50003"
var ipFulcrum3 = "dist36:50004"

//Struct para guardar los relojes por planeta
type clock struct{
	planeta string
	relojx int
	relojy int
	relojz int
}

//guarda un set con los planetas y sus relojes por archivo.
var clkPlanets []clock

type server struct {
	pb.UnimplementedEntradaMensajeServer
}

func (s *server ) Intercambio (ctx context.Context, req *pb.Mensaje) (*pb.Mensaje, error) {
	ans := ""
	fmt.Println("Fulcrum 1 recibió el siguiente mensaje: "+ req.Body)
	if(strings.Split(req.Body, " ")[0] == "GetNumberRebelds"){
		ans = LeiaProcess(req.Body)
	} else if (strings.Split(req.Body, " ")[0] == "CLK"){
		ans = getCLK(req.Body) 
	} else if ((strings.Split(req.Body, ",")[0] == "MERGEU") || (strings.Split(req.Body, ",")[0] == "MERGEUA")){
		ans = processMergeu(req.Body)
	} else if (strings.Split(req.Body, ",")[0] == "MERGECLK"){
		ans = processMergeclk(req.Body)
	} else {
		ans = processInformante(req.Body)
	}
	return &pb.Mensaje{Body: ans}, nil 
}

//Procesa el request del informante (Dividiendo en los comandos que se enviaron)
func processInformante(comando string)(respuesta string){
	tipoC := strings.Split(comando, " ")[0]
	if(tipoC == "AddCity"){
		cP := strings.Split(comando, " ")
		if(len(cP) > 3){
			//Si no esta dentro del archivo, se agrega y si no existe, se crea agregando la linea.
    		f, err := os.OpenFile("planetas/"+cP[1]+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    		if err!= nil {
    			panic(err)
    		}
    		defer f.Close()

    		if _,err = f.WriteString(cP[1]+" "+cP[2]+" "+cP[3]+"\n"); err != nil {
    			panic(err)
    		}
    		//Si no esta dentro del archivo log, se agrega y si no existe, se crea agregando la linea.
    		f, err = os.OpenFile("log_planetas/"+cP[1]+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    		if err!= nil {
    			panic(err)
    		}
    		defer f.Close()

    		if _,err = f.WriteString(strings.Join(cP," ")+"\n"); err != nil {
    			panic(err)
    		}
		}else {
			//Si no esta dentro del archivo, se agrega y si no existe, se crea agregando la linea.
			//Solo que no se agrega el valor, sino un 0
    		f, err := os.OpenFile("planetas/"+cP[1]+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    		if err!= nil {
    			panic(err)
    		}
    		defer f.Close()

    		if _,err = f.WriteString(cP[1]+" "+cP[2]+" 0\n"); err != nil {
    			panic(err)
    		}
    		//Si no esta dentro del archivo log, se agrega y si no existe, se crea agregando la linea.
    		f, err = os.OpenFile("log_planetas/"+cP[1]+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    		if err!= nil {
    			panic(err)
    		}
    		defer f.Close()

    		if _,err = f.WriteString(strings.Join(cP," ")+"\n"); err != nil {
    			panic(err)
    		}

		}
		//En ambos casos, debe crearse/actualizarse el reloj del archivo para el planeta.
		for i:= range clkPlanets{
			if (clkPlanets[i].planeta == cP[1]){
				clkPlanets[i].relojx++
				relojx := strconv.Itoa(clkPlanets[i].relojx)
				relojy := strconv.Itoa(clkPlanets[i].relojy)
				relojz := strconv.Itoa(clkPlanets[i].relojz)
				respuesta = relojx+" "+relojy+" "+relojz
				return
			}
		}
		clkPlanets = append(clkPlanets, clock{cP[1], 1, 0, 0})
		respuesta = "1 0 0"
		return
	}else if (tipoC == "UpdateName"){
		cP := strings.Split(comando, " ")
		//Se procede con cambiar la linea
    	input,err := ioutil.ReadFile("planetas/"+cP[1]+".txt")
    	if err != nil {
    		log.Fatalln(err)
    	}
    	lines := strings.Split(string(input), "\n")

    	for i, line := range lines {
    		if strings.Contains(line, cP[1]+" "+cP[2]){
    			linea := strings.Split(lines[i], " ")
    			//fmt.Println("Lo encontre, escribiré: "+cP[1]+" "+cP[2]+" "+linea[2])
    			lines[i] = cP[1]+" "+cP[3]+" "+linea[2]
    		}
    	}
    	output := strings.Join(lines,"\n")
    	err = ioutil.WriteFile("planetas/"+cP[1]+".txt", []byte(output), 0644)
    	if err != nil {
    		log.Fatalln(err)
    	}
    	//Se actualiza Log
    	f, err := os.OpenFile("log_planetas/"+cP[1]+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    	if err!= nil {
    		panic(err)
    	}
    	defer f.Close()

    	if _,err = f.WriteString(strings.Join(cP," ")+"\n"); err != nil {
    		panic(err)
    	}
    	for i:= range clkPlanets{
			if (clkPlanets[i].planeta == cP[1]){
				clkPlanets[i].relojx++
				relojx := strconv.Itoa(clkPlanets[i].relojx)
				relojy := strconv.Itoa(clkPlanets[i].relojy)
				relojz := strconv.Itoa(clkPlanets[i].relojz)
				respuesta = relojx+" "+relojy+" "+relojz
			}
		}
		return
	}else if (tipoC == "UpdateNumber"){
		cP := strings.Split(comando, " ")
		//Se procede con cambiar la linea
    	input,err := ioutil.ReadFile("planetas/"+cP[1]+".txt")
    	if err != nil {
    		log.Fatalln(err)
    	}
    	lines := strings.Split(string(input), "\n")

    	for i, line := range lines {
    		if strings.Contains(line, cP[1]+" "+cP[2]){
    			lines[i] = cP[1]+" "+cP[2]+" "+cP[3]
    		}
    	}
    	output := strings.Join(lines,"\n")
    	err = ioutil.WriteFile("planetas/"+cP[1]+".txt", []byte(output), 0644)
    	if err != nil {
    		log.Fatalln(err)
    	}
    	//Se actualiza Log
    	f, err := os.OpenFile("log_planetas/"+cP[1]+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    	if err!= nil {
    		panic(err)
    	}
    	defer f.Close()

    	if _,err = f.WriteString(strings.Join(cP," ")+"\n"); err != nil {
    		panic(err)
    	}
    	for i:= range clkPlanets{
			if (clkPlanets[i].planeta == cP[1]){
				clkPlanets[i].relojx++
				relojx := strconv.Itoa(clkPlanets[i].relojx)
				relojy := strconv.Itoa(clkPlanets[i].relojy)
				relojz := strconv.Itoa(clkPlanets[i].relojz)
				respuesta = relojx+" "+relojy+" "+relojz
			}
		}
		return
	}else{
		cP := strings.Split(comando, " ")
		//Se procede con eliminar la linea
    	input,err := ioutil.ReadFile("planetas/"+cP[1]+".txt")
    	if err != nil {
    		log.Fatalln(err)
    	}
    	lines := strings.Split(string(input), "\n")

    	for i, line := range lines {
    		if strings.Contains(line, cP[1]+" "+cP[2]){
    			lines[i] = lines[len(lines)-1]
    			lines[len(lines)-1] = ""
    			lines = lines[:len(lines)-1]
    		}
    	}
    	output := strings.Join(lines,"\n")
    	err = ioutil.WriteFile("planetas/"+cP[1]+".txt", []byte(output), 0644)
    	if err != nil {
    		log.Fatalln(err)
    	}
    	//Se actualiza Log
    	f, err := os.OpenFile("log_planetas/"+cP[1]+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    	if err!= nil {
    		panic(err)
    	}
    	defer f.Close()

    	if _,err = f.WriteString(strings.Join(cP," ")+"\n"); err != nil {
    		panic(err)
    	}
    	for i:= range clkPlanets{
			if (clkPlanets[i].planeta == cP[1]){
				clkPlanets[i].relojx++
				relojx := strconv.Itoa(clkPlanets[i].relojx)
				relojy := strconv.Itoa(clkPlanets[i].relojy)
				relojz := strconv.Itoa(clkPlanets[i].relojz)
				respuesta = relojx+" "+relojy+" "+relojz
			}
		}
		return
	}
	return
}

//Procesa el MERGECLK (actualiza el reloj interno del archivo)
func processMergeclk(comando string)(respuesta string){
	planeta := strings.Split(comando, ",")[1]
	reloj := strings.Split(comando, ",")[2]
	relojX,_ := strconv.Atoi(strings.Split(reloj," ")[0])
	relojY,_ := strconv.Atoi(strings.Split(reloj," ")[1])
	relojZ,_ := strconv.Atoi(strings.Split(reloj," ")[2])
	for i:=range clkPlanets{
		if (clkPlanets[i].planeta == planeta){
			if(clkPlanets[i].relojx >= relojX){
				clkPlanets[i].relojx = relojX
			}
			if(clkPlanets[i].relojy >= relojY){
				clkPlanets[i].relojy = relojY
			}
			if(clkPlanets[i].relojz >= relojZ){
				clkPlanets[i].relojz = relojZ
			}
		}
	}
	respuesta = "Listo"
	return
}


//Procesa el MERGEU (para cuando se desea pasar las lineas del archivo).
func processMergeu(comando string)(respuesta string){
	planeta := strings.Split(comando, ",")[1]
	if(strings.Split(comando,",")[0] == "MERGEU"){
		//Orden de borrado y reset del planeta
		path:="planetas/"+strings.Split(comando,",")[1]+".txt"
    	f,err1:=os.Create(path)
    	if err1 != nil {
    		log.Fatal(err1)
    	}
    	defer f.Close()
    	f.Close()
	} else {
		linea := strings.Split(comando, ",")[2]
		//Si no esta dentro del archivo, se agrega y si no existe, se crea agregando la linea.
    	f, err := os.OpenFile("planetas/"+planeta+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    	if err!= nil {
    		panic(err)
    	}
    	defer f.Close()
    	if _,err = f.WriteString(linea+"\n"); err != nil {
    		panic(err)
    	}
    	flag:=0
    	for i := range clkPlanets{
    		if(clkPlanets[i].planeta == planeta){
    			flag=1
    		}
    	}
    	if(flag==0){
    		//Se crea ademas el clock para que luego sea actualizado
    		clkPlanets = append(clkPlanets,clock{planeta, 1, 0, 0})
    	}
    	
	}
	respuesta = "manejado"
	return
}
	/*
	planeta := strings.Split(comando, ",")[1]
	fmt.Println("Planeta: "+planeta)
	linea := strings.Split(comando, ",")[2]
	if _,err := os.Stat("planetas/"+planeta+".txt"); err == nil {
		file, err := os.Open("planetas/"+planeta+".txt")
    	if err!=nil{
    		log.Fatalf("Fallo en abrir archivo")
    	}
    	flag := 0
    	scanner := bufio.NewScanner(file)
    	scanner.Split(bufio.ScanLines)
    	for scanner.Scan() {
    		text := scanner.Text()
    		//Lucho seba || Lucho seba
    		if(strings.Contains(text, strings.Split(linea," ")[0]+" "+strings.Split(linea," ")[1])){
    			respuesta = "Sin Cambios"
    			return
    		} else if(strings.Split(text, " ")[1] == strings.Split(linea, " ")[1]){
    			//Es la misma ciudad solo que con un valor distinto, se deberá cambiar
    			flag = 1
    		}
    	}
    	file.Close()
    	if(flag == 1){
    		//Se procede con cambiar la linea
    		input,err := ioutil.ReadFile("planetas/"+planeta+".txt")
    		if err != nil {
    			log.Fatalln(err)
    		}
    		lines := strings.Split(string(input), "\n")

    		for i, line := range lines {
    			if strings.Contains(line, planeta+" "+strings.Split(linea," ")[1]){
    				lines[i] = linea
    			}
    		}
    		output := strings.Join(lines,"\n")
    		err = ioutil.WriteFile("planetas/"+planeta+".txt", []byte(output), 0644)
    		if err != nil {
    			log.Fatalln(err)
    		}
    	} else {
    		//Si no esta dentro del archivo, se agrega y si no existe, se crea agregando la linea.
    		f, err := os.OpenFile("planetas/"+planeta+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    		if err!= nil {
    			panic(err)
    		}
    		defer f.Close()

    		if _,err = f.WriteString(linea+"\n"); err != nil {
    			panic(err)
    		}
    	}	
	} else {
		//Si no esta dentro del archivo, se agrega y si no existe, se crea agregando la linea.
    	f, err := os.OpenFile("planetas/"+planeta+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    	if err!= nil {
    		panic(err)
    	}
    	defer f.Close()
    	if _,err = f.WriteString(linea+"\n"); err != nil {
    		panic(err)
    	}
    	//Se crea ademas el clock para que luego sea actualizado
    	clkPlanets = append(clkPlanets,clock{planeta, 1, 0, 0}) 

	}
	respuesta = "Manejado"
	return
}
*/

//Obtiene el reloj del archivo de un planeta, si no existe, se retorna un Err404
func getCLK(comando string)(clock string){
	planeta := strings.Split(comando," ")[1]
	for i:= range clkPlanets{
		if(clkPlanets[i].planeta == planeta){
			relojx := strconv.Itoa(clkPlanets[i].relojx)
			relojy := strconv.Itoa(clkPlanets[i].relojy)
			relojz := strconv.Itoa(clkPlanets[i].relojz)
			clock = relojx+" "+relojy+" "+relojz
			return
		}
	}
	clock = "Err404"
	return
}

//Funcion para el proceso de consultas de leia, arma la respuesta y la envia directamente
func getCity(planeta string, ciudad string)(data string){
	if _,err := os.Stat("planetas/"+planeta+".txt"); err == nil {
    	file, err := os.Open("planetas/"+planeta+".txt")
    	if err!=nil{
    		log.Fatalf("Fallo en abrir archivo")
    	}
    	flag:=1
    	scanner := bufio.NewScanner(file)
    	scanner.Split(bufio.ScanLines)
    	for scanner.Scan() {
    		text := scanner.Text()
    		if(strings.Split(text, " ")[1] == ciudad){
    			data = text
    			flag=0
    			break;
    		}
    	}
    	if (flag==1){
    		data = "Error: Ciudad no existe, verifique y vuelva a reintentar"
    	}
    	file.Close()
    	for i:=range clkPlanets{
    		if(clkPlanets[i].planeta == planeta){
    			relojx := strconv.Itoa(clkPlanets[i].relojx)
    			relojy := strconv.Itoa(clkPlanets[i].relojy)
    			relojz := strconv.Itoa(clkPlanets[i].relojz)
    			data = data+" "+relojx+","+relojy+","+relojz
    			break;
    		}
    	}
    } else if errors.Is(err, os.ErrNotExist){
    	data = "Error: Planeta no existe, verifique y vuelva a reintentar"
    }
    return
}

func LeiaProcess(comando string)(data string){
	planeta := strings.Split(comando, " ")
	data = getCity(planeta[1], planeta[2])
	return
}


//Funcion ejecutada por gRPC para enviar el mensaje
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

//Funcion que toma la ip a la que se desea enviar, se conecta y realiza el envio del mensaje. Retorna la respuesta
func enviarMsg(ip string, msg string)(answer string){
	conn, err := grpc.Dial(ip, grpc.WithInsecure())

	if err != nil {
		panic("No se puede conectar al servidor "+ err.Error())
	}

	serviceClient := pb.NewEntradaMensajeClient(conn)

	answer = Solicitud(serviceClient, msg)
	
	defer conn.Close()
	
	return
}

//Se encargará de realizar el Merge cada 2 minutos.
func timer(){
	c := time.Tick(120 * time.Second)
	for i := range c {
		if (i == i.Add(time.Second*1)){
			fmt.Println("xd")	
		}
		//Cada 2 min hará lo de aquí
    	//Obtener los nombres de archivos de planetas
		planets, err := filepath.Glob("planetas/*")
    	if err != nil {
        	log.Fatal(err)
    	}
    	var toSend2 []string
    	var toSend3 []string
    	var planetas []string
    	for i:= range planets {
    		planetas = append(planetas, strings.Split(strings.Split(planets[i],"\\")[1],".")[0])
    	}
    	for i:= range planetas{
    		consulta := "CLK "+planetas[i]
    		res2 := enviarMsg(ipFulcrum2, consulta)
    		res3 := enviarMsg(ipFulcrum3, consulta)
    		if(res2 == "Err404"){
    			//Fulcrum 2 no lo tiene, tendre que mandarle esa info.
    			toSend2 = append(toSend2, planetas[i])
    		}
    		if(res3 == "Err404"){
    			//Fulcrum 3 no lo tiene, tendre que mandarle esa info.
    			toSend3 = append(toSend3, planetas[i])
    		}
    		//Ambos fulcrum tienen el archivo, ahora debo saber si este fulcrum es el más actualizado para sino propagar a ambos.
    		if((res2 == res3) && (res2 != "Err404")){
    			//Son iguales, por lo que yo soy distinto a ellos y tengo el archivo, propago a los demás.
    			toSend2 = append(toSend2, planetas[i])
    			toSend3 = append(toSend3, planetas[i])
    		} else if ((res2 != res3) && (res2 != "Err404") && (res3 != "Err404")){
    			//Son distintos, por lo que yo estoy desactualizado, no se hace nada, asumiendo que otro fulcrum esta mas actualizado.
    		}
    		path:="log_planetas/"
    		err1:=os.RemoveAll(path)
    		if err1 != nil {
    			log.Fatal(err1)
    		}
    		err1 = os.Mkdir("log_planetas", 0700)
    		if err1 != nil {
    			log.Fatal(err1)
    		}
    	}
    	for i:= range toSend2{
    		//MERGEU,<planeta> (Resetea el archivo del fulcrum a consultar)
    		//MERGEUA,<planeta>,<linea a añadir>
    		//Se envia cada linea del archivo para que el fulcrum correspondiente actualice la info y cree las ciudades/planetas en caso de no tenerlas.
    		if _,err := os.Stat("planetas/"+toSend2[i]+".txt"); err == nil {
    			file, err := os.Open("planetas/"+toSend2[i]+".txt")
    			if err!=nil{
    				log.Fatalf("Fallo en abrir archivo")
    			}
    			scanner := bufio.NewScanner(file)
    			scanner.Split(bufio.ScanLines)
    			enviarMsg(ipFulcrum2, "MERGEU,"+toSend2[i])
    			for scanner.Scan() {
    				_ = enviarMsg(ipFulcrum2, "MERGEUA,"+toSend2[i]+","+scanner.Text())
    			}
    			file.Close()
    		}
    		//MERGECLK,<planeta>,<clock a updatear>
    		for j:=range clkPlanets{
    			if(clkPlanets[j].planeta == toSend2[i]){
    				_ = enviarMsg(ipFulcrum2, "MERGECLK,"+toSend2[i]+","+(strconv.Itoa(clkPlanets[j].relojx))+" "+(strconv.Itoa(clkPlanets[j].relojy))+" "+(strconv.Itoa(clkPlanets[j].relojz)))
    				break;
    			}
    		}
    	}
    	for i:= range toSend3{
    		//MERGEU,<planeta>,<linea a cambiar>
    		//Se envia cada linea del archivo para que el fulcrum correspondiente actualice la info y cree las ciudades/planetas en caso de no tenerlas.
    		if _,err := os.Stat("planetas/"+toSend3[i]+".txt"); err == nil {
    			file, err := os.Open("planetas/"+toSend3[i]+".txt")
    			if err!=nil{
    				log.Fatalf("Fallo en abrir archivo")
    			}
    			scanner := bufio.NewScanner(file)
    			scanner.Split(bufio.ScanLines)
    			enviarMsg(ipFulcrum3, "MERGEU,"+toSend3[i])
    			for scanner.Scan() {
    				_ = enviarMsg(ipFulcrum3, "MERGEUA,"+toSend3[i]+","+scanner.Text())
    			}
    			file.Close()
    		}
    		//MERGECLK,<planeta>,<clock a updatear>
    		for j:=range clkPlanets{
    			if(clkPlanets[j].planeta == toSend3[i]){
    				_ = enviarMsg(ipFulcrum3, "MERGECLK,"+toSend3[i]+","+(strconv.Itoa(clkPlanets[j].relojx))+" "+(strconv.Itoa(clkPlanets[j].relojy))+" "+(strconv.Itoa(clkPlanets[j].relojz)))
    				break;
    			}
    		}
    	}

	}

}



func main(){
	go timer()
	//Solo inicializo server, Funciones se encargan del resto 
	listener, err := net.Listen("tcp", ":50002")

	if err != nil {
		panic("No se puede crear la conexión tcp: "+ err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterEntradaMensajeServer(serv, &server{})
	if err = serv.Serve(listener); err != nil {
		panic("No se ha podido inicializar el servidor: "+ err.Error())
	}
	

}