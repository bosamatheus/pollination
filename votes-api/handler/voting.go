package handler

import (
	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func SendVote() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Print("sending vote")
		return c.Status(fiber.StatusAccepted).JSON("vote accepted")
	}
}
