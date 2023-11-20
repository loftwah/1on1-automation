package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"1on1-automation/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	slackService := services.NewSlackService()
	if slackService != nil {
		go slackService.StartListening() // Start listening to Slack events if service is initialized
	}

	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	// Routes
	e.GET("/", handleRoot)
	e.GET("/health", handleHealthCheck)
	e.GET("/generate-questions", handleGenerateQuestions)
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
	prompt := "Create a set of engaging and open-ended questions in Markdown format for a casual and friendly one-on-one work meeting. Include questions about current projects, personal development, challenges, work satisfaction, and any positive aspects of work life."

	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, response)
}

func handleGenerateReport(c echo.Context) error {
	meetingResponses := c.FormValue("responses")

	prompt := "Based on these one-on-one meeting responses: " + meetingResponses + ", generate suggestions for discussion topics for the next meeting in Markdown format. Focus on areas such as progress, challenges, and opportunities for growth."

	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, response)
}
