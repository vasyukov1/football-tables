package repository

import (
	"context"
	"errors"

	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
	"github.com/vasyukov1/football-tables/backend/internal/domain/repository"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/converters"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/model"
	"gorm.io/gorm"
)

// MatchRepo — GORM-реализация MatchRepository.
type MatchRepo struct {
    db *gorm.DB
}

// NewMatchRepository конструктор.
func NewMatchRepository(db *gorm.DB) repository.MatchRepository {
    return &MatchRepo{db: db}
}

func (r *MatchRepo) Create(ctx context.Context, e *entity.Match) error {
    m := converters.ConvertToModelMatch(e)
    return r.db.WithContext(ctx).Create(m).Error
}

func (r *MatchRepo) GetByID(ctx context.Context, id int) (*entity.Match, error) {
    var m model.Match
    if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, repository.ErrNotFound
        }
        return nil, err
    }
    return converters.ConvertToEntityMatch(&m), nil
}

func (r *MatchRepo) GetAll(ctx context.Context) ([]*entity.Match, error) {
    var arr []model.Match
    if err := r.db.WithContext(ctx).Find(&arr).Error; err != nil {
        return nil, err
    }
    out := make([]*entity.Match, len(arr))
    for i := range arr {
        out[i] = converters.ConvertToEntityMatch(&arr[i])
    }
    return out, nil
}

func (r *MatchRepo) Update(ctx context.Context, e *entity.Match) error {
    m := converters.ConvertToModelMatch(e)
    return r.db.WithContext(ctx).Save(m).Error
}

func (r *MatchRepo) Delete(ctx context.Context, id int) error {
    res := r.db.WithContext(ctx).Delete(&model.Match{}, id)
    if res.Error != nil {
        return res.Error
    }
    if res.RowsAffected == 0 {
        return repository.ErrNotFound
    }
    return nil
}

// func (r *MatchRepo) Update(ctx context.Context, match *entity.Match) error {
// 	// mock
// 	return r.db.WithContext(ctx).Create(match).Error
// }

// func (r *MatchRepo) Delete(ctx context.Context, id int) error {
// 	// mock
// 	var dbMatch model.Match
// 	return r.db.WithContext(ctx).First(&dbMatch, id).Error
// }



func (r *MatchRepo) TeamExists(ctx context.Context, teamID int) (bool, error) {
	var count int64
	err := r.db.Model(&model.Team{}).
		Where("id = ?", teamID).
		Count(&count).
		Error
	return count > 0, err
}
