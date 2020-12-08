package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	db_storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
)

type (
	Handlers struct {
		storage storage
		tgBot tgBot
	}


	storage interface {
		GetUser(ctx context.Context, userID int64) (*db_storage.User, error)
		GetUserLocation(ctx context.Context, userID int64) (*db_storage.UserLocation, error)
		SaveUserLocation(ctx context.Context, userLocation *db_storage.UserLocation) error
	}

	tgBot interface {
		Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	}
)
