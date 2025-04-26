package repository

import (
	"context"
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
)

type TeamRepository interface {
	Create(ctx context.Context, match *entity.Team) error
	GetAll(ctx context.Context) ([]*entity.Team, error)
	GetByID(ctx context.Context, id int) (*entity.Team, error)
	GetByName(ctx context.Context, name string) (*entity.Team, error)
	Update(ctx context.Context, match *entity.Team) error
	Delete(ctx context.Context, id int) error

	GetMatchesByID(ctx context.Context, id int) ([]*entity.Match, error)
}
