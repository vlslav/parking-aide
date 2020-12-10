package handlers

func New() *Handlers {
	return &Handlers{}
}

func (h *Handlers) SetStorage(storage storage) {
	h.storage = storage
}

func (h *Handlers) SetTgBot(tgBot tgBot) {
	h.tgBot = tgBot
}
