package main

import (
	"net/http"
	"os"
	"strings"

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
	// Define categories of questions
	categories := []string{"About Manager", "Career development", "Conversation starters", "Job satisfaction", "Team and company", "Work-life"}

	// Generate a friendly and personalized prompt for OpenAI
	prompt := "Create a personalized and friendly set of questions for a one-on-one meeting with a manager. Include questions from the following categories: " + strings.Join(categories, ", ") + ". Format the questions in Markdown for easy readability. Advise the user that they don't have to answer anything they're not comfortable with and this is just to help guide the conversation."

	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Format response in Markdown
	markdownResponse := "### One-on-One Meeting Questions\n\n" + response

	return c.String(http.StatusOK, markdownResponse)
}

func handleGenerateReport(c echo.Context) error {
	// Assume the responses from the one-on-one meeting are passed as a parameter
	meetingResponses := c.FormValue("responses")

	// Example prompt to generate a report
	prompt := "Generate a report summarizing the key points and actionable items from the following one-on-one meeting responses: " + meetingResponses

	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"report": response})
}
