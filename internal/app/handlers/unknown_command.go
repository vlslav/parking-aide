package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handlers) UnknownCommand(ctx context.Context, msg *tgbotapi.Message) {
	responseMsg := tgbotapi.NewMessage(msg.Chat.ID, "Такая команда недоступна")

	h.tgBot.Send(responseMsg)
}
