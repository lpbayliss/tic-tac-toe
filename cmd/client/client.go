package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lpbayliss/tic-tac-toe/game"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:4000"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.FindGame(ctx, &pb.FindGameRequest{})

	log.Printf("Game Found: %v", r.GetToken())
}