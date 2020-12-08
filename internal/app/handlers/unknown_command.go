package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handlers)UnknownCommand(msg *tgbotapi.Message) {
	responseMsg := tgbotapi.NewMessage(msg.Chat.ID, "Такая команда недоступна")


	h.tgBot.Send(responseMsg)
}
