package profilehandler

import (
	"github.com/clubo-app/packages/utils"
	ug "github.com/clubo-app/protobuf/user"
	"github.com/gofiber/fiber/v2"

	"github.com/clubo-app/packages/utils/middleware"
)

func (h userGatewayHandler) GetMe(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	res, err := h.uc.GetUser(c.Context(), &ug.GetUserRequest{Id: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
