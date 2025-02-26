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

type CategoryUpdateInput struct {
	ID   uuid.UUID
	Name string
}

type Service interface {
	Create(ctx context.Context, input CategoryCreateInput) (*models.Category, error)
	ById(ctx context.Context, id uuid.UUID) (*models.Category, error)
	All(ctx context.Context) ([]*models.Category, error)
	Update(ctx context.Context, input CategoryUpdateInput) (*models.Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
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

func (s *ServiceImpl) ById(ctx context.Context, id uuid.UUID) (*models.Category, error) {
	return s.storage.FindById(ctx, id)
}

func (s *ServiceImpl) All(ctx context.Context) ([]*models.Category, error) {
	return s.storage.All(ctx)
}

func (s *ServiceImpl) Update(ctx context.Context, input CategoryUpdateInput) (*models.Category, error) {
	color, err := s.ById(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	color.Name = input.Name

	if err = s.storage.Update(ctx, color); err != nil {
		return nil, err
	}

	return color, nil
}

func (s *ServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.ById(ctx, id); err != nil {
		return err
	}
	return s.storage.Delete(ctx, id)
}
