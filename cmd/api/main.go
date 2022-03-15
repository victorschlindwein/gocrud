package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/victorschlindwein/gocrud/src/handlers"
	"github.com/victorschlindwein/gocrud/src/repositories"
	"github.com/victorschlindwein/gocrud/src/services"
	"github.com/victorschlindwein/gocrud/src/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	server := fiber.New()

	dialector := postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"tuffi.db.elephantsql.com",
		"bphzpcdy",
		"NbvS91vJW_soAh9jCbE0iz5JKVSwHyty",
		"bphzpcdy",
		"5432",
	))

	db, err := gorm.Open(dialector)

	if err != nil {
		log.Fatalf("Banco não conectou. Err=%v", err.Error())
	}

	db.AutoMigrate(&structs.Installment{})

	repositoryContainer := repositories.GetRepositories(db)
	serviceContainer := services.GetServices(repositoryContainer)
	handlers.NewHandlerContainer(server, serviceContainer)

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("Server on :)")
	})

	if err := server.Listen(":3001"); err != nil {
		log.Fatalf("Aplicação não subiu. Err=%v", err.Error())
	}
}
