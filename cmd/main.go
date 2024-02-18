package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	config "github.com/gui-laranjeira/testyfy/configs"
	"github.com/gui-laranjeira/testyfy/internal/guests"
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
	guestHandler := guests.GuestFactory(db)

	app := fiber.New()

	app.Get("/api/v1/ping", func(c *fiber.Ctx) error { return c.SendString("pong") })
	app.Post("/api/v1/users/register", userHandler.CreateUser)
	app.Post("/api/v1/auth/login", userHandler.Authenticate)

	app.Post("/api/v1/guests", guestHandler.CreateGuest)
	app.Get("/api/v1/guests/:id", guestHandler.GetGuestById)
	app.Get("/api/v1/guests/user/:userId", guestHandler.GetGuestByUserId)
	app.Get("/api/v1/guests/email/:email", guestHandler.GetGuestByEmail)
	app.Put("/api/v1/guests/:id", guestHandler.UpdateGuest)
	app.Delete("/api/v1/guests/:id", guestHandler.DeleteGuest)
	
	app.Listen(":" + cfg.Server.Port)
}
