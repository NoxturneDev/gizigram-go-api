package handlers

import (
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"os"

	"github.com/berkatps/model"
	"github.com/berkatps/services"
	"github.com/gofiber/websocket/v2"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
)

var connections = make([]*model.WebSocketConnection, 0)

func AiChatHandler(c *websocket.Conn) {
	defer c.Close()

	phoneNumber := c.Query("phoneNumber")

	currentConn := model.WebSocketConnection{Conn: c, PhoneNumber: phoneNumber}
	connections = append(connections, &currentConn)

	handleChat(&currentConn, connections)
}

func handleChat(currentConn *model.WebSocketConnection, connections []*model.WebSocketConnection) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	geminiApiKey := os.Getenv("API_KEY")

	// Initialize AI service
	aiModel, err := services.InitAIService(ctx, geminiApiKey)
	if err != nil {
		log.Println("ERROR", err)
		sendStringMessage(currentConn, "Failed to initialize AI service")
		return
	}

	cs := aiModel.StartChat()

	// Simulate connecting to Gemini
	fmt.Println("Connecting to Gemini...")
	sendStringMessage(currentConn, "Connecting to Gemini...")

	// Connection successful
	sendStringMessage(currentConn, "Successfully connected to Gemini")

	for {
		var payload model.SocketPayload
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			log.Println("ERROR reading JSON:", err.Error())
			sendStringMessage(currentConn, "Invalid JSON format")
			continue
		}

		log.Printf("Received JSON: %+v\n", payload)

		iter := cs.SendMessageStream(ctx, genai.Text(payload.Prompt))
		for {
			resp, err := iter.Next()
			if err != nil {
				if err == iterator.Done {
					log.Println("All items in iterator processed.")
					sendStringMessage(currentConn, "All responses processed. Waiting for new input.")
					break
				}
				log.Println("stream error:", err)
				sendStringMessage(currentConn, "Failed to generate content")
				break
			}

			sendAiResult(currentConn, resp.Candidates[0].Content.Parts[0])
		}
	}
}

func sendStringMessage(currentConn *model.WebSocketConnection, message string) {
	currentConn.WriteJSON(model.SocketResponse{
		Message: message,
	})
}

func sendAiResult(currentConn *model.WebSocketConnection, part genai.Part) {
	currentConn.WriteJSON(model.SocketResponse{
		Part: part,
	})
}
