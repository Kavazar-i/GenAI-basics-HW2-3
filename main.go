package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	openai "github.com/sashabaranov/go-openai"
	"GenAI-basics-HW2-3/utils"
)

var (
	bot *tgbotapi.BotAPI
	openaiClient *openai.Client
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal("Failed to initialize bot:", err)
	}

	openaiClient = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}

func generateStory(ctx context.Context) (string, error) {
	resp, err := openaiClient.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a creative storyteller. Generate a short sci-fi story in maximum 400 characters.",
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func main() {
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Create context that will be canceled on SIGINT or SIGTERM
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Received shutdown signal, cleaning up...")
		cancel()
	}()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down bot...")
			return
		case update := <-updates:
			if update.Message == nil {
				continue
			}

			var response string
			var err error

			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "story":
					storyCtx, storyCancel := context.WithTimeout(ctx, 10*time.Second)
					response, err = generateStory(storyCtx)
					storyCancel()
					if err != nil {
						log.Printf("Error generating story: %v", err)
						response = "Sorry, I couldn't generate a story at the moment."
					}
				default:
					response = "Unknown command"
				}
			} else {
				response = utils.ReverseString(update.Message.Text)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	}
} 