package app

import (
	"net/http"

	"github.com/Vasyl-Trefilov/bestBack/internal/handler"
	"github.com/Vasyl-Trefilov/bestBack/internal/repository"
	"github.com/Vasyl-Trefilov/bestBack/internal/service"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type App struct {
	handler *handler.Handler
}

func NewApp(configDB string) (*App, error) {
	db, err := repository.NewPostgresDB(configDB)
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	handler := handler.NewHandler(srv)

	return &App{
		handler: handler,
	}, nil
}

func (a *App) Run(port string) error {
	err := http.ListenAndServe(":"+port, a.routes())
	if err != nil {
		return err
	}
	return nil
}

func (a *App) routes() http.Handler {
	r := gin.Default()
	r.GET("/hello-world", a.handler.HelloWorld)
	r.GET("health", a.handler.HealthCheck)
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/list", a.handler.ListUsers)
		userRoutes.POST("/create", a.handler.CreateUser)
		userRoutes.GET("/get/:id", a.handler.GetUser)
		userRoutes.PUT("/update/:id", a.handler.UpdateUser)
		userRoutes.DELETE("/delete/:id", a.handler.DeleteUser)
	}
	// userRoutes.Use(AuthMiddleware())
	return r
}
