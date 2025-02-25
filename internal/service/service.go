package service

import (
	"context"

	"github.com/demyforge/category-service/internal/domain/models"
	"github.com/demyforge/category-service/internal/storage"
	"github.com/google/uuid"
)

type CategoryCreateInput struct {
	Name string
}

type Service interface {
	Create(ctx context.Context, input CategoryCreateInput) (*models.Category, error)
	All(ctx context.Context) ([]*models.Category, error)
}

type ServiceImpl struct {
	storage storage.Storage
}

func New(storage storage.Storage) *ServiceImpl {
	return &ServiceImpl{storage: storage}
}

func (s *ServiceImpl) Create(ctx context.Context, input CategoryCreateInput) (*models.Category, error) {
	category := &models.Category{
		ID:   uuid.New(),
		Name: input.Name,
	}

	if err := s.storage.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *ServiceImpl) All(ctx context.Context) ([]*models.Category, error) {
	return s.storage.All(ctx)
}
