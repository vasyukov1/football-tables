package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vasyukov1/football-tables/backend/internal/config"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/handler"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/middleware"
	_ "github.com/vasyukov1/football-tables/docs"
)

func SetupAPIRouter(
	matchHandler *handler.MatchHandler,
	teamHandler *handler.TeamHandler,
	cfg *config.Config,
) *gin.Engine {
	router := gin.New()

	if cfg.Env == "development" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.Use(
		middleware.CORS(cfg),
		gin.Recovery(),
	)

	public := router.Group("")
	{
		public.POST("/teams", teamHandler.CreateTeam)
		public.GET("/teams", teamHandler.GetTeams)
		public.POST("/matches", matchHandler.CreateMatch)
		public.GET("/matches", matchHandler.GetMatches)
	}

	// It will be for admins
	//
	//private := router.Group("/admin")
	//private.Use(middleware.JWT(cfg.JWT.Secret))
	//{
	//	private.POST("/matches", matchHandler.CreateMatch)
	//}

	return router
}
