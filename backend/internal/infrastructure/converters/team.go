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


func ConvertToModelTeam(e *entity.Team) *model.Team {
    return &model.Team{
        ID:   e.ID,   // ID можно оставить 0 для Create, либо передать существующий для Update
        Name: e.Name,
    }
}
