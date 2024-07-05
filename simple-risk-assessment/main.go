package main

import (
	"context"
	"gin/simple-risk-assessment/entity"
	"gin/simple-risk-assessment/handler"
	"gin/simple-risk-assessment/repository/postgres_gorm"
	"gin/simple-risk-assessment/repository/postgres_pgx"
	"gin/simple-risk-assessment/router"
	"gin/simple-risk-assessment/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	db, err := gorm.Open(postgres.Open("postgresql://postgres:postgres@localhost:5432/simple-risk-assessment"), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Submission{})
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := postgres_gorm.NewUserRepository(db)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	log.Println("Running server on port 8080")
	r.Run(":8080")
}

func connectDB(dbURL string) (postgres_pgx.PgxPoolIface, error) {
	return pgxpool.New(context.Background(), dbURL)
}
