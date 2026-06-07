package main

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {

	
	database.ConnectDB()

	fmt.Println("Running database migrations...")
	err := database.DB.AutoMigrate(new(models.Blog), new(models.RFQ), new(models.RFQItem), new(models.Quotation))
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed. Refresh pgAdmin!")


	app := fiber.New()

	routes.SetupRoutes(app)
	
	app.Listen(":8000")
}

