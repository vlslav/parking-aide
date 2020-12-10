package handlers

import (
	"context"

	"github.com/pkg/errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	db_storage "github.com/vlslav/parking-aide/internal/pkg/db-storage"
)

func (s *TestHandlersSuite) TestReceiveLocation() {
	ctx := context.Background()
	const userID int64 = 42
	const longitude, latitude float64 = 0.123, 0.345

	s.Run("Пользователь успешно сохранил геолокацию", func() {
		h := s.NewHandlers()
		s.storageMock.SaveUserLocationMock.Set(func(ctx context.Context, userLocation *db_storage.UserLocation) (err error) {
			s.EqualValues(userID, userLocation.UserID)
			s.EqualValues(longitude, userLocation.Longitude)
			s.EqualValues(latitude, userLocation.Latitude)

			return nil
		})

		responseMsg := tgbotapi.NewMessage(userID, "Спасибо, ваше местоположение сохранено!")
		s.tgBotMock.SendMock.Set(func(c tgbotapi.Chattable) (m1 tgbotapi.Message, err error) {
			s.EqualValues(responseMsg, c)

			return m1, nil
		})

		h.ReceiveLocation(ctx, s.getTgMsg(userID, withLatitude(latitude), withLongitude(longitude)))
	})

	s.Run("Ошибка БД", func() {
		h := s.NewHandlers()
		s.storageMock.SaveUserLocationMock.Set(func(ctx context.Context, userLocation *db_storage.UserLocation) (err error) {
			s.EqualValues(userID, userLocation.UserID)
			s.EqualValues(longitude, userLocation.Longitude)
			s.EqualValues(latitude, userLocation.Latitude)

			return errors.New("qwerty")
		})

		responseMsg := tgbotapi.NewMessage(userID, "Ой! Что-то пошло не так, попробуйте позднее")
		s.tgBotMock.SendMock.Set(func(c tgbotapi.Chattable) (m1 tgbotapi.Message, err error) {
			s.EqualValues(responseMsg, c)

			return m1, nil
		})

		h.ReceiveLocation(ctx, s.getTgMsg(userID, withLongitude(longitude), withLatitude(latitude)))
	})
}

func (s *TestHandlersSuite) getTgMsg(userID int64, options ...func(m *tgbotapi.Message)) *tgbotapi.Message {
	msg := &tgbotapi.Message{
		Chat: &tgbotapi.Chat{ID: userID},
	}

	for _, option := range options {
		option(msg)
	}

	return msg
}

func withLongitude(longitude float64) tgMsgOption {
	return func(m *tgbotapi.Message) {
		if m.Location == nil {
			m.Location = &tgbotapi.Location{}
		}
		m.Location.Longitude = longitude
	}
}

func withLatitude(latitude float64) tgMsgOption {
	return func(m *tgbotapi.Message) {
		if m.Location == nil {
			m.Location = &tgbotapi.Location{}
		}
		m.Location.Latitude = latitude
	}
}
