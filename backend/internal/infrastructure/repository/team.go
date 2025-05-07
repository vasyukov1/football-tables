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

type TeamRepo struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) repository.TeamRepository {
	return &TeamRepo{db: db}
}


func (r *TeamRepo) Create(ctx context.Context, team *entity.Team) error {
	return r.db.WithContext(ctx).Create(team).Error
}

func (r *TeamRepo) GetAll(ctx context.Context) ([]*entity.Team, error) {
	var dbTeams []*model.Team
	result := r.db.WithContext(ctx).Find(&dbTeams)
	if result.Error != nil {
		return nil, result.Error
	}

	teams := make([]*entity.Team, len(dbTeams))
	for i, t := range dbTeams {
		teams[i] = converters.ConvertToEntityTeam(t)
	}

	return teams, nil
}

func (r *TeamRepo) GetByID(ctx context.Context, id int) (*entity.Team, error) {
	var team model.Team
	if err := r.db.WithContext(ctx).First(&team, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return converters.ConvertToEntityTeam(&team), nil
}

func (r *TeamRepo) GetByName(ctx context.Context, name string) (*entity.Team, error) {
	var team model.Team
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return converters.ConvertToEntityTeam(&team), nil
}

// func (r *TeamRepo) Update(ctx context.Context, match *entity.Team) error {
// 	//TODO implement me
// 	panic("implement me")
// }

// func (r *TeamRepo) Delete(ctx context.Context, id int) error {
// 	//TODO implement me
// 	panic("implement me")
// }

// func (r *TeamRepo) GetMatchesByID(ctx context.Context, id int) ([]*entity.Match, error) {
// 	//TODO implement me
// 	panic("implement me")
// }

func (r *TeamRepo) Update(ctx context.Context, team *entity.Team) error {
    m := converters.ConvertToModelTeam(team)
    // Save обновит по первичному ключу, либо создаст, если нет
    return r.db.WithContext(ctx).Save(m).Error
}

func (r *TeamRepo) Delete(ctx context.Context, id int) error {
    res := r.db.WithContext(ctx).Delete(&model.Team{}, id)
    if res.Error != nil {
        return res.Error
    }
    if res.RowsAffected == 0 {
        return repository.ErrNotFound
    }
    return nil
}

func (r *TeamRepo) GetMatchesByID(ctx context.Context, id int) ([]*entity.Match, error) {
    var dbMatches []model.Match
    err := r.db.WithContext(ctx).
        Where("team1_id = ? OR team2_id = ?", id, id).
        Find(&dbMatches).Error
    if err != nil {
        return nil, err
    }

    out := make([]*entity.Match, len(dbMatches))
    for i, m := range dbMatches {
        out[i] = converters.ConvertToEntityMatch(&m)
    }
    return out, nil
}