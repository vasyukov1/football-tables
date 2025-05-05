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

// StageRepo — GORM-реализация repository.StageRepository.
type StageRepo struct {
    db *gorm.DB
}

// NewStageRepository возвращает новый StageRepository.
func NewStageRepository(db *gorm.DB) repository.StageRepository {
    return &StageRepo{db: db}
}

func (r *StageRepo) Create(ctx context.Context, e *entity.Stage) error {
    m := converters.ConvertToModelStage(e)
    return r.db.WithContext(ctx).Create(m).Error
}

func (r *StageRepo) GetByID(ctx context.Context, id int) (*entity.Stage, error) {
    var m model.Stage
    err := r.db.WithContext(ctx).
        Preload("Matches").
        First(&m, id).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, repository.ErrNotFound
        }
        return nil, err
    }
    return converters.ConvertToEntityStage(&m), nil
}

func (r *StageRepo) Update(ctx context.Context, e *entity.Stage) error {
    m := converters.ConvertToModelStage(e)
    return r.db.WithContext(ctx).Save(m).Error
}

func (r *StageRepo) Delete(ctx context.Context, id int) error {
    res := r.db.WithContext(ctx).Delete(&model.Stage{}, id)
    if res.Error != nil {
        return res.Error
    }
    if res.RowsAffected == 0 {
        return repository.ErrNotFound
    }
    return nil
}

func (r *StageRepo) GetMatchesByID(ctx context.Context, id int) ([]*entity.Match, error) {
    var m model.Stage
    err := r.db.WithContext(ctx).
        Preload("Matches").
        First(&m, id).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, repository.ErrNotFound
        }
        return nil, err
    }
    // Конвертируем матчи
    var out []*entity.Match
    for _, mm := range m.Matches {
        out = append(out, converters.ConvertToEntityMatch(mm))
    }
    return out, nil
}
