package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"log"

	"github.com/gofiber/fiber/v3"
)

//main functions

func BlogList(c fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg": "Blog List",
	}

	db := database.DB

	var records []models.Blog
	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg": "Add a Blog",
	} 

	record := new(models.Blog)

	if err := c.Bind().Body(&record); err != nil {
		log.Println("Error in Parsing.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong. "
	}

	result := database.DB.Create(record)

	if result.Error != nil {
		log.Println("Error in Saving Data.")
	}

	context["msg"] = "Record is Saved. "
	context["data"] = record

	c.Status(201)
	return c.JSON(context)

}

func BlogUpdate(c fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg": "Blog Updated",
	}

	id := c.Params("id")

	var record models.Blog

	database.DB.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["msg"] = "Record not found"
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.Bind().Body(&record); err != nil {
		log.Println("Error in parsing.")
	}

	result := database.DB.Save(record)

	if  result.Error != nil {
		log.Println("Error in saving data")
	}

	context["msg"] = "Record updated successfully"
	context["data"] = record


	c.Status(200)
	return c.JSON(context)

}

func BlogDelete(c fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg": "Blog Deleted of the ID ",
	}

	id :=  c.Params("id")

	var record models.Blog

	database.DB.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["msg"] = "Record not found"
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

	context["msg"] = "Record Deleted"
	context["statusText"] = "ok"

	c.Status(200)
	return c.JSON(context)

}