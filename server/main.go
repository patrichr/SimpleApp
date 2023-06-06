package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/elsumanta/grpcserver/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	//Import internal modules
	handler "github.com/elsumanta/grpcserver/server/handler"
	repo "github.com/elsumanta/grpcserver/server/repo"
)

var (
	port     = flag.Int("port", 50051, "The server port")
	dbhost   = "192.168.0.108"
	dbport   = 5656
	user     = "server"
	password = "server"
	dbname   = "postgres"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//init database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbhost, user, password, dbname, dbport)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Init repo modules
	repo := repo.New(db)

	// Init handler module
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterGreeterServer(s, handler.Init(context.Background(), repo))

	// Serve port
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
