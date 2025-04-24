package usecase

import (
	"context"
	"github.com/vasyukov1/football-tables/backend/internal/domain/repository"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/repository/postgres_repo/model"
)

type MatchUsecase struct {
	matchRepo repository.MatchRepository
}

func NewMatchUsecase(matchRepo repository.MatchRepository) *MatchUsecase {
	return &MatchUsecase{matchRepo: matchRepo}
}

func (u MatchUsecase) GetMatches(ctx context.Context) ([]*model.Match, error) {
	return []*model.Match{}, nil // mock
}
