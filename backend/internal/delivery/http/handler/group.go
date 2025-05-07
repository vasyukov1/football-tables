// backend/internal/delivery/http/handler/group_handler.go
package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/vasyukov1/football-tables/backend/internal/usecase"
)

type GroupHandler struct {
    uc *usecase.GroupUsecase
}

func NewGroupHandler(uc *usecase.GroupUsecase) *GroupHandler {
    return &GroupHandler{uc: uc}
}

func (h *GroupHandler) Create(c *gin.Context) {
    var req struct {
        Name    string `json:"name" binding:"required"`
        TeamIDs []int  `json:"teamIds"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    grp, err := h.uc.CreateGroup(c.Request.Context(), req.Name, req.TeamIDs)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, grp)
}

func (h *GroupHandler) List(c *gin.Context) {
    groups, err := h.uc.ListGroups(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, groups)
}