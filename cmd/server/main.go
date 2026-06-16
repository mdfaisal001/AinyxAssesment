package main

import (
	"log"

	"go-user-api/config"
	db "go-user-api/db/sqlc/generated"
	"go-user-api/internal/handler"
	"go-user-api/internal/logger"
	"go-user-api/internal/repository"
	"go-user-api/internal/routes"
	"go-user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	logger.InitLogger()

	conn, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	queries := db.New(conn)

	repo := repository.NewUserRepository(queries)

	userService := service.NewUserService(repo)

	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server Running")
	})

	routes.SetupRoutes(app, userHandler)

	log.Fatal(app.Listen(":3000"))
}