package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
	"strings"
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

	// Generate content using the prompt.
	promptDesignBase := "Hai, saya ingin kamu bertindak seperti ahli gizi pribadi saya. Tapi kamu perlu mengirim tanggapannya dalam format JSON (tidak perlu menambahkan JSON sebagai awalan untuk jawabannya) dengan objek: saran (advice), fakta nutrisi (nutrition fact sebagai objek), data anak (nama, usia). Berikan saya beberapa saran sehat berdasarkan makanan dan nutrisi. Pertama-tama, anak saya berusia [age] dan namanya [name]. Saya ingin memberinya makanan ini pada gambar. dan ini adalah deskripsi lebih lanjut tentangnya. [prompt]. Tolong juga tulis beberapa saran apakah makanan ini direkomendasikan berdasarkan nutrisinya. Saya khawatir tentang masalah stunting, jadi saya ingin anak saya se-sehat mungkin."

	promptDesign1 := strings.Replace(promptDesignBase, "[age]", "5", -1)
	promptDesign2 := strings.Replace(promptDesign1, "[name]", "galih adhi kusuma", -1)
	promptDesign3 := strings.Replace(promptDesign2, "[prompt]", "describe the food", -1)

	prompt := []genai.Part{
		genai.FileData{URI: img1.URI},
		genai.Text(promptDesign3),
	}

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
