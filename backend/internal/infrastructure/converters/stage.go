package converters

import (
    "github.com/vasyukov1/football-tables/backend/internal/domain/entity"
    "github.com/vasyukov1/football-tables/backend/internal/infrastructure/model"
)

// ConvertToModelStage конвертирует бизнес‑сущность Stage в GORM‑модель.
func ConvertToModelStage(e *entity.Stage) *model.Stage {
    return &model.Stage{
        ID:        e.ID,
        PlayoffID: e.PlayoffID, // e.PlayoffID — int
        Name:      e.Name,
    }
}

// ConvertToEntityStage конвертирует GORM‑модель Stage в бизнес‑сущность.
func ConvertToEntityStage(m *model.Stage) *entity.Stage {
    return &entity.Stage{
        ID:        m.ID,
        PlayoffID: m.PlayoffID, // m.PlayoffID — int
        Name:      m.Name,
        Matches:   nil,         // если нужны матчи — грузить отдельно
    }
}
