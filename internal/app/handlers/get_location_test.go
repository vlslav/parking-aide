package handlers

import (
	"context"

	"github.com/pkg/errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	db_storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
)

func (s *TestHandlersSuite) TestGetLocation() {
	ctx := context.Background()
	const userID int64 = 42
	const longitude, latitude float64 = 0.123, 0.345

	s.Run("Пользователь успешно получил геолокацию", func() {
		h := s.NewHandlers()
		s.storageMock.GetUserLocationMock.Set(func(ctx context.Context, userID1 int64) (up1 *db_storage.UserLocation, err error) {
			s.EqualValues(userID, userID1)
			return &db_storage.UserLocation{
				UserID:    userID,
				Latitude:  latitude,
				Longitude: longitude,
			}, nil
		})

		responseMsg := tgbotapi.NewLocation(userID, latitude, longitude)
		s.tgBotMock.SendMock.Set(func(c tgbotapi.Chattable) (m1 tgbotapi.Message, err error) {
			s.EqualValues(responseMsg, c)

			return m1, nil
		})

		h.GetLocation(ctx, s.getTgMsg(userID))
	})

	s.Run("Ошибка БД", func() {
		h := s.NewHandlers()
		s.storageMock.GetUserLocationMock.Set(func(ctx context.Context, userID1 int64) (up1 *db_storage.UserLocation, err error) {
			s.EqualValues(userID, userID1)
			return nil, errors.New("qwerty")
		})

		responseMsg := tgbotapi.NewMessage(userID, ErrorMessage)
		s.tgBotMock.SendMock.Set(func(c tgbotapi.Chattable) (m1 tgbotapi.Message, err error) {
			s.EqualValues(responseMsg, c)

			return m1, nil
		})

		h.GetLocation(ctx, s.getTgMsg(userID))
	})
}
