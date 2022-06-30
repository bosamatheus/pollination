package handler

import (
	"github.com/bosamatheus/pollination/voting-api/internal/api/presenter"
	"github.com/bosamatheus/pollination/voting-api/internal/usecase/voting"
	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

// SendVote sends vote to be processed.
func SendVote(service voting.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		vote := &presenter.Vote{}
		if err := c.BodyParser(&vote); err != nil {
			log.Errorln(err)
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		err := service.SendVote(vote.MapToDomain())
		if err != nil {
			log.Errorln(err)
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		return c.Status(fiber.StatusAccepted).JSON("vote accepted")
	}
}
