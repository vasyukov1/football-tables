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

// GroupRepo — GORM-реализация repository.GroupRepository.
type GroupRepo struct {
    db *gorm.DB
}

// NewGroupRepository возвращает новый GroupRepository.
func NewGroupRepository(db *gorm.DB) repository.GroupRepository {
    return &GroupRepo{db: db}
}

func (r *GroupRepo) Create(ctx context.Context, e *entity.Group) error {
    m := converters.ConvertToModelGroup(e)
    return r.db.WithContext(ctx).Create(m).Error
}

func (r *GroupRepo) GetByID(ctx context.Context, id int) (*entity.Group, error) {
    var m model.Group
    err := r.db.WithContext(ctx).
        Preload("Teams").
        Preload("Matches").
        First(&m, id).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, repository.ErrNotFound
        }
        return nil, err
    }
    return converters.ConvertToEntityGroup(&m), nil
}

func (r *GroupRepo) Update(ctx context.Context, e *entity.Group) error {
    m := converters.ConvertToModelGroup(e)
    return r.db.WithContext(ctx).Save(m).Error
}

func (r *GroupRepo) Delete(ctx context.Context, id int) error {
    res := r.db.WithContext(ctx).Delete(&model.Group{}, id)
    if res.Error != nil {
        return res.Error
    }
    if res.RowsAffected == 0 {
        return repository.ErrNotFound
    }
    return nil
}

func (r *GroupRepo) GetTeamsByID(ctx context.Context, id int) ([]*entity.Team, error) {
    var m model.Group
    err := r.db.WithContext(ctx).
        Preload("Teams").
        First(&m, id).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, repository.ErrNotFound
        }
        return nil, err
    }
    // Конвертируем команды
    var out []*entity.Team
    for _, t := range m.Teams {
        out = append(out, converters.ConvertToEntityTeam(t))
    }
    return out, nil
}
