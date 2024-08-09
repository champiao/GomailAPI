package mail

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SendMail(data []byte) {
	c := gin.Context{}
	err := godotenv.Load()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var dataJ map[string]any
	json.Unmarshal(data, &dataJ)
	fmt.Println(dataJ["flag"])
	var (
		smtpHost string
		smtpPort string
		smtpUser string
		smtpPass string
		smtpFrom string
		smtpTo   string
	)
	if dataJ["flag"] == "FIRST" {
		smtpHost = getEnv("SMTP_HOST_FIRST", "")
		smtpPort = getEnv("SMTP_PORT_FIRST", "")
		smtpUser = getEnv("SMTP_USER_FIRST", "")
		smtpPass = getEnv("SMTP_PASS_FIRST", "")
		smtpFrom = getEnv("SMTP_USER_FIRST", "")
		smtpTo = dataJ["mailTo"].(string)

	}

	if dataJ["flag"] == "SECOND" {
		smtpHost = getEnv("SMTP_HOST_SECOND", "")
		smtpPort = getEnv("SMTP_PORT_SECOND", "")
		smtpUser = getEnv("SMTP_USER_SECOND", "")
		smtpPass = getEnv("SMTP_PASS_SECOND", "")
		smtpFrom = getEnv("SMTP_USER_SECOND", "")
		smtpTo = dataJ["mailTo"].(string)
	}
	if dataJ["flag"] == "THIRD" {
		smtpHost = getEnv("SMTP_HOST_THIRD", "")
		smtpPort = getEnv("SMTP_PORT_THIRD", "")
		smtpUser = getEnv("SMTP_USER_THIRD", "")
		smtpPass = getEnv("SMTP_PASS_THIRD", "")
		smtpFrom = getEnv("SMTP_USER_THIRD", "")
		smtpTo = dataJ["mailTo"].(string)
	}

	subjectData := dataJ["subject"].(string)
	msg := dataJ["message"].(string)
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = sendEmail(smtpFrom, smtpTo, msg, subjectData, smtpUser, smtpPass, smtpHost, port)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func sendEmail(smtpFrom string, smtpTo string, msg string, subject string, smtpUser string, smtpPass string, smtpHost string, smtpPort int) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpFrom)
	m.SetHeader("To", smtpTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", msg)
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	fmt.Printf("mensagem enviada de %v, para %v", smtpFrom, smtpTo)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
