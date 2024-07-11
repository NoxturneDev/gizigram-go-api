package model

import (
	"context"

	"github.com/gofiber/websocket/v2"
	"github.com/google/generative-ai-go/genai"
)

type AI struct {
	GenerativeModel *genai.GenerativeModel
	Context         context.Context
}

type WebSocketConnection struct {
	*websocket.Conn
	PhoneNumber string
}

type SocketResponse struct {
	Message string
	Result  *genai.GenerateContentResponse
	Part    genai.Part
}

type SocketPayload struct {
	Prompt string
}
