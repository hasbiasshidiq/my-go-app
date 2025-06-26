package user

import (
	"my-go-app/module/user/interfaces"
	repository "my-go-app/module/user/repository/postgres"
	"my-go-app/module/user/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserModule struct {
	UserUsecase interfaces.UserUsecase
}

var userModuleInstance UserModule
var PayrollUsecase usecase.UserUsecase

func InitModule() {
	app := fiber.New()

	// create repository
	userRepo := repository.NewPostgresUserRepo()

	// create usecase
	userUC := usecase.NewUserUsecase(userRepo)

	userModuleInstance = UserModule{
		UserUsecase: userUC,
	}

	registerUserHandlers(app)

	app.Listen(":8080")

}
