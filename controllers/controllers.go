package controllers

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/leocm889/server/models"
)

func sendEmail(name, email, phone, message string) error {

	usernameEmail := os.Getenv("EMAIL")
	passwordEmail := os.Getenv("PASSWORD")

	var from string = usernameEmail
	var password string = passwordEmail
	var to string = usernameEmail

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	var subject string = "Contact Form Submission"

	var body string = fmt.Sprintf(
		"Name: %s\r\n"+
			"Email: %s\r\n"+
			"Phone: %s\r\n"+
			"Message: %s\r\n",
		name, email, phone, message)

	var msg string = fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n%s",
		from, to, subject, body)
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}

func SubmitFormHandler(c *fiber.Ctx) error {
	data := new(models.FormData)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	name := data.FirstName + " " + data.LastName
	email := data.Email
	phone := data.Phone
	message := data.Message

	if err := sendEmail(name, email, phone, message); err != nil {
		return err
	}

	// return c.SendString("Thanks for contacting us!")
	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{"message": "Thanks for contacting us!", "data": data})
}
