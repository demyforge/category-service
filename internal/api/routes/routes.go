package routes

import (
	"github.com/demyforge/category-service/internal/api/handlers"
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(app fiber.Router, h *handlers.Handler) {
	app.Post("/category", h.CreateCategory)
	app.Get("/category", h.GetAllCategories)
}
