package main

import (
	"log"

	"github.com/JustSelection/users-service/internal/database"
	"github.com/JustSelection/users-service/internal/transport/grpc"
	"github.com/JustSelection/users-service/internal/user"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	repo := user.NewRepository(db)

	svc := user.NewService(repo)

	if err := grpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
