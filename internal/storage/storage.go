package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/demyforge/category-service/internal/domain/models"
	"github.com/demyforge/category-service/pkg/database/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	Create(ctx context.Context, category *models.Category) error
	All(ctx context.Context) ([]*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*models.Category, error)
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

func (s *StorageImpl) Update(ctx context.Context, color *models.Category) error {
	query := "UPDATE categories SET name = :name WHERE id = :id"
	if _, err := s.db.NamedExecContext(ctx, query, color); err != nil {
		if postgres.IsDuplicate(err) {
			return errors.New("category already exists")
		}
		return err
	}
	return nil
}

func (s *StorageImpl) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := s.db.ExecContext(ctx, query, id)
	return err
}

func (s *StorageImpl) FindById(ctx context.Context, id uuid.UUID) (*models.Category, error) {
	var category models.Category
	query := "SELECT * FROM categories WHERE id = $1"
	if err := s.db.GetContext(ctx, &category, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}
