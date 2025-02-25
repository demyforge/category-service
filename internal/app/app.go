package app

import (
	"github.com/demyforge/category-service/internal/api/handlers"
	"github.com/demyforge/category-service/internal/api/routes"
	"github.com/demyforge/category-service/internal/service"
	"github.com/demyforge/category-service/internal/storage"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type App struct {
	fiber *fiber.App
}

func New(storageDsn string) *App {
	app := &App{
		fiber: fiber.New(),
	}

	app.fiber.Use(cors.New())

	store, err := storage.New(storageDsn)
	if err != nil {
		panic(err)
	}

	s := service.New(store)
	h := handlers.New(s)

	routes.InitRoutes(app.fiber, h)

	return app
}

func (app *App) Listen(addr string) error {
	return app.fiber.Listen(addr)
}
