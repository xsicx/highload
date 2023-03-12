package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/xsicx/highload/internal/interfaces/api"
	"github.com/xsicx/highload/internal/interfaces/config"
	"github.com/xsicx/highload/internal/interfaces/database"
)

func main() {
	cfg := config.Initialize()

	must(database.ConnectToPostgres(cfg.DB))
	defer database.Close()

	apiServer := fiber.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	serverShutdown := make(chan struct{})

	go func() {
		<-c

		log.Print("Gracefully shutting down...")

		shutdownErr := apiServer.Shutdown()

		must(shutdownErr)

		log.Print("API server stopped")

		serverShutdown <- struct{}{}
	}()

	api.New(apiServer, api.Dependencies{
		UsersGateway: database.NewUsersGateway(database.DB()),
	})

	must(apiServer.Listen(":9000"))

	<-serverShutdown
}

func must(err error) {
	if err == nil {
		return
	}

	log.Panic(err)
}
