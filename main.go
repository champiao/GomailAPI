package main

import (
	"example/mail"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func handler(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "message: could not read request body")
		return
	}
	SendMail(data, c)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	router.POST("/mail", handler)
	router.Run("localhost:" + getEnv("APP_PORT", ""))
}

func SendMail(data []byte, c *gin.Context) {
	mail.SendMail(data, c)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
