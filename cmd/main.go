package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vlslav/parking-aide/internal/app/handlers"
	command_resolver "github.com/vlslav/parking-aide/internal/pkg/command-resolver"
	storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	// Получаем токен телеги, бота, канал обновлений
	tgToken := getTgToken()
	tgBot, tgBotUpdates := getTgBotAndUpdates(tgToken)

	// Инициализируем db-storage
	dataStorage := storage.New()

	// Инициализируем обработчики сообщений пользователя
	handlers := handlers.New()
	handlers.SetTgBot(tgBot)
	handlers.SetStorage(dataStorage) // TODO

	// Инициализируем резолверы сообщений в обработчики
	resolver := command_resolver.New()
	resolver.SetUpdatesChannel(tgBotUpdates)
	resolver.SetResolvers(map[string]func(msg *tgbotapi.Message){
		"/send_location" : handlers.ReceiveLocation,
		"/unknown_command" : handlers.UnknownCommand,
	})

	// Запускаем резолвер
	resolver.Start(ctx)
}

func getTgToken() string {
	tgToken := os.Getenv("TELEGRAM_TOKEN")
	if tgToken == "" {
		log.Fatal("tg token not specified")
	}

	return tgToken
}


func getTgBotAndUpdates(token string) (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	tgBot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("can't create bot api: %v", err)
	}

	tgBot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := tgBot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("can't get updates: %v", err)
	}

	return tgBot, updates
}