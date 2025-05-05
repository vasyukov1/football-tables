package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
	"github.com/vasyukov1/football-tables/backend/internal/domain/repository"
)

type TeamUsecase struct {
	teamRepo repository.TeamRepository
}

func NewTeamUsecase(tr repository.TeamRepository) *TeamUsecase {
	return &TeamUsecase{teamRepo: tr}
}

func (uc *TeamUsecase) CreateTeam(ctx context.Context, name string) (*entity.Team, error) {
	existingTeam, err := uc.teamRepo.GetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing team: %w", err)
	}
	if existingTeam != nil {
		return nil, errors.New("team name already exists")
	}

	team := &entity.Team{Name: name}
	if err := uc.teamRepo.Create(ctx, team); err != nil {
		return nil, err
	}

	return team, nil
}

func (uc *TeamUsecase) GetTeams(ctx context.Context) ([]*entity.Team, error) {
	teams, err := uc.teamRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return teams, nil
}
