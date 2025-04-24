package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vasyukov1/football-tables/backend/internal/usecase"
	"net/http"
)

type MatchHandler struct {
	matchUC *usecase.MatchUsecase
}

func NewMatchHandler(matchUC *usecase.MatchUsecase) *MatchHandler {
	return &MatchHandler{matchUC: matchUC}
}

func (h *MatchHandler) CreateMatch(c *gin.Context) {

}

func (h *MatchHandler) GetMatches(c *gin.Context) {
	matches, err := h.matchUC.GetMatches(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, matches)
}
