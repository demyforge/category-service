package storage

import (
	"context"
	"errors"

	"github.com/demyforge/category-service/internal/domain/models"
	"github.com/demyforge/category-service/pkg/database/postgres"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	Create(ctx context.Context, category *models.Category) error
	All(ctx context.Context) ([]*models.Category, error)
}

type StorageImpl struct {
	db *sqlx.DB
}

func New(dsn string) (*StorageImpl, error) {
	db, err := postgres.Connect(dsn)
	if err != nil {
		return nil, err
	}

	return &StorageImpl{db: db}, nil
}

func (s *StorageImpl) Create(ctx context.Context, category *models.Category) error {
	query := "INSERT INTO categories (id, name) VALUES (:id, :name)"
	if _, err := s.db.NamedExecContext(ctx, query, category); err != nil {
		if postgres.IsDuplicate(err) {
			return errors.New("category already exists")
		}
		return err
	}
	return nil
}

func (s *StorageImpl) All(ctx context.Context) ([]*models.Category, error) {
	categories := []*models.Category{}
	query := "SELECT * FROM categories"
	if err := s.db.SelectContext(ctx, &categories, query); err != nil {
		return nil, err
	}
	return categories, nil
}
