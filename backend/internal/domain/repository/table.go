package repository

import (
	"context"
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
)

type TableRepository interface {
	Create(ctx context.Context, match *entity.Table) error
	GetByID(ctx context.Context, id int) (*entity.Table, error)
	Update(ctx context.Context, match *entity.Table) error
	Delete(ctx context.Context, id int) error
}
