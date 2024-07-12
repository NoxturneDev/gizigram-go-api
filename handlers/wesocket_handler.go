package handlers

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/berkatps/model"
	"github.com/berkatps/services"
	"github.com/gofiber/websocket/v2"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
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

	// Inisialisasi layanan AI
	aiModel, err := services.InitAIService(ctx, geminiApiKey)
	if err != nil {
		log.Println("ERROR", err)
		sendStringMessage(currentConn, "Gagal menginisialisasi layanan AI")
		return
	}

	cs := aiModel.StartChat()

	// Simulasikan koneksi ke Gemini
	fmt.Println("Menghubungkan ke Gemini...")
	sendStringMessage(currentConn, "Menghubungkan ke Gemini...")

	// Koneksi berhasil
	sendStringMessage(currentConn, "Berhasil terhubung ke Gemini")

	for {
		var payload model.SocketPayload
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			log.Println("ERROR reading JSON:", err.Error())
			sendStringMessage(currentConn, "Format JSON tidak valid")
			continue
		}

		log.Printf("Received JSON: %+v\n", payload)

		if payload.Image != "" {
			// Dekode data gambar base64
			imageData, err := base64.StdEncoding.DecodeString(payload.Image)
			if err != nil {
				log.Println("ERROR decoding image:", err)
				sendStringMessage(currentConn, "Data gambar tidak valid")
				continue
			}

			// Simpan gambar sementara
			imagePath := filepath.Join(os.TempDir(), "uploaded_image.png")
			err = ioutil.WriteFile(imagePath, imageData, 0644)
			if err != nil {
				log.Println("ERROR saving image:", err)
				sendStringMessage(currentConn, "Gagal menyimpan gambar")
				continue
			}

			// Mengenali gambar
			imageResult, err := services.RecognizeImage(imagePath)
			if err != nil {
				log.Println("ERROR recognizing image:", err)
				sendStringMessage(currentConn, "Gagal mengenali gambar")
				continue
			}

			// Kirim hasil kembali ke klien
			sendStringMessage(currentConn, fmt.Sprintf("Hasil pengenalan gambar: %s", imageResult))
		} else {
			iter := cs.SendMessageStream(ctx, genai.Text(payload.Prompt))
			for {
				resp, err := iter.Next()
				if err != nil {
					if err == iterator.Done {
						log.Println("Semua item dalam iterator telah diproses.")
						sendStringMessage(currentConn, "Semua respons telah diproses. Menunggu input baru.")
						break
					}
					log.Println("stream error:", err)
					sendStringMessage(currentConn, "Gagal menghasilkan konten")
					break
				}

				sendAiResult(currentConn, resp.Candidates[0].Content.Parts[0])
			}
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

func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		log.Println("ERROR: Gagal mendapatkan file form:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Gagal mengunggah gambar")
	}

	// Log detail file
	log.Printf("File: %+v\n", file)
	err = os.MkdirAll("/uploads", os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory", err)
	}

	// Simpan file ke disk
	filePath := fmt.Sprintf("./uploads/%d_%s", time.Now().Nanosecond(), file.Filename)
	log.Println("Menyimpan file ke:", filePath)

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal Membuat Direktori",
		})

	}

	if err := c.SaveFile(file, filePath); err != nil {
		log.Println("ERROR: Gagal menyimpan file:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menyimpan gambar"})
	}
	imageData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("ERROR: gagal membaca file: ", err)
	}

	base64Image := base64.StdEncoding.EncodeToString(imageData)

	if len(connections) > 0 {
		currentCon := connections[0]
		err := currentCon.WriteJSON(model.SocketPayload{
			Prompt: "Rancang aplikasi yang ramah pengguna untuk menganalisis data anak-anak dalam jangka waktu tertentu guna mengidentifikasi potensi stunting berdasarkan standar WHO, dengan fitur seperti entri data, analisis data, pelaporan, tindakan pencegahan, manajemen pengguna, aksesibilitas mobile, dan dukungan serta umpan balik, serta integrasikan data WHO tentang nutrisi, perawatan kesehatan, kebersihan, sanitasi, pendidikan, dan kesadaran, untuk memberikan rekomendasi pencegahan stunting dan isyarat visual ketika metrik pertumbuhan anak di bawah ambang batas WHO\n",
			Image:  base64Image,
		})
		if err != nil {
			log.Println("ERROR: gagal mengirim gambar melalui websocket: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "gagal mengirim gambar melalui websocket",
			})
		}
		log.Println("Gambar berhasil dikirim ke WebSocket")
	} else {
		log.Println("Tidak ada koneksi WebSocket aktif")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mengirim gambar melalui websocket",
		})
	}
	log.Println("File berhasil disimpan:", filePath)

	return c.JSON(fiber.Map{
		"message":  "Gambar berhasil diunggah",
		"filePath": filePath,
	})
}
