package handlers

package handlers

import (
"context"
"errors"
"fmt"
"log"
"os"
"strings"

"github.com/gofiber/websocket/v2"
"github.com/joho/godotenv"
"google.golang.org/api/iterator"

)

var connections = make([]*model.WebSocketConnection, 0)

func AiChatHandler(c *websocket.Conn) {
	defer c.Close()

	phoneNumber := c.Query("name")

	currentConn := model.WebSocketConnection{Conn: c, PhoneNumber: phoneNumber}
	connections = append(connections, &currentConn)

	handleChat(&currentConn, connections)
}

func handleChat(currentConn *model.WebSocketConnection, connections []*model.WebSocketConnection) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: ", fmt.Sprintf("%v", r))
		}
	}()

	ctx := context.Background()
	geminiApiKey := os.Getenv("GEMINI_API_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Println("GEMINI_API_KEY:", geminiApiKey)
	fmt.Println("SECRET_KEY:", secretKey)

	model, err := services.InitAIService(ctx, geminiApiKey)
	if err != nil {
		log.Println("ERROR", err)
		return
	}

	cs := model.StartChat()

	for {
		payload := model.SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				sendStringMessage(currentConn, "LEAVE")
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		iter := cs.SendMessageStream(ctx, genai.Text(payload.Prompt))
		for {
			resp, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				fmt.Println(iter.MergedResponse().Candidates[0].Content.Parts[0])
				break
			}
			if err != nil {
				log.Println("stream error:", err)
				sendStringMessage(currentConn, "Failed to generate content")
			}

			log.Println("Sending message:", resp)
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
