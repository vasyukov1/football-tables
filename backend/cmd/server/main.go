package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vasyukov1/football-tables/backend/internal/config"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/handler"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/routes"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/repository"
	"github.com/vasyukov1/football-tables/backend/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// @title Football Tables API
// @version 1.0
// @description API for managing football tournaments
// @host localhost:8080
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}

	cfg := config.Load()

	db := initDB(cfg)

	matchRepo := repository.NewMatchRepository(db)
	teamRepo := repository.NewTeamRepository(db)

	matchUC := usecase.NewMatchUsecase(matchRepo, teamRepo)
	teamUC := usecase.NewTeamUsecase(teamRepo)

	matchHandler := handler.NewMatchHandler(matchUC)
	teamHandler := handler.NewTeamHandler(teamUC)

	router := routes.SetupAPIRouter(
		matchHandler,
		teamHandler,
		cfg,
	)

	log.Printf("Сервер запущен на порту %s", cfg.HTTP.Port)
	_ = router.Run(":" + cfg.HTTP.Port)
}

func initDB(cfg *config.Config) *gorm.DB {
	dsn := buildDSN(cfg)
	log.Println("DSN:", dsn)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Warn,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

func buildDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)
}
