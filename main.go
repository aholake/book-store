package main

import (
	"log"
	"net"

	"github.com/aholake/book-store/database"
	pb "github.com/aholake/book-store/proto"
	"github.com/aholake/book-store/server"
	"google.golang.org/grpc"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Unable to create listener on port 8089: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, server.Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Unable to start server %v", err)
	}
}
