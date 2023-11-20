package services

import (
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

// SlackService struct holds necessary data to communicate with the Slack API
type SlackService struct {
	api          *slack.Client
	socketClient *socketmode.Client
}

// NewSlackService creates a new instance of SlackService
func NewSlackService() *SlackService {
	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		log.Println("No Slack bot token provided. Slack integration is disabled.")
		return nil
	}

	api := slack.New(token)
	socketClient := socketmode.New(api, socketmode.OptionDebug(true), socketmode.OptionLog(nil))

	return &SlackService{
		api:          api,
		socketClient: socketClient,
	}
}

// StartListening listens for events from Slack and handles them
func (s *SlackService) StartListening() {
	if s == nil {
		return // Skip if SlackService is not initialized
	}

	go func() {
		for evt := range s.socketClient.Events {
			switch evt.Type {
			case socketmode.EventTypeInteractive:
				// Handle interactive events here
			case socketmode.EventTypeEventsAPI:
				// Handle events API callbacks here
				s.socketClient.Ack(*evt.Request)
				// Add processing logic here
			}
		}
	}()

	s.socketClient.Run()
}

// SendMessage sends a message to a specified channel
func (s *SlackService) SendMessage(channel, message string) error {
	if s == nil {
		return nil // Skip if SlackService is not initialized
	}
	_, _, err := s.api.PostMessage(channel, slack.MsgOptionText(message, false))
	return err
}

// More methods for handling different Slack interactions can be added here
