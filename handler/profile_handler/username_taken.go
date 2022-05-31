package profilehandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type UsernameTakenResponse struct {
	Taken bool `json:"taken"`
}

func (h profileGatewayHandler) UsernameTaken(c *fiber.Ctx) error {
	uName := c.Params("username")

	res, err := h.pc.UsernameTaken(c.Context(), &profile.UsernameTakenRequest{Username: uName})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(&UsernameTakenResponse{Taken: res.Taken})
}
