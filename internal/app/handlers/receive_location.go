package handlers

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	db_storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
)

func (h *Handlers) ReceiveLocation(ctx context.Context, msg *tgbotapi.Message) {
	responseMsg := tgbotapi.NewMessage(msg.Chat.ID, "Спасибо, ваше местоположение сохранено!")

	if err := h.storage.SaveUserLocation(ctx, &db_storage.UserLocation{
		UserID:    msg.Chat.ID,
		Latitude:  msg.Location.Latitude,
		Longitude: msg.Location.Longitude,
	}); err != nil {
		log.Printf("can't save user location: %v", err)

		responseMsg.Text = ErrorMessage
	}

	h.tgBot.Send(responseMsg)
}
