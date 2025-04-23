package repository

import (
	"context"
	"football-tables/Backend/internal/domain/entity"
)

type TeamRepository interface {
	Create(ctx context.Context, match *entity.Team) error
	GetByID(ctx context.Context, id int) (*entity.Team, error)
	Update(ctx context.Context, match *entity.Team) error
	Delete(ctx context.Context, id int) error

	GetMatchesByID(ctx context.Context, id int) ([]*entity.Match, error)
}
