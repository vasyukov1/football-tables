package repository

import (
	"context"
	"football-tables/Backend/internal/domain/entity"
)

type MatchRepository interface {
	Create(ctx context.Context, match *entity.Match) error
	GetByID(ctx context.Context, id int) (*entity.Match, error)
	Update(ctx context.Context, match *entity.Match) error
	Delete(ctx context.Context, id int) error
}
