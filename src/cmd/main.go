package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/caarlos0/env/v6"
	pb "github.com/mehmetaligok/venom-example-project/src/proto/user"
	"github.com/mehmetaligok/venom-example-project/src/repository"
	"github.com/mehmetaligok/venom-example-project/src/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type appConfig struct {
	DSN  string `env:"DB_DSN"`
	Port string `env:"PORT"`
}

func main() {
	cfg := &appConfig{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatalf("failed to start grpc server. Error: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", cfg.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	userRepo := repository.NewUserRepo(cfg.DSN)
	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, server.NewUserServer(userRepo))
	reflection.Register(s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			s.GracefulStop()
		}
	}()

	log.Println("Starting server...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
