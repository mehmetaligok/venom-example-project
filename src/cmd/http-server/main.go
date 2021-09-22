package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/mehmetaligok/venom-example-project/src/repository"
	httpServer "github.com/mehmetaligok/venom-example-project/src/server/http"
)

type appConfig struct {
	DSN  string `env:"DB_DSN"`
	Port string `env:"PORT"`
}

func main() {
	cfg := &appConfig{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatalf("failed to start http server. Error: %v", err)
	}

	userRepo := repository.NewUserRepo(cfg.DSN)
	userServer := httpServer.NewUserServer(userRepo)
	server := gin.New()

	server.GET("/user/:id", userServer.GetUserHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), server))
}
