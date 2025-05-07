// backend/internal/usecase/group.go
package usecase

import (
    "context"
    "errors"
    "fmt"

    "github.com/vasyukov1/football-tables/backend/internal/domain/entity"
    "github.com/vasyukov1/football-tables/backend/internal/domain/repository"
)

// GroupUsecase хранит зависимости для работы с группами.
type GroupUsecase struct {
    groupRepo repository.GroupRepository
    teamRepo  repository.TeamRepository
}

// NewGroupUsecase создаёт новый GroupUsecase.
func NewGroupUsecase(
    gr repository.GroupRepository,
    tr repository.TeamRepository,
) *GroupUsecase {
    return &GroupUsecase{
        groupRepo: gr,
        teamRepo:  tr,
    }
}

// CreateGroup создаёт новую группу с указанным именем и списком teamIDs.
// Проверяет, что каждая команда существует.
func (uc *GroupUsecase) CreateGroup(ctx context.Context, name string, teamIDs []int) (*entity.Group, error) {
    // Проверяем команды
    var teams []*entity.Team
    for _, id := range teamIDs {
        t, err := uc.teamRepo.GetByID(ctx, id)
        if err != nil {
            return nil, fmt.Errorf("team %d not found", id)
        }
        teams = append(teams, t)
    }

    grp := &entity.Group{
        Name:  name,
        Teams: teams,
    }
    if err := uc.groupRepo.Create(ctx, grp); err != nil {
        return nil, fmt.Errorf("failed to create group: %w", err)
    }
    return grp, nil
}

// GetGroup возвращает группу по ID, включая список входящих в неё команд.
func (uc *GroupUsecase) GetGroup(ctx context.Context, id int) (*entity.Group, error) {
    grp, err := uc.groupRepo.GetByID(ctx, id)
    if err != nil {
        if errors.Is(err, repository.ErrNotFound) {
            return nil, fmt.Errorf("group %d not found", id)
        }
        return nil, err
    }
    return grp, nil
}

// UpdateGroup обновляет имя и состав группы.
// Для существующих групп также проверяет наличие команд.
func (uc *GroupUsecase) UpdateGroup(ctx context.Context, id int, name string, teamIDs []int) (*entity.Group, error) {
    // Сначала получаем существующую группу (для проверки)
    grp, err := uc.groupRepo.GetByID(ctx, id)
    if err != nil {
        if errors.Is(err, repository.ErrNotFound) {
            return nil, fmt.Errorf("group %d not found", id)
        }
        return nil, err
    }

    // Проверяем команды
    var teams []*entity.Team
    for _, tid := range teamIDs {
        t, err := uc.teamRepo.GetByID(ctx, tid)
        if err != nil {
            return nil, fmt.Errorf("team %d not found", tid)
        }
        teams = append(teams, t)
    }

    // Перезаполняем поля
    grp.Name = name
    grp.Teams = teams

    if err := uc.groupRepo.Update(ctx, grp); err != nil {
        return nil, fmt.Errorf("failed to update group: %w", err)
    }
    return grp, nil
}

// DeleteGroup удаляет группу по ID.
func (uc *GroupUsecase) DeleteGroup(ctx context.Context, id int) error {
    // Проверим существование
    if _, err := uc.groupRepo.GetByID(ctx, id); err != nil {
        if errors.Is(err, repository.ErrNotFound) {
            return fmt.Errorf("group %d not found", id)
        }
        return err
    }
    return uc.groupRepo.Delete(ctx, id)
}

// ListGroups возвращает все группы.
// Замечание: интерфейс GroupRepository должен иметь метод GetAll.
func (uc *GroupUsecase) ListGroups(ctx context.Context) ([]*entity.Group, error) {
    // Предполагаем, что метод GetAll реализован:
    groups, err := uc.groupRepo.GetAll(ctx)
    if err != nil {
        return nil, err
    }
    return groups, nil
}
