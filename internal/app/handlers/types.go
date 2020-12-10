package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	db_storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
)

const ErrorMessage = "Ой! Что-то пошло не так, попробуйте позднее"

type (
	Handlers struct {
		storage storage
		tgBot   tgBot
	}

	storage interface {
		// таблица users
		SaveUser(ctx context.Context, user *db_storage.User) error
		GetUser(ctx context.Context, userID int64) (*db_storage.User, error)
		// таблица user_location
		GetUserLocation(ctx context.Context, userID int64) (*db_storage.UserLocation, error)
		SaveUserLocation(ctx context.Context, userLocation *db_storage.UserLocation) error
		DeleteUserLocation(ctx context.Context, userLocation *db_storage.UserLocation) error
	}

	tgBot interface {
		Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	}
)
