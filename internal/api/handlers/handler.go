package handlers

import (
	"errors"

	"github.com/demyforge/category-service/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Handler struct {
	service service.Service
}

func New(s service.Service) *Handler {
	return &Handler{service: s}
}

type categoryCreateInput struct {
	Name string `json:"name"`
}

func (h *Handler) CreateCategory(c fiber.Ctx) error {
	var input categoryCreateInput
	if err := c.Bind().JSON(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	category, err := h.service.Create(c.Context(), service.CategoryCreateInput{
		Name: input.Name,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

func (h *Handler) GetAllCategories(c fiber.Ctx) error {
	categories, err := h.service.All(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(categories)
}

func parseUUID(uid string) (uuid.UUID, error) {
	id, err := uuid.Parse(uid)
	if err != nil {
		return uuid.Nil, errors.New("invalid uuid")
	}
	return id, nil
}
