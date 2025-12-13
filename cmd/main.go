package main

import (
	"SnipSnapBot/internal/handlers"
	"SnipSnapBot/internal/repositories"
	"SnipSnapBot/internal/services"
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	botToken := os.Getenv("BOT_TOKEN")
	mongoURI := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping mongo: %v", err)
	}
	log.Println("Connected to Mongo")

	collection := client.Database("snipsnap").Collection("links")

	repo := repositories.NewMongoLinkRepo(collection)
	service := services.NewLinkService(repo)

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("failed to create telegram bot: %v", err)
	}

	h := handlers.NewHandlers(*bot, service)

	bot.Handle("/save", h.SaveHandler)
	bot.Handle("/list", h.ListHandler)
	bot.Handle("/get", h.GetHandler)
	bot.Handle("/remove", h.RemoveHandler)
	bot.Handle("/edit", h.EditHandler)
	bot.Handle(telebot.OnText, h.HandleUnkownCommand)
	log.Println("Bot started")
	bot.Start()
}
