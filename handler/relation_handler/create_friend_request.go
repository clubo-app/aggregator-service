package relationhandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationGatewayHandler) CreateFriendRequest(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	fId := c.Params("id")

	if user.Sub == fId {
		return fiber.NewError(fiber.StatusBadRequest, "You can't add yourself")
	}

	ok, err := h.rc.CreateFriendRequest(c.Context(), &relation.CreateFriendRequestRequest{UserId: user.Sub, FriendId: fId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(ok)
}
