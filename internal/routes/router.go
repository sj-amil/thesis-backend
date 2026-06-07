package routes

import (
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
)

//routes

func SetupRoutes(app *fiber.App) {
	
	app.Get("/", services.BlogList)
	app.Post("/", services.BlogCreate)
	app.Put("/:id",  services.BlogUpdate)
	app.Delete("/:id", services.BlogDelete)

	app.Get("/all-rfqs", services.RfqsList)
	app.Post("/create-rfqs", services.RfqsCreate)
	app.Patch("/edit-rfqs/:id", services.RfqsUpdate)
	app.Delete("/delete-rfqs/:id", services.RfqsDelete)

	app.Get("/all-quotes", services.QuoteList)
	app.Post("/create-quotes", services.QuoteCreate)
	app.Patch("/edit-quotes/:id", services.QuoteUpdate)
	app.Delete("/delete-quotes/:id", services.QuoteDelete)

	app.Get("/all-po", services.PoList)
	app.Post("/create-po", services.PoCreate)
	app.Patch("/edit-po/:id", services.PoUpdate)
	app.Delete("/delete-po/:id", services.PoDelete)

	app.Get("/all-supplier", services.SupplierList)
	app.Post("/create-supplier", services.SupplierCreate)
	app.Patch("/edit-supplier/:id", services.SupplierUpdate)
	app.Delete("/delete-supplier/:id", services.SupplierDelete)

	app.Get("/all-buyer", services.BuyerList)
	app.Post("/create-buyer", services.BuyerCreate)
	app.Patch("/edit-buyer/:id", services.BuyerUpdate)
	app.Delete("/delete-buyer/:id", services.BuyerDelete)


}