package services

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func InitAIService(ctx context.Context, apiKey string) (*genai.GenerativeModel, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %v", err)
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	return model, nil
}
func RecognizeImage(imagePath string) (string, error) {
	// Placeholder: Replace this with actual API call to Gemini image recognition
	// Here we simulate the response from the API
	simulatedResult := "Recognized image: Nasi Goreng"
	return simulatedResult, nil
}

//func RecognizeImage(imagePath string) (string, error) {
//	// Open the image file
//	file, err := os.Open(imagePath)
//	if err != nil {
//		return "", fmt.Errorf("failed to open image file: %v", err)
//	}
//	defer file.Close()
//
//	// Read the image file into a byte array
//	imageData, err := ioutil.ReadAll(file)
//	if err != nil {
//		return "", fmt.Errorf("failed to read image file: %v", err)
//	}
//
//	// Create a request to send the image to the Gemini API
//	req, err := http.NewRequest("POST", "https://api.gemini.com/v1/image-recognition", bytes.NewReader(imageData))
//	if err != nil {
//		return "", fmt.Errorf("failed to create request: %v", err)
//	}
//
//	// Set the appropriate headers
//	req.Header.Set("Content-Type", "application/octet-stream")
//	req.Header.Set("Authorization", "Bearer YOUR_API_KEY_HERE")
//
//	// Send the request
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return "", fmt.Errorf("failed to send request: %v", err)
//	}
//	defer resp.Body.Close()
//
//	// Read the response
//	respBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", fmt.Errorf("failed to read response body: %v", err)
//	}
//
//	// Process the response (assuming it's a simple text response)
//	return string(respBody), nil
//}
