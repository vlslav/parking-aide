package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	"github.com/vlslav/parking-aide/internal/app/handlers"
	command_resolver "github.com/vlslav/parking-aide/internal/pkg/command-resolver"
	storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Получаем токен телеги, бота, канал обновлений
	tgToken := getTgToken()
	tgBot, tgBotUpdates := getTgBotAndUpdates(tgToken)

	// Инициализируем db-storage
	dataStorage := storage.New(newDB())

	// Инициализируем обработчики сообщений пользователя
	handlers := handlers.New()
	handlers.SetTgBot(tgBot)
	handlers.SetStorage(dataStorage)

	// Инициализируем резолверы сообщений в обработчики
	resolver := command_resolver.New()
	resolver.SetUpdatesChannel(tgBotUpdates)
	resolver.SetResolvers(map[string]func(ctx context.Context, msg *tgbotapi.Message){
		"/start":           handlers.Register,
		"/get_location":    handlers.GetLocation,
		"/unknown_command": handlers.UnknownCommand,
	})

	gracefulShutdown(cancel)

	// Запускаем резолвер
	resolver.Start(ctx)
}

func getTgToken() string {
	tgToken := os.Getenv("TELEGRAM_TOKEN")
	if tgToken == "" {
		//log.Fatal("tg token not specified")
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

func newDB() *sqlx.DB {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("tg token not specified")
	}

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("can't connect db: %v", err)
	}

	return db
}

func gracefulShutdown(cancelFunc context.CancelFunc) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGILL)

	go func() {
		sig := <-sigs
		log.Printf("catched signal: %v", sig)
		cancelFunc()
	}()
}
