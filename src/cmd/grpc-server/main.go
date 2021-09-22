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
	grpcServer "github.com/mehmetaligok/venom-example-project/src/server/grpc"
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	userRepo := repository.NewUserRepo(cfg.DSN)
	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, grpcServer.NewUserServer(userRepo))
	reflection.Register(s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")

			s.GracefulStop()
		}
	}()

	log.Println("Starting server...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
