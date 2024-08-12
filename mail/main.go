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

func SendMail(data []byte, c *gin.Context) {

	err := godotenv.Load()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error env não inicializada abortando")
		return
	}
	var dataJ map[string]any
	json.Unmarshal(data, &dataJ)
	var (
		smtpHost string
		smtpPort string
		smtpUser string
		smtpPass string
		smtpFrom string
		smtpTo   string
	)
	if dataJ["flag"] == "first" {
		smtpHost = getEnv("SMTP_HOST_FIRST", "")
		smtpPort = getEnv("SMTP_PORT_FIRST", "")
		smtpUser = getEnv("SMTP_USER_FIRST", "")
		smtpPass = getEnv("SMTP_PASS_FIRST", "")
		smtpFrom = getEnv("SMTP_USER_FIRST", "")
		smtpTo = dataJ["mailTo"].(string)
	} else if dataJ["flag"] == "second" {
		smtpHost = getEnv("SMTP_HOST_SECOND", "")
		smtpPort = getEnv("SMTP_PORT_SECOND", "")
		smtpUser = getEnv("SMTP_USER_SECOND", "")
		smtpPass = getEnv("SMTP_PASS_SECOND", "")
		smtpFrom = getEnv("SMTP_USER_SECOND", "")
		smtpTo = dataJ["mailTo"].(string)
	} else if dataJ["flag"] == "third" {
		smtpHost = getEnv("SMTP_HOST_THIRD", "")
		smtpPort = getEnv("SMTP_PORT_THIRD", "")
		smtpUser = getEnv("SMTP_USER_THIRD", "")
		smtpPass = getEnv("SMTP_PASS_THIRD", "")
		smtpFrom = getEnv("SMTP_USER_THIRD", "")
		smtpTo = dataJ["mailTo"].(string)
	} else {
		c.Error(fmt.Errorf("invalid flag")).SetType(gin.ErrorTypePrivate)
		c.AbortWithStatusJSON(http.StatusBadRequest, "flag invalida")
		return
	}

	subjectData := dataJ["subject"].(string)
	msg := dataJ["message"].(string)
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		c.Error(fmt.Errorf("invalid port")).SetType(gin.ErrorTypePrivate)
		c.AbortWithStatusJSON(http.StatusBadRequest, "port invalida")
		return
	}
	err = sendEmail(smtpFrom, smtpTo, msg, subjectData, smtpUser, smtpPass, smtpHost, port, c)
	if err != nil {
		c.Error(fmt.Errorf("error ao enviar email")).SetType(gin.ErrorTypePrivate)
		c.AbortWithStatusJSON(http.StatusBadRequest, "erro ao enviar email")
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

func sendEmail(smtpFrom string, smtpTo string, msg string, subject string, smtpUser string, smtpPass string, smtpHost string, smtpPort int, c *gin.Context) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpFrom)
	m.SetHeader("To", smtpTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", msg)
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	fmt.Printf("mensagem enviada de %v, para '%v'", smtpFrom, smtpTo)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "email enviado de " + smtpFrom + " para " + smtpTo})
	return nil
}
