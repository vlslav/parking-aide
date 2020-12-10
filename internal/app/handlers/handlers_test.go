package handlers

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/stretchr/testify/suite"
	. "github.com/vlslav/parking-aide/internal/app/handlers/mocks"
)

type (
	TestHandlersSuite struct {
		suite.Suite

		storageMock *StorageMock
		tgBotMock   *TgBotMock

		h *Handlers
	}

	tgMsgOption func(m *tgbotapi.Message)
)

func (s *TestHandlersSuite) SetupTest() {
	s.h = s.NewHandlers()
}

func (s *TestHandlersSuite) NewHandlers() *Handlers {
	s.storageMock = NewStorageMock(s.T())
	s.tgBotMock = NewTgBotMock(s.T())

	return &Handlers{
		storage: s.storageMock,
		tgBot:   s.tgBotMock,
	}
}

func TestHandlersSuiteTest(t *testing.T) {
	suite.Run(t, new(TestHandlersSuite))
}
