package converters

import (
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/model"
)

func ConvertToModelMatch(e *entity.Match) *model.Match {
    return &model.Match{
        ID:           e.ID,
        Team1ID:      e.Team1ID,
        Team2ID:      e.Team2ID,
        Score1:       e.Score1,
        Score2:       e.Score2,
        Stage:        e.Stage,
        GroupID:      e.GroupID,      // *int
        PlayoffID:    e.PlayoffID,    // *int
        NextMatchID:  e.NextMatchID,  // *int
        IsCompleted:  e.IsCompleted,
    }
}

// ConvertToEntityMatch конвертирует GORM-модель в бизнес-сущность.
func ConvertToEntityMatch(m *model.Match) *entity.Match {
    return &entity.Match{
        ID:           m.ID,
        Team1ID:      m.Team1ID,
        Team2ID:      m.Team2ID,
        Score1:       m.Score1,
        Score2:       m.Score2,
        Stage:        m.Stage,
        GroupID:      m.GroupID,      // *int
        PlayoffID:    m.PlayoffID,    // *int
        NextMatchID:  m.NextMatchID,  // *int
        IsCompleted:  m.IsCompleted,
    }
}