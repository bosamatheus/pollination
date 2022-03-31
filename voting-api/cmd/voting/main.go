package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bosamatheus/pollination/voting-api/handler"
	fiber "github.com/gofiber/fiber/v2"
	fiberCORS "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	log "github.com/sirupsen/logrus"
)

func main() {
	// logger
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// app
	app := fiber.New()

	// middlewares
	app.Use(fiberCORS.New())
	app.Use(fiberRecover.New())
	app.Use(fiberLogger.New(fiberLogger.Config{
		TimeFormat: time.RFC3339,
	}))

	// routers
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Post("/api/v1/votes", handler.SendVote())

	// start server
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Panic(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)                    // create channel to signify a signal being sent
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // when an interrupt or termination signal is sent, notify the channel
	<-quit                                             // blocks the main thread until an interrupt is received
	log.Print("gracefully shutting down")

	err := app.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
