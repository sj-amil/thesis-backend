package main

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {

	
	database.ConnectDB()

	fmt.Println("Running database migrations...")
	err := database.DB.AutoMigrate(new(models.Blog), new(models.RFQ), new(models.RFQItem), new(models.Quotation), new(models.SupplierProfile))
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed. Refresh pgAdmin!")


	app := fiber.New()


	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		 AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	routes.SetupRoutes(app)
	
	app.Listen(":8000")
}

