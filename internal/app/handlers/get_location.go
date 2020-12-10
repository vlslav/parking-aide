package handlers

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handlers) GetLocation(ctx context.Context, msg *tgbotapi.Message) {
	userLocation, err := h.storage.GetUserLocation(ctx, msg.Chat.ID)
	if err != nil {
		log.Printf("can't get user location: %v", err)

		responseMsg := tgbotapi.NewMessage(msg.Chat.ID, ErrorMessage)
		h.tgBot.Send(responseMsg)
		return
	}

	if _, err := h.tgBot.Send(tgbotapi.NewLocation(
		userLocation.UserID,
		userLocation.Latitude,
		userLocation.Longitude,
	)); err == nil {
		h.storage.DeleteUserLocation(ctx, userLocation)
	}
}
