package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"log"

	"github.com/gofiber/fiber/v3"
)

func RfqsList(c fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "ok",
		"msg": "Rfq List",
	}

	db := database.DB

	var records []models.RFQ
	db.Find(&records)

	

	context["rqfs_records"] = records

	c.Status(200)
	return c.JSON(context)
}


func RfqsCreate(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "Rfq Created",
	}

	record := new(models.RFQ)

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

func RfqsUpdate(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "Rfq Updated",
	}

	id := c.Params("id")

	var record models.RFQ
	database.DB.First(&record, id)

	if record.ID == 0 {
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



func RfqsDelete(c fiber.Ctx)  error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "Rfq Deleted",
	}

	id := c.Params("id")

	var record models.RFQ
	database.DB.First(&record, id)
	
	if record.ID == 0 {
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