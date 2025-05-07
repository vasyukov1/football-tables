package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vasyukov1/football-tables/backend/internal/config"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/handler"
	_ "github.com/vasyukov1/football-tables/docs"
)

func SetupAPIRouter(
	matchHandler *handler.MatchHandler,
	teamHandler *handler.TeamHandler,
	groupHandler *handler.GroupHandler,
	cfg *config.Config,
) *gin.Engine {
	router := gin.Default()

	if cfg.Env == "development" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//router.Use(
	//	middleware.CORS(cfg),
	//	gin.Recovery(),
	//)

	router.Use(CORSMiddleware())

	public := router.Group("")
	{
		public.POST("/teams", teamHandler.CreateTeam)
		public.GET("/teams", teamHandler.GetTeams)
		public.POST("/matches", matchHandler.CreateMatch)
		public.GET("/matches", matchHandler.GetMatches)
		public.GET("/groups", groupHandler.List)
        public.POST("/groups", groupHandler.Create)
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}