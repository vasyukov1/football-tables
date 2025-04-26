package usecase

import (
	"context"
	"fmt"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/dto/request"
	"github.com/vasyukov1/football-tables/backend/internal/domain/entity"
	"github.com/vasyukov1/football-tables/backend/internal/domain/repository"
)

type MatchUsecase struct {
	matchRepo repository.MatchRepository
	teamRepo  repository.TeamRepository
}

func NewMatchUsecase(
	matchRepo repository.MatchRepository,
	teamRepo repository.TeamRepository,
) *MatchUsecase {
	return &MatchUsecase{
		matchRepo: matchRepo,
		teamRepo:  teamRepo,
	}
}

func (uc *MatchUsecase) CreateMatch(ctx context.Context, req *request.CreateMatchRequest) (*entity.Match, error) {
	_, err := uc.teamRepo.GetByID(ctx, req.Team1ID)
	if err != nil {
		return nil, fmt.Errorf("team1 not found")
	}
	_, err = uc.teamRepo.GetByID(ctx, req.Team2ID)
	if err != nil {
		return nil, fmt.Errorf("team2 not found")
	}

	match := &entity.Match{
		Team1ID:   req.Team1ID,
		Team2ID:   req.Team2ID,
		Stage:     req.Stage,
		GroupID:   req.GroupID,
		PlayoffID: req.PlayoffID,
	}

	if err := uc.matchRepo.Create(ctx, match); err != nil {
		return nil, fmt.Errorf("failed to create match: %w", err)
	}

	return match, nil
}

func (uc *MatchUsecase) GetMatches(ctx context.Context) ([]*entity.Match, error) {
	matches, err := uc.matchRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return matches, nil
}
