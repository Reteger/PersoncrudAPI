package app

import (
	"log"
	"personcrud/internal/config"
	"personcrud/internal/database"
	"personcrud/internal/handlers"
	"personcrud/internal/repository"
	"personcrud/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type App struct {
	db     *bun.DB
	router *gin.Engine
	config *config.Config
}

func New() (*App, error) {
	// Загружаем конфигурацию
	cfg := config.Load()

	// Подключаемся к базе с использованием конфига
	db, err := database.Connect(cfg)
	if err != nil {
		return nil, err
	}

	// Инициализируем слои
	personRepo := repository.NewPersonRepository(db)
	personService := services.NewPersonService(personRepo)
	personHandler := handlers.NewPersonHandler(personService)

	// Настраиваем роутер
	router := gin.Default()

	// Регистрируем маршруты
	router.GET("/persons", personHandler.GetAll)
	router.GET("/persons/:id", personHandler.GetByID)
	router.POST("/persons", personHandler.Create)
	router.PUT("/persons/:id", personHandler.Update)
	router.DELETE("/persons/:id", personHandler.Delete)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return &App{
		db:     db,
		router: router,
		config: cfg,
	}, nil
}

func (a *App) Run() error {
	log.Printf("🚀 Server starting on :%s", a.config.ServerPort)
	return a.router.Run(":" + a.config.ServerPort)
}

func (a *App) Close() {
	// Обрабатываем ошибку закрытия базы данных
	if err := a.db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	} else {
		log.Println("Database connection closed")
	}
}
