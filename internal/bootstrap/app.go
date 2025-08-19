package bootstrap

import (
	"database/sql"
	"log"

	"github.com/didinj/go-clean-architecture/internal/handler"
	"github.com/didinj/go-clean-architecture/internal/infrastructure"
	"github.com/didinj/go-clean-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
}

func InitializeApp() *App {
	dsn := "host=localhost port=5432 user=postgres password=yourpassword dbname=yourdb sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository (Postgres)
	userRepo := infrastructure.NewPostgresUserRepository(db)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Initialize usecase
	userUC := usecase.NewUserUsecase(userRepo)

	// Initialize handler
	userHandler := handler.NewUserHandler(userUC)

	// Setup Gin router
	r := gin.Default()
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)

	return &App{engine: r}
}

func (a *App) Run(addr string) error {
	return a.engine.Run(addr)
}
