package handler

import (
	"github.com/bosamatheus/pollination/vote-api/internal/api/presenter"
	"github.com/bosamatheus/pollination/vote-api/internal/usecase/vote"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// SendVote sends vote to be processed.
func SendVote(service vote.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		vote := &presenter.Vote{}
		if err := c.BodyParser(&vote); err != nil {
			log.Errorln(err)
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		err := service.Send(vote.MapToEntity())
		if err != nil {
			log.Errorln(err)
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		return c.Status(fiber.StatusAccepted).JSON("vote accepted")
	}
}
