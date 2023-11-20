package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"1on1-automation/services"
	"1on1-automation/slack_service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	slackService := slack_service.NewSlackService()
	go slackService.StartListening() // Start listening to Slack events

	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	// Routes
	e.GET("/", handleRoot)
	e.GET("/health", handleHealthCheck)
	e.POST("/slack/events", slackService.HandleSlackEvents)
	e.POST("/generate-questions", handleGenerateQuestions)
	e.POST("/generate-report", handleGenerateReport)

	e.Logger.Fatal(e.Start(":" + getServerPort()))
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	return port
}

func handleRoot(c echo.Context) error {
	return c.String(http.StatusOK, "AI-Enhanced One-on-One Meeting Assistant is running")
}

func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK", "message": "Service is operational"})
}

func handleGenerateQuestions(c echo.Context) error {
	// Example prompt to generate questions
	prompt := "Generate a list of discussion questions for a one-on-one meeting focusing on career development."

	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"questions": response})
}

func handleGenerateReport(c echo.Context) error {
	// Example prompt to generate a report
	prompt := "Generate a report summarizing the key points from a one-on-one meeting about project management."

	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"report": response})
}
