package repository

import (
	"context"
	"football-tables/Backend/internal/domain/entity"
)

type PlayoffRepository interface {
	Create(ctx context.Context, match *entity.Playoff) error
	GetByID(ctx context.Context, id int) (*entity.Playoff, error)
	Update(ctx context.Context, match *entity.Playoff) error
	Delete(ctx context.Context, id int) error
}
