package main

import (
    "fmt"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/leonidlivshits/football-tables/Backend/internal/repository/postgres/model"
)

func main() {
    // Попытка загрузить .env (если файла нет — просто логируем предупреждение)
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  .env not found, using environment variables")
    }

    // Собираем DSN из переменных окружения
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )
    log.Printf("Connecting to %s:%s as %s/%s",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"), os.Getenv("DB_NAME"),
    )

    // Открываем GORM
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Автомиграции
    if err := db.AutoMigrate(
		&model.Team{},
		&model.Group{},
		&model.Playoff{},
		&model.Stage{},   // сначала stages
		&model.Match{},   // потом matches
		&model.Table{},
    ); err != nil {
        log.Fatalf("Migration failed: %v", err)
    }
    log.Println("✅ Миграции выполнены успешно")

    // Запуск сервера
    router := gin.Default()
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok", "db": "connected"})
    })
    log.Println("🚀 Сервер запущен на порту :8080")
    router.Run(":8080")
}
