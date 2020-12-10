package handlers

import (
	"context"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	db_storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
)

func (h *Handlers) Register(ctx context.Context, msg *tgbotapi.Message) {
	responseMsg := tgbotapi.NewMessage(msg.Chat.ID, "Привет!\n\nЧтобы сохранить местоположение, отправь геопозицию мне в чат\nСкрепка->Геопозиция")

	if err := h.storage.SaveUser(ctx, &db_storage.User{
		UserID:       msg.Chat.ID,
		RegisteredAt: time.Now(),
	}); err != nil {
		log.Printf("can't save user: %v", err)

		responseMsg.Text = ErrorMessage
	}

	h.tgBot.Send(responseMsg)
}
