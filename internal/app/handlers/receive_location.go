package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	db_storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
)

func (h *Handlers)ReceiveLocation(msg *tgbotapi.Message) {
	responseMsg := tgbotapi.NewMessage(msg.Chat.ID, "Спасибо, ваше местоположение сохранено!")

	if err := h.storage.SaveUserLocation(context.Background(), &db_storage.UserLocation{
		UserID:    msg.Chat.ID,
		Latitude:  msg.Location.Latitude,
		Longitude: msg.Location.Longitude,
	}); err != nil {
		responseMsg.Text = "Ой! Что-то пошло не так, попробуйте позднее"
	}

	h.tgBot.Send(responseMsg)
}
