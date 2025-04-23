package repository

import (
	"context"
	"football-tables/Backend/internal/domain/entity"
)

type StageRepository interface {
	Create(ctx context.Context, match *entity.Stage) error
	GetByID(ctx context.Context, id int) (*entity.Stage, error)
	Update(ctx context.Context, match *entity.Stage) error
	Delete(ctx context.Context, id int) error

	GetMatchesByID(ctx context.Context, id int) ([]*entity.Match, error)
}
