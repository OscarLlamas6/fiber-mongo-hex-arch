package ports

import (
	fiber "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	Login(email string, password string) error
	Register(email string, password string, name string) (*mongo.InsertOneResult, error)
}

type IUserRepository interface {
	Login(email string, password string) error
	Register(email string, password string, name string) (*mongo.InsertOneResult, error)
}

type IUserHandlers interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
