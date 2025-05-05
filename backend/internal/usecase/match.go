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
    // Проверяем команды
    if _, err := uc.teamRepo.GetByID(ctx, req.Team1ID); err != nil {
        return nil, fmt.Errorf("team1 not found")
    }
    if _, err := uc.teamRepo.GetByID(ctx, req.Team2ID); err != nil {
        return nil, fmt.Errorf("team2 not found")
    }

    // Формируем сущность
    match := &entity.Match{
        Team1ID:   req.Team1ID,
        Team2ID:   req.Team2ID,
        Stage:     req.Stage,
        GroupID:   req.GroupID,   // уже *int
        PlayoffID: req.PlayoffID, // уже *int
    }

    // Сохраняем
    if err := uc.matchRepo.Create(ctx, match); err != nil {
        return nil, fmt.Errorf("failed to create match: %w", err)
    }
    return match, nil
}

func (uc *MatchUsecase) GetMatches(ctx context.Context) ([]*entity.Match, error) {
    return uc.matchRepo.GetAll(ctx)
}