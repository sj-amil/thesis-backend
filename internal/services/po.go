package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func PoList(c fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "ok",
		"msg": "PO List",
	}

	db := database.DB

	var records []models.PurchaseOrder
	db.Find(&records)

	

	context["quote_records"] = records

	c.Status(200)
	return c.JSON(context)
}


func PoCreate(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "PO Created",
	}

	record := new(models.PurchaseOrder)

	if err := c.Bind().Body(&record); err != nil {
		log.Println("Error in Parsing.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	result := database.DB.Create(record)

	if result.Error != nil {
		log.Println("Error in Saving data.")
	}

	context["msg"] = "Record is created"
	context["data"] = record

	
	c.Status(201)
	return c.JSON(context)

}

func PoUpdate(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "PO Updated",
	}

	id := c.Params("id")

	var record models.PurchaseOrder
	database.DB.First(&record, id)

	if record.ID == uuid.Nil {
		log.Println("Record not found")
		context["msg"] = "Record not found"
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.Bind().Body(&record); err != nil {
		log.Println("Error in Parsing.")
	}

	result := database.DB.Save(&record)

	if result.Error != nil {
		log.Println("Error in saving data.")
	}

	context["msg"] = "Record updated successfully."
	context["data"] = record 
	
	c.Status(200)
	return c.JSON(context)
}



func PoDelete(c fiber.Ctx)  error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "PO Deleted",
	}

	id := c.Params("id")

	var record models.PurchaseOrder
	database.DB.First(&record, id)
	
	if record.ID == uuid.Nil {
		log.Println("Record not found")
		context["msg"] = "Record not found."
		c.Status(400)
		return c.JSON(context)
	}

	result := database.DB.Delete(record)

	if result.Error != nil {
		log.Println("Something went wrong")
		context["msg"] = "something went wrong"
		c.Status(400)
		return c.JSON(context)
	}

	context["msg"] = "Data deleted"
	
	c.Status(200)
	return c.JSON(context)
}