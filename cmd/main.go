package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/lbayliss/tic-tac-toe/proto"
	"google.golang.org/grpc"
)

const (
	port = ":4000"
)

type GameServer struct {

}

func main() {
	fmt.Println("Hello World!")
}