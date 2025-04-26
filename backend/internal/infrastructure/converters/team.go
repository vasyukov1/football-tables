package converters

import (
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/model"
)

func ConvertToEntityTeam(t *model.Team) *entity.Team {
	return &entity.Team{
		ID:   t.ID,
		Name: t.Name,
	}
}
