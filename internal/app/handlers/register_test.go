package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	db_storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
)

func (s *TestHandlersSuite) TestRegister() {
	ctx := context.Background()
	const userID int64 = 42

	s.Run("Пользователь успешно зарегистрировался", func() {
		h := s.NewHandlers()
		s.storageMock.SaveUserMock.Set(func(ctx context.Context, user *db_storage.User) (err error) {
			s.EqualValues(userID, user.UserID)
			return nil
		})

		responseMsg := tgbotapi.NewMessage(userID, "Привет!\n\nЧтобы сохранить местоположение, отправь геопозицию мне в чат\nСкрепка->Геопозиция")
		s.tgBotMock.SendMock.Set(func(c tgbotapi.Chattable) (m1 tgbotapi.Message, err error) {
			s.EqualValues(responseMsg, c)

			return m1, nil
		})

		h.Register(ctx, s.getTgMsg(userID))
	})

	s.Run("Ошибка БД", func() {
		h := s.NewHandlers()
		s.storageMock.SaveUserMock.Set(func(ctx context.Context, user *db_storage.User) (err error) {
			s.EqualValues(userID, user.UserID)
			return errors.New("qwerty")
		})

		responseMsg := tgbotapi.NewMessage(userID, ErrorMessage)
		s.tgBotMock.SendMock.Set(func(c tgbotapi.Chattable) (m1 tgbotapi.Message, err error) {
			s.EqualValues(responseMsg, c)

			return m1, nil
		})

		h.Register(ctx, s.getTgMsg(userID))
	})
}
