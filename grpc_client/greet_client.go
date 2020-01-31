package main

import (
	"fmt"
	"log"

	"github.com/caiopapai/golang-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Iniciando um client GRPC")
	cc, err := grpc.Dial("localhost:5001", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Não foi possivel conectar no servidor")
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Cliente connectado", c)

}
