package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vasyukov1/football-tables/backend/internal/config"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/handler"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/middleware"
)

func SetupAPIRouter(
	matchHandler *handler.MatchHandler,
	cfg *config.Config,
) *gin.Engine {
	router := gin.New()

	router.Use(
		middleware.CORS(cfg),
		gin.Recovery(),
	)

	public := router.Group("/api")
	{
		public.GET("/matches", matchHandler.GetMatches)
	}

	private := router.Group("/api/admin")
	private.Use(middleware.JWT(cfg.JWT.Secret))
	{
		private.POST("/matches", matchHandler.CreateMatch)
	}

	return router
}
