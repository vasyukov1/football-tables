package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vasyukov1/football-tables/backend/internal/config"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/handler"
	"github.com/vasyukov1/football-tables/backend/internal/delivery/http/routes"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/repository/postgres_repo"
	"github.com/vasyukov1/football-tables/backend/internal/infrastructure/repository/postgres_repo/model"
	"github.com/vasyukov1/football-tables/backend/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Загрузка .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env not found, using environment variables")
	}

	// Загрузка конфига
	cfg := config.Load()

	// Инициализация БД
	db := initDB(cfg)

	// Инициализация репозиториев
	matchRepo := postgres_repo.NewMatchRepository(db)

	// Инициализация usecase
	matchUC := usecase.NewMatchUsecase(matchRepo)

	// Middleware
	//authMiddleware := middleware.NewAuthMiddleware(cfg.JWT.Secret)

	// Handlers
	matchHandler := handler.NewMatchHandler(matchUC)

	// Роутер
	router := routes.SetupAPIRouter(
		matchHandler,
		cfg,
	)

	// Запуск сервера
	log.Printf("🚀 Сервер запущен на порту %s", cfg.HTTP.Port)
	_ = router.Run(":" + cfg.HTTP.Port)
}

func initDB(cfg *config.Config) *gorm.DB {
	dsn := buildDSN(cfg)
	log.Println("DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Автомиграции (временное решение)
	if err := db.AutoMigrate(
		&model.Team{},
		&model.Group{},
		&model.Playoff{},
		&model.Stage{},
		&model.Match{},
		&model.Table{},
	); err != nil {
		log.Fatalf("Migration failed: %v", err)
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
