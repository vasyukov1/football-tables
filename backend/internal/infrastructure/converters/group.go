package converters

import (
    "github.com/vasyukov1/football-tables/backend/internal/domain/entity"
    "github.com/vasyukov1/football-tables/backend/internal/infrastructure/model"
)

// ConvertToModelGroup конвертирует бизнес-сущность Group в GORM-модель.
// При создании/обновлении групп хранится только Name и ID.
func ConvertToModelGroup(e *entity.Group) *model.Group {
    return &model.Group{
        ID:   e.ID,
        Name: e.Name,
        // Teams и Matches подтягиваются через Preload при чтении
    }
}

// ConvertToEntityGroup конвертирует GORM-модель Group в бизнес-сущность,
// включая вложенные Teams и Matches.
func ConvertToEntityGroup(m *model.Group) *entity.Group {
    // конвертация команд
    var teams []*entity.Team
    for _, t := range m.Teams {
        teams = append(teams, &entity.Team{
            ID:   t.ID,
            Name: t.Name,
        })
    }
    // конвертация матчей
    var matches []*entity.Match
    for _, mm := range m.Matches {
        matches = append(matches, &entity.Match{
            ID:          mm.ID,
            Team1ID:     mm.Team1ID,
            Team2ID:     mm.Team2ID,
            Score1:      mm.Score1,
            Score2:      mm.Score2,
            Stage:       mm.Stage,
            GroupID:     mm.GroupID,
            PlayoffID:   mm.PlayoffID,
            NextMatchID: mm.NextMatchID,
            IsCompleted: mm.IsCompleted,
        })
    }
    return &entity.Group{
        ID:      m.ID,
        Name:    m.Name,
        Teams:   teams,
        Matches: matches,
    }
}
