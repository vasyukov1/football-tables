package repository

import (
	"context"
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
)

type MatchRepository interface {
	Create(ctx context.Context, match *entity.Match) error
	GetByID(ctx context.Context, id int) (*entity.Match, error)
	GetAll(ctx context.Context) ([]*entity.Match, error)
	Update(ctx context.Context, match *entity.Match) error
	Delete(ctx context.Context, id int) error
}
