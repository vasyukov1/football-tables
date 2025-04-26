package repository

import (
	"context"
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
	"github.com/vasyukov1/football-tables/backend/internal/domain/repository"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/converters"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/model"
	"gorm.io/gorm"
)

type MatchRepo struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) repository.MatchRepository {
	return &MatchRepo{db: db}
}

func (r *MatchRepo) Create(ctx context.Context, match *entity.Match) error {
	return r.db.WithContext(ctx).Create(match).Error
}

func (r *MatchRepo) GetByID(ctx context.Context, id int) (*entity.Match, error) {
	var dbMatch model.Match
	if err := r.db.WithContext(ctx).First(&dbMatch, id).Error; err != nil {
		return nil, err
	}
	return converters.ConvertToEntityMatch(&dbMatch), nil
}

func (r *MatchRepo) GetAll(ctx context.Context) ([]*entity.Match, error) {
	var dbMatches []*model.Match
	result := r.db.WithContext(ctx).Find(&dbMatches)
	if result.Error != nil {
		return nil, result.Error
	}

	matches := make([]*entity.Match, len(dbMatches))
	for i, m := range dbMatches {
		matches[i] = converters.ConvertToEntityMatch(m)
	}

	return matches, nil
}

func (r *MatchRepo) Update(ctx context.Context, match *entity.Match) error {
	// mock
	return r.db.WithContext(ctx).Create(match).Error
}

func (r *MatchRepo) Delete(ctx context.Context, id int) error {
	// mock
	var dbMatch model.Match
	return r.db.WithContext(ctx).First(&dbMatch, id).Error
}

func (r *MatchRepo) TeamExists(ctx context.Context, teamID int) (bool, error) {
	var count int64
	err := r.db.Model(&model.Team{}).
		Where("id = ?", teamID).
		Count(&count).
		Error
	return count > 0, err
}
