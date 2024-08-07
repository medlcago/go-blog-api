package app

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/internal/app/config"
	"go-blog-api/internal/app/controllers"
	"go-blog-api/internal/app/db/postgres"
	"go-blog-api/internal/app/repositories"
	"go-blog-api/internal/app/routers"
	"go-blog-api/internal/app/services"
	"gorm.io/gorm"
	"log"
)

type App struct {
	r           *gin.Engine
	db          *gorm.DB
	repos       *repositories.Repository
	services    *services.Service
	controllers *controllers.Controller
}

func NewApp() *App {
	config.Load()

	app := new(App)
	app.r = gin.Default()

	// Dependency injection
	app.db = postgres.ConnectDB(config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)
	app.repos = repositories.NewRepository(app.db)
	app.services = services.NewServices(app.repos)
	app.controllers = controllers.NewController(app.services)

	return app
}

func (app *App) setupRouters() {
	app.r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Not found",
		})
	})

	router := app.r.Group("api/v1")
	{
		routers.SetupAuthRoutes(router, app.controllers.AuthController)
		routers.SetupUserRoutes(router, app.controllers.UsersController)
		routers.SetupPostsRoutes(router, app.controllers.PostsController)
	}
}

func (app *App) Run() {
	app.setupRouters()

	log.Fatal(app.r.Run(":8080"))
}
