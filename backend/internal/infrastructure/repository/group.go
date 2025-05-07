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

// NewGroupRepository конструктор.
func NewGroupRepository(db *gorm.DB) repository.GroupRepository {
    return &GroupRepo{db: db}
}

// Create создаёт группу и привязывает к ней команды.
func (r *GroupRepo) Create(ctx context.Context, e *entity.Group) error {
    // 1) Создаём только саму группу
    m := &model.Group{
        Name: e.Name,
    }
    if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
        return err
    }

    // 2) Если есть команды — связываем их через many2many, не трогая таблицу teams
    if len(e.Teams) > 0 {
        var teamModels []*model.Team
        for _, te := range e.Teams {
            teamModels = append(teamModels, &model.Team{ID: te.ID})
        }
        if err := r.db.WithContext(ctx).
            Model(m).
            Association("Teams").
            Replace(teamModels); err != nil {
            return err
        }
    }
    return nil
}

// GetByID подгружает группу вместе с командами и матчами.
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

// Update обновляет группу и — если переданы — состав команд.
func (r *GroupRepo) Update(ctx context.Context, e *entity.Group) error {
    m := &model.Group{
        ID:   e.ID,
        Name: e.Name,
    }
    if err := r.db.WithContext(ctx).Save(m).Error; err != nil {
        return err
    }
    if e.Teams != nil {
        var teamModels []*model.Team
        for _, te := range e.Teams {
            teamModels = append(teamModels, &model.Team{ID: te.ID})
        }
        if err := r.db.WithContext(ctx).Model(m).Association("Teams").Replace(teamModels); err != nil {
            return err
        }
    }
    return nil
}

// Delete удаляет группу.
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

// GetAll возвращает все группы с командами и матчами.
func (r *GroupRepo) GetAll(ctx context.Context) ([]*entity.Group, error) {
    var ms []model.Group
    if err := r.db.WithContext(ctx).
        Preload("Teams").
        Preload("Matches").
        Find(&ms).Error; err != nil {
        return nil, err
    }
    var out []*entity.Group
    for i := range ms {
        out = append(out, converters.ConvertToEntityGroup(&ms[i]))
    }
    return out, nil
}

// GetTeamsByID отдаёт только команды группы.
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
    var teams []*entity.Team
    for _, t := range m.Teams {
        teams = append(teams, converters.ConvertToEntityTeam(t))
    }
    return teams, nil
}
