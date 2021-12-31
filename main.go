package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Sms struct {
	Body string `json:"body"`
}

func main() {
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ping Pong!")
	})

	app.Get("/sms/:phoneNumber", func(c *fiber.Ctx) error {
		phoneNumber := c.Params("phoneNumber")
		sms := new(Sms)
		if err := c.BodyParser(sms); err != nil {
			return err
		}
		accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
		authToken := os.Getenv("TWILIO_AUTH_TOKEN")

		client := twilio.NewRestClientWithParams(twilio.RestClientParams{
			Username: accountSid,
			Password: authToken,
		})

		params := &openapi.CreateMessageParams{}
		params.SetTo(phoneNumber)
		params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
		params.SetBody(sms.Body)

		resp, err := client.ApiV2010.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
			err = nil
		}
		return c.Status(200).SendString("Message Sid: " + *resp.Sid)
	})

	app.Listen(":3006")

}
