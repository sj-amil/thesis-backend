package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"log"

	"github.com/gofiber/fiber/v3"
)

func UserList(c fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "ok",
		"msg": "User List",
	}

	db := database.DB

	var records []models.User
	db.Find(&records)

	

	context["user_records"] = records

	c.Status(200)
	return c.JSON(context)
}


func UserCreate(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "User Created",
	}

	record := new(models.User)

	if err := c.Bind().Body(&record); err != nil {
		log.Println("Error in Parsing.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	result := database.DB.Create(record)

	if result.Error != nil {
		log.Println("Error in Saving data.")
	}

	context["msg"] = "User is created"
	context["data"] = record

	
	c.Status(201)
	return c.JSON(context)

}





func UserDelete(c fiber.Ctx)  error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "User Deleted",
	}

	id := c.Params("id")

	var record models.User
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