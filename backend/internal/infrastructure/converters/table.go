package converters

import (
    "github.com/vasyukov1/football-tables/backend/internal/domain/entity"
    "github.com/vasyukov1/football-tables/backend/internal/infrastructure/model"
)

func ConvertToModelTable(e *entity.Table) *model.Table {
    return &model.Table{
        ID:       e.ID,
        GroupIDs: e.GroupIDs,
    }
}

func ConvertToEntityTable(m *model.Table) *entity.Table {
    return &entity.Table{
        ID:       m.ID,
        GroupIDs: m.GroupIDs,
    }
}
