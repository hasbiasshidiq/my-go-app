package user

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func registerUserHandlers(router fiber.Router) {
	router.Get("/users/:id", userModuleInstance.GetUserByID)

}

func (um UserModule) GetUserByID(ctx *fiber.Ctx) error {

	idParam := ctx.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("invalid ID")
	}

	user, err := um.UserUsecase.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).SendString("user not found")
	}

	return ctx.Status(http.StatusOK).JSON(user)
}
