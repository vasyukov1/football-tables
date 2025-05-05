package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/dto/request"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/dto/response"
	"github.com/vasyukov1/football-tables/backend/internal/usecase"
	"net/http"
)

type TeamHandler struct {
	teamUC *usecase.TeamUsecase
}

func NewTeamHandler(tu *usecase.TeamUsecase) *TeamHandler {
	return &TeamHandler{teamUC: tu}
}

// CreateTeam godoc
// @Summary Create new team
// @Description Create new team
// @Tags teams
// @Accept json
// @Produce json
// @Param input body request.CreateTeamRequest true "Team data"
// Example: {"name": "Barcelona"}
// @Success 201 {object} entity.Team
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Router /teams [post]
func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var req request.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &response.ErrorResponse{
			Message: "Invalid request body format",
		})
		return
	}

	team, err := h.teamUC.CreateTeam(c.Request.Context(), req.Name)
	if err != nil {
		c.JSON(http.StatusConflict, &response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, team)
}

// GetTeams godoc
// @Summary Get all teams
// @Description Get list of all teams
// @Tags teams
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Team
// @Failure 500 {object} response.ErrorResponse
// @Router /teams [get]
func (h *TeamHandler) GetTeams(c *gin.Context) {
	teams, err := h.teamUC.GetTeams(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, teams)
}
