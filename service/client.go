package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var connection *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %s", err)
	}
	defer conn.Close()

	s := service.NewServiceServiceClient(conn)
	message := service.Info{
		Name:   "Karan",
		School: "CSUEB",
		AGE:    "24",
		GPA:    "3.87",
	}

	response, err := s.InfoObtained(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling InfoObtained: %s", err)
	}
	log.Printf("Server: %s", response.Body)

}
