package app

import (
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/cors"
)

type App struct {
    fiber *fiber.App
}

func New() *App {
    app := &App{
        fiber: fiber.New(),
    }
    app.fiber.Use(cors.New())
    return app
}

func (app *App) Listen(addr string) error {
    return app.fiber.Listen(addr)
}
