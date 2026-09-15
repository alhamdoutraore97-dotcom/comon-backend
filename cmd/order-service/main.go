package main

import (
	_ "github.com/alhamdoutraore97-dotcom/backend/cmd/order-service/docs"
	"github.com/alhamdoutraore97-dotcom/backend/internal/order/delivery/http"
	"github.com/alhamdoutraore97-dotcom/backend/internal/order/repository"
	"github.com/alhamdoutraore97-dotcom/backend/internal/order/usecase"
	"github.com/alhamdoutraore97-dotcom/comon-backend/db"
	"github.com/alhamdoutraore97-dotcom/comon-backend/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Order Service API
// @version 1.0
// @description Service de gestion des commandes / Сервис управления заказами
// @host localhost:8082
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

	// Initialisation des couches / Инициализация слоёв
	orderRepo := repository.NewOrderPostgresRepo(database)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	orderHandler := http.NewOrderHandler(orderUsecase)

	// Routeur Gin / Маршрутизатор Gin
	r := gin.Default()

	// Routes / Маршруты
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.POST("/orders", orderHandler.CreateOrder)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Info().Msg("Order-service démarré sur le port 8082 / Order-service запущен на порту 8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal().Err(err).Msg("Échec du lancement du serveur / Не удалось запустить сервер")
	}
}
