package relationhandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationGatewayHandler) AcceptFriendRequest(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	uId := c.Params("id")

	if user.Sub == uId {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Friend Id")
	}

	ok, err := h.rc.AcceptFriendRequest(c.Context(), &relation.AcceptFriendRequestRequest{UserId: uId, FriendId: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	return c.Status(fiber.StatusOK).JSON(ok)
}
