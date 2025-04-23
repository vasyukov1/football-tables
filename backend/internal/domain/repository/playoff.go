package repository

import (
	"context"
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
)

type PlayoffRepository interface {
	Create(ctx context.Context, match *entity.Playoff) error
	GetByID(ctx context.Context, id int) (*entity.Playoff, error)
	Update(ctx context.Context, match *entity.Playoff) error
	Delete(ctx context.Context, id int) error
}
