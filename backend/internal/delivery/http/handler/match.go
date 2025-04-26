package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/dto/request"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/dto/response"
	"github.com/vasyukov1/football-tables/backend/internal/usecase"
	"net/http"
)

type MatchHandler struct {
	matchUC *usecase.MatchUsecase
}

func NewMatchHandler(matchUC *usecase.MatchUsecase) *MatchHandler {
	return &MatchHandler{matchUC: matchUC}
}

// CreateMatch godoc
// @Summary Create match
// @Description Create new match
// @Tags matches
// @Accept json
// @Produce json
// @Param input body request.CreateMatchRequest true "Match data"
// Example: {"team1_id": 1, "team2_id": 2, "stage": "group", "group_id": 1, "playoff_id": 0}
// @Success 201 {object} entity.Match
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Router /matches [post]
func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var req request.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	match, err := h.matchUC.CreateMatch(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusConflict, &response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, match)
}

// GetMatches godoc
// @Summary Get all matches
// @Description Get list of all matches
// @Tags matches
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Match
// @Failure 500 {object} response.ErrorResponse
// @Router /matches [get]
func (h *MatchHandler) GetMatches(c *gin.Context) {
	matches, err := h.matchUC.GetMatches(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, matches)
}
