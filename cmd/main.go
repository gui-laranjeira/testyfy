package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	config "github.com/gui-laranjeira/testyfy/configs"
	"github.com/gui-laranjeira/testyfy/internal/user"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", cfg.Database.User, cfg.Database.Password,
		cfg.Database.Container, cfg.Database.Port, cfg.Database.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	userHandler := user.UserFactory(db)

	app := fiber.New()

	app.Get("/api/v1/ping", func(c *fiber.Ctx) error { return c.SendString("pong") })
	app.Post("/api/v1/users/register", userHandler.CreateUser)
	app.Post("/api/v1/auth/login", userHandler.Authenticate)

	app.Listen(":" + cfg.Server.Port)
}
