package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	pb "github.com/lpbayliss/tic-tac-toe/game"
	"google.golang.org/grpc"
)

const (
	port = ":4000"
)

type GameServer struct {
	pb.UnimplementedGameServer
}

func (s* GameServer) FindGame(context.Context, *pb.FindGameRequest) (*pb.FindGameResponse, error) {
	log.Print("Received FindGame Request");
	return &pb.FindGameResponse{Token: uuid.NewString()}, nil
}

func main() {
	fmt.Println("Starting Game Server!")

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGameServer(s, &GameServer{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}