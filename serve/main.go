package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hendrihmwn/dating-app-api/app/config"
	"github.com/hendrihmwn/dating-app-api/app/domain/repository/pg"
	"github.com/hendrihmwn/dating-app-api/app/domain/service"
	"github.com/hendrihmwn/dating-app-api/app/handler"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var (
	r *gin.Engine
)

func init() {
	db, err := sqlx.Open("postgres", config.GetDBConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	packageRepository := pg.NewPackageRepository(db)
	userRepository := pg.NewUserRepository(db)
	userPairRepository := pg.NewUserPairRepository(db)
	userPackageRepository := pg.NewUserPackageRepository(db)
	orderRepository := pg.NewOrderRepository(db)

	packageService := service.NewPackageService(packageRepository)
	userService := service.NewUserService(userRepository, userPairRepository, userPackageRepository)
	userPackageService := service.NewUserPackageService(userPackageRepository, packageRepository, orderRepository)

	userHandler := handler.NewUserHandler(userService)
	packageHandler := handler.NewPackageHandler(packageService)
	userPackageHandler := handler.NewUserPackageHandler(userPackageService)

	r = gin.Default()
	// Define your handler
	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Register)
	r.GET("/packages", packageHandler.List)
	r.GET("/package/:id", packageHandler.Get)
	r.Use(handler.TokenAuth).POST("/purchase", userPackageHandler.Create)
	r.Use(handler.TokenAuth).GET("/subscribes", userPackageHandler.List)
	r.Use(handler.TokenAuth).GET("/candidates", userHandler.List)
}

func run() error {
	// NOTE: Setup context, so the requets can be cancelled
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	return srv.ListenAndServe()
}

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
