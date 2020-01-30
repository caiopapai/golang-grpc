package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/caiopapai/golang-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {

	port := 50051

	//Listen announces on the local network address
	fmt.Println("Criando novo listner")
	lis, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("Fail to listen", err.Error())
	}
	fmt.Print("Listner criado")

	//NewServer creates a gRPC server which has no service registered and has not started to accept requests yet
	fmt.Println("Criando novo servidor")
	s := grpc.NewServer()
	fmt.Println("Registrando Service Server")
	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println("Subindo servidor")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Fail to serve: ", err.Error())
	}

}
