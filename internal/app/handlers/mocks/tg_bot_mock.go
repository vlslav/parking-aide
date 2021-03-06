package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/vlslav/parking-aide/internal/app/handlers.tgBot -o ./internal/app/handlers/mocks/tg_bot_mock.go

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gojuno/minimock/v3"
)

// TgBotMock implements handlers.tgBot
type TgBotMock struct {
	t minimock.Tester

	funcSend          func(c tgbotapi.Chattable) (m1 tgbotapi.Message, err error)
	inspectFuncSend   func(c tgbotapi.Chattable)
	afterSendCounter  uint64
	beforeSendCounter uint64
	SendMock          mTgBotMockSend
}

// NewTgBotMock returns a mock for handlers.tgBot
func NewTgBotMock(t minimock.Tester) *TgBotMock {
	m := &TgBotMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendMock = mTgBotMockSend{mock: m}
	m.SendMock.callArgs = []*TgBotMockSendParams{}

	return m
}

type mTgBotMockSend struct {
	mock               *TgBotMock
	defaultExpectation *TgBotMockSendExpectation
	expectations       []*TgBotMockSendExpectation

	callArgs []*TgBotMockSendParams
	mutex    sync.RWMutex
}

// TgBotMockSendExpectation specifies expectation struct of the tgBot.Send
type TgBotMockSendExpectation struct {
	mock    *TgBotMock
	params  *TgBotMockSendParams
	results *TgBotMockSendResults
	Counter uint64
}

// TgBotMockSendParams contains parameters of the tgBot.Send
type TgBotMockSendParams struct {
	c tgbotapi.Chattable
}

// TgBotMockSendResults contains results of the tgBot.Send
type TgBotMockSendResults struct {
	m1  tgbotapi.Message
	err error
}

// Expect sets up expected params for tgBot.Send
func (mmSend *mTgBotMockSend) Expect(c tgbotapi.Chattable) *mTgBotMockSend {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("TgBotMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &TgBotMockSendExpectation{}
	}

	mmSend.defaultExpectation.params = &TgBotMockSendParams{c}
	for _, e := range mmSend.expectations {
		if minimock.Equal(e.params, mmSend.defaultExpectation.params) {
			mmSend.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSend.defaultExpectation.params)
		}
	}

	return mmSend
}

// Inspect accepts an inspector function that has same arguments as the tgBot.Send
func (mmSend *mTgBotMockSend) Inspect(f func(c tgbotapi.Chattable)) *mTgBotMockSend {
	if mmSend.mock.inspectFuncSend != nil {
		mmSend.mock.t.Fatalf("Inspect function is already set for TgBotMock.Send")
	}

	mmSend.mock.inspectFuncSend = f

	return mmSend
}

// Return sets up results that will be returned by tgBot.Send
func (mmSend *mTgBotMockSend) Return(m1 tgbotapi.Message, err error) *TgBotMock {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("TgBotMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &TgBotMockSendExpectation{mock: mmSend.mock}
	}
	mmSend.defaultExpectation.results = &TgBotMockSendResults{m1, err}
	return mmSend.mock
}

//Set uses given function f to mock the tgBot.Send method
func (mmSend *mTgBotMockSend) Set(f func(c tgbotapi.Chattable) (m1 tgbotapi.Message, err error)) *TgBotMock {
	if mmSend.defaultExpectation != nil {
		mmSend.mock.t.Fatalf("Default expectation is already set for the tgBot.Send method")
	}

	if len(mmSend.expectations) > 0 {
		mmSend.mock.t.Fatalf("Some expectations are already set for the tgBot.Send method")
	}

	mmSend.mock.funcSend = f
	return mmSend.mock
}

// When sets expectation for the tgBot.Send which will trigger the result defined by the following
// Then helper
func (mmSend *mTgBotMockSend) When(c tgbotapi.Chattable) *TgBotMockSendExpectation {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("TgBotMock.Send mock is already set by Set")
	}

	expectation := &TgBotMockSendExpectation{
		mock:   mmSend.mock,
		params: &TgBotMockSendParams{c},
	}
	mmSend.expectations = append(mmSend.expectations, expectation)
	return expectation
}

// Then sets up tgBot.Send return parameters for the expectation previously defined by the When method
func (e *TgBotMockSendExpectation) Then(m1 tgbotapi.Message, err error) *TgBotMock {
	e.results = &TgBotMockSendResults{m1, err}
	return e.mock
}

// Send implements handlers.tgBot
func (mmSend *TgBotMock) Send(c tgbotapi.Chattable) (m1 tgbotapi.Message, err error) {
	mm_atomic.AddUint64(&mmSend.beforeSendCounter, 1)
	defer mm_atomic.AddUint64(&mmSend.afterSendCounter, 1)

	if mmSend.inspectFuncSend != nil {
		mmSend.inspectFuncSend(c)
	}

	mm_params := &TgBotMockSendParams{c}

	// Record call args
	mmSend.SendMock.mutex.Lock()
	mmSend.SendMock.callArgs = append(mmSend.SendMock.callArgs, mm_params)
	mmSend.SendMock.mutex.Unlock()

	for _, e := range mmSend.SendMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.m1, e.results.err
		}
	}

	if mmSend.SendMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSend.SendMock.defaultExpectation.Counter, 1)
		mm_want := mmSend.SendMock.defaultExpectation.params
		mm_got := TgBotMockSendParams{c}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSend.t.Errorf("TgBotMock.Send got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSend.SendMock.defaultExpectation.results
		if mm_results == nil {
			mmSend.t.Fatal("No results are set for the TgBotMock.Send")
		}
		return (*mm_results).m1, (*mm_results).err
	}
	if mmSend.funcSend != nil {
		return mmSend.funcSend(c)
	}
	mmSend.t.Fatalf("Unexpected call to TgBotMock.Send. %v", c)
	return
}

// SendAfterCounter returns a count of finished TgBotMock.Send invocations
func (mmSend *TgBotMock) SendAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.afterSendCounter)
}

// SendBeforeCounter returns a count of TgBotMock.Send invocations
func (mmSend *TgBotMock) SendBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.beforeSendCounter)
}

// Calls returns a list of arguments used in each call to TgBotMock.Send.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSend *mTgBotMockSend) Calls() []*TgBotMockSendParams {
	mmSend.mutex.RLock()

	argCopy := make([]*TgBotMockSendParams, len(mmSend.callArgs))
	copy(argCopy, mmSend.callArgs)

	mmSend.mutex.RUnlock()

	return argCopy
}

// MinimockSendDone returns true if the count of the Send invocations corresponds
// the number of defined expectations
func (m *TgBotMock) MinimockSendDone() bool {
	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSend != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendInspect logs each unmet expectation
func (m *TgBotMock) MinimockSendInspect() {
	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TgBotMock.Send with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		if m.SendMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TgBotMock.Send")
		} else {
			m.t.Errorf("Expected call to TgBotMock.Send with params: %#v", *m.SendMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSend != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		m.t.Error("Expected call to TgBotMock.Send")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TgBotMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSendInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TgBotMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *TgBotMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendDone()
}
