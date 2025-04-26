package converters

import (
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/model"
)

func ConvertToDBMatch(e *entity.Match) *model.Match {
	return &model.Match{
		Team1ID:   e.Team1ID,
		Team2ID:   e.Team2ID,
		Stage:     e.Stage,
		GroupID:   e.GroupID,
		PlayoffID: e.PlayoffID,
	}
}

func ConvertToEntityMatch(m *model.Match) *entity.Match {
	return &entity.Match{
		ID:          m.ID,
		Team1ID:     m.Team1ID,
		Team2ID:     m.Team2ID,
		Score1:      m.Score1,
		Score2:      m.Score2,
		Stage:       m.Stage,
		GroupID:     m.GroupID,
		PlayoffID:   m.PlayoffID,
		IsCompleted: m.IsCompleted,
	}
}
