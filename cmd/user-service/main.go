package main

import (
	_ "github.com/alhamdoutraore97-dotcom/backend/cmd/user-service/docs" // swagger docs
	"github.com/alhamdoutraore97-dotcom/backend/internal/user/delivery/http"
	"github.com/alhamdoutraore97-dotcom/backend/internal/user/repository"
	"github.com/alhamdoutraore97-dotcom/backend/internal/user/usecase"
	"github.com/alhamdoutraore97-dotcom/comon-backend/db"
	"github.com/alhamdoutraore97-dotcom/comon-backend/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Service API
// @version 1.0
// @description Service de gestion des utilisateurs / Сервис управления пользователями
// @host localhost:8081
// @BasePath /
func main() {
	// Initialisation du logger / Инициализация логгера
	logger.Init()
	log := logger.Get()

	// Connexion à PostgreSQL / Подключение к PostgreSQL
	connStr := "postgresql://postgres:1234@localhost:5432/postgres?sslmode=disable"
	database, err := db.Connect(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Impossible de se connecter à PostgreSQL / Не удалось подключиться к PostgreSQL")
	}
	defer database.Close()

	// Vérification de la connexion / Проверка подключения
	if err := database.Ping(); err != nil {
		log.Fatal().Err(err).Msg("PostgreSQL ne répond pas / PostgreSQL не отвечает")
	}
	log.Info().Msg("Connecté à PostgreSQL / Подключено к PostgreSQL")

	// Initialisation des couches (Clean Architecture) / Инициализация слоёв
	userRepo := repository.NewUserPostgresRepo(database)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := http.NewUserHandler(userUsecase)

	// Routeur Gin / Маршрутизатор Gin
	r := gin.Default()

	// Routes / Маршруты
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Info().Msg("User-service démarré sur le port 8081 / User-service запущен на порту 8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal().Err(err).Msg("Échec du lancement du serveur / Не удалось запустить сервер")
	}
}
