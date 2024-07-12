package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
)

func AiScanner(c *fiber.Ctx) error {
	log.Println("AiScanner")
	ctx := context.Background()
	// Access your API key as an environment variable
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyANaDL7jL7ZhbbAzz05JEBAibYvyzVuK7c"))
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "message": "Failed to create AI client", "data": err.Error()})
	}
	defer client.Close()

	// Use client.UploadFile to upload a file to the service.
	// Pass it an io.Reader.
	f, err := os.Open("img.png")
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "message": "Failed to create AI client", "data": err.Error()})
	}
	defer f.Close()

	//Optionally set a display name.
	opts := genai.UploadFileOptions{DisplayName: "burger"}
	// Let the API generate a unique `name` for the file by passing an empty string.
	// If you specify a `name`, then it has to be globally unique.
	img1, err := client.UploadFile(ctx, "", f, &opts)
	if err != nil {
		log.Fatal(err)
	}

	// View the response.
	model := client.GenerativeModel("gemini-1.5-pro")

	// Create a prompt using text and the URI reference for the uploaded file.
	prompt := []genai.Part{
		genai.FileData{URI: img1.URI},
		genai.Text("Describe this food nutrient facts."),
	}

	// Generate content using the prompt.
	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Fatal(err)
	}

	// Handle the response of generated text
	for _, c := range resp.Candidates {
		if c.Content != nil {
			fmt.Println(*c.Content)
		}
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "AI Scanner", "data": resp.Candidates})
}
