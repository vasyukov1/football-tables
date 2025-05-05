// backend/internal/infrastructure/repository/playoff_repo.go
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

type PlayoffRepo struct {
    db *gorm.DB
}

func NewPlayoffRepository(db *gorm.DB) repository.PlayoffRepository {
    return &PlayoffRepo{db: db}
}

func (r *PlayoffRepo) Create(ctx context.Context, p *entity.Playoff) error {
    m := converters.ConvertToModelPlayoff(p)
    return r.db.WithContext(ctx).Create(m).Error
}

func (r *PlayoffRepo) GetByID(ctx context.Context, id int) (*entity.Playoff, error) {
    var m model.Playoff
    if err := r.db.WithContext(ctx).
        Preload("Rounds").
        First(&m, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, repository.ErrNotFound
        }
        return nil, err
    }
    return converters.ConvertToEntityPlayoff(&m), nil
}

func (r *PlayoffRepo) Update(ctx context.Context, p *entity.Playoff) error {
    m := converters.ConvertToModelPlayoff(p)
    return r.db.WithContext(ctx).Save(m).Error
}

func (r *PlayoffRepo) Delete(ctx context.Context, id int) error {
    res := r.db.WithContext(ctx).Delete(&model.Playoff{}, id)
    if res.Error != nil {
        return res.Error
    }
    if res.RowsAffected == 0 {
        return repository.ErrNotFound
    }
    return nil
}
