package converters

import (
    "github.com/vasyukov1/football-tables/backend/internal/domain/entity"
    "github.com/vasyukov1/football-tables/backend/internal/infrastructure/model"
)

// ConvertToModelPlayoff конвертирует бизнес‑сущность Playoff в GORM‑модель.
func ConvertToModelPlayoff(e *entity.Playoff) *model.Playoff {
    // Создаём слайс моделей Stage
    rounds := make([]*model.Stage, len(e.Rounds))
    for i, r := range e.Rounds {
        // у entity.Stage.PlayoffID — int, значит передаём прямо
        rounds[i] = &model.Stage{
            ID:        r.ID,
            PlayoffID: r.PlayoffID,
            Name:      r.Name,
        }
    }
    return &model.Playoff{
        ID:     e.ID,
        Rounds: rounds,
    }
}

// ConvertToEntityPlayoff конвертирует GORM‑модель Playoff в бизнес‑сущность.
func ConvertToEntityPlayoff(m *model.Playoff) *entity.Playoff {
    // У entity.Playoff.Rounds тип []entity.Stage (не указатели)
    rounds := make([]entity.Stage, len(m.Rounds))
    for i, r := range m.Rounds {
        rounds[i] = entity.Stage{
            ID:        r.ID,
            PlayoffID: r.PlayoffID, // r.PlayoffID — int
            Name:      r.Name,
        }
    }
    return &entity.Playoff{
        ID:     m.ID,
        Rounds: rounds,
    }
}
