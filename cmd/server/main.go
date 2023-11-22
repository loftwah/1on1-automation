package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"1on1-automation/internal/services"
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

	// Static files
	e.Static("/", "../../frontend/src")
	e.Static("/css", "../../frontend/src/css")
	e.Static("/js", "../../frontend/src/js")
	e.Static("/images", "../../frontend/src/images")

	// Routes
	e.GET("/", handleRoot)
	e.GET("/dynamic-content", handleDynamicContent)
	e.GET("/health", handleHealthCheck)
	e.GET("/htmx", handleHtmx)
	e.GET("/generate-questions", handleGenerateQuestions)
	e.POST("/generate-suggestions", handleGenerateSuggestions)

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
	return c.File("frontend/src/index.html")
}

func handleHtmx(c echo.Context) error {
	return c.File("frontend/src/htmx.html")
}

func handleDynamicContent(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "image/png")
	return c.File("frontend/src/images/htmx.png")
}

func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK", "message": "Service is operational"})
}

func handleGenerateQuestions(c echo.Context) error {
	prompt := "Craft a series of engaging, thought-provoking, and open-ended questions in Markdown format for a relaxed yet insightful one-on-one work meeting. The questions should delve into areas like current project status, personal and professional growth, workplace challenges, overall job satisfaction, and the bright spots in their work life. Aim to foster a sense of comfort and candid conversation, encouraging both reflection and forward-thinking."

	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, response)
}

func handleGenerateSuggestions(c echo.Context) error {
	meetingResponses := c.FormValue("responses")

	prompt := "Analyze these one-on-one meeting responses: " + meetingResponses + ". Synthesize thoughtful, constructive, and actionable suggestions for discussion topics in the next meeting, presented in Markdown format. Emphasize areas such as project milestones, emerging challenges, professional development opportunities, team dynamics, and paths for future growth. Ensure the suggestions are framed to promote positive dialogue, engagement, and mutual understanding."

	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, response)
}
