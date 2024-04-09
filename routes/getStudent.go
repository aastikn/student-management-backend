package routes

import (
	"student-management/database"
	"student-management/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)



// func GetallMarks(c *fiber.Ctx, id uint) error {
// 	StudentId := id;
// 	marks := []models.Mark{}
// 	err := database.Database.Db.Where("studentId = ?", StudentId).Error
// 	if err != nil {
// 		return c.Status(500).JSON("Error fetching users")
// 	}
// 	return c.Status(200).JSON(marks)
// }

func GetStudent(c *fiber.Ctx) error {
	StudentId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON("Error parsing JSON: " + err.Error())
	}
	student := models.Student{}
	err = database.Database.Db.Where("id = ?", StudentId).First(&student).Error
	if err != nil {
		return c.Status(404).JSON("product does not exist")
	}
	marks := []models.Mark{}
	err = database.Database.Db.Find(&marks).Error
	if err != nil {
		return c.Status(500).JSON("Error fetching marks"+err.Error())
	}
	studentDetails := fiber.Map{
		"id":           student.ID,
        "regno":        student.Regno,
        "password":     student.Password,
        "student_name": student.StudentName,
        "marks":        marks,
	}

	return c.Status(200).JSON(studentDetails)
}