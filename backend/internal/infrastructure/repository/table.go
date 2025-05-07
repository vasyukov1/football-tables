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

// TableRepo — GORM-реализация repository.TableRepository.
type TableRepo struct {
    db *gorm.DB
}

// NewTableRepository возвращает новый TableRepository.
func NewTableRepository(db *gorm.DB) repository.TableRepository {
    return &TableRepo{db: db}
}

func (r *TableRepo) Create(ctx context.Context, e *entity.Table) error {
    m := converters.ConvertToModelTable(e)
    return r.db.WithContext(ctx).Create(m).Error
}

func (r *TableRepo) GetByID(ctx context.Context, id int) (*entity.Table, error) {
    var m model.Table
    err := r.db.WithContext(ctx).First(&m, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, repository.ErrNotFound
        }
        return nil, err
    }
    return converters.ConvertToEntityTable(&m), nil
}

func (r *TableRepo) Update(ctx context.Context, e *entity.Table) error {
    m := converters.ConvertToModelTable(e)
    return r.db.WithContext(ctx).Save(m).Error
}

func (r *TableRepo) Delete(ctx context.Context, id int) error {
    res := r.db.WithContext(ctx).Delete(&model.Table{}, id)
    if res.Error != nil {
        return res.Error
    }
    if res.RowsAffected == 0 {
        return repository.ErrNotFound
    }
    return nil
}
