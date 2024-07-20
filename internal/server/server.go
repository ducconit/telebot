package server

import (
	"github.com/gofiber/fiber/v2"

	"telebot/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "telebot",
			AppName:      "telebot",
		}),

		db: database.New(),
	}

	return server
}
