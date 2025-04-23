package repository

import (
	"context"
	"football-tables/Backend/internal/domain/entity"
)

type GroupRepository interface {
	Create(ctx context.Context, match *entity.Group) error
	GetByID(ctx context.Context, id int) (*entity.Group, error)
	Update(ctx context.Context, match *entity.Group) error
	Delete(ctx context.Context, id int) error

	GetTeamsByID(ctx context.Context, id int) ([]*entity.Team, error)
}
