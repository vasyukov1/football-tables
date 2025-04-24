package postgres_repo

import (
	"context"
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
	"github.com/vasyukov1/football-tables/backend/internal/domain/repository"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/repository/postgres_repo/model"
	"gorm.io/gorm"
)

type MatchRepo struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) repository.MatchRepository {
	return &MatchRepo{db: db}
}

func (r *MatchRepo) Create(ctx context.Context, match *entity.Match) error {
	dbMatch := convertToDBModel(match)
	return r.db.WithContext(ctx).Create(dbMatch).Error
}

func (r *MatchRepo) GetByID(ctx context.Context, id int) (*entity.Match, error) {
	var dbMatch model.Match
	if err := r.db.WithContext(ctx).First(&dbMatch, id).Error; err != nil {
		return nil, err
	}
	return convertToEntity(&dbMatch), nil
}

func (r *MatchRepo) Update(ctx context.Context, match *entity.Match) error {
	// mock
	dbMatch := convertToDBModel(match)
	return r.db.WithContext(ctx).Create(dbMatch).Error
}

func (r *MatchRepo) Delete(ctx context.Context, id int) error {
	// mock
	var dbMatch model.Match
	return r.db.WithContext(ctx).First(&dbMatch, id).Error
}

func convertToDBModel(e *entity.Match) *model.Match {
	return &model.Match{
		ID:          e.ID,
		Team1ID:     e.Team1ID,
		Team2ID:     e.Team2ID,
		Score1:      e.Score1,
		Score2:      e.Score2,
		Stage:       e.Stage,
		NextMatchID: e.NextMatchID,
		IsCompleted: e.IsCompleted,
	}
}

func convertToEntity(m *model.Match) *entity.Match {
	return &entity.Match{
		ID:          m.ID,
		Team1ID:     m.Team1ID,
		Team2ID:     m.Team2ID,
		Score1:      m.Score1,
		Score2:      m.Score2,
		Stage:       m.Stage,
		NextMatchID: m.NextMatchID,
		IsCompleted: m.IsCompleted,
	}
}
