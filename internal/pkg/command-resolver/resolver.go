package command_resolver

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandResolver struct {
	resolvers map[string]func(ctx context.Context, msg *tgbotapi.Message)
	tgUpdates tgbotapi.UpdatesChannel
}

func New() *CommandResolver {
	return &CommandResolver{}
}

func (c *CommandResolver) SetResolvers(resolvers map[string]func(ctx context.Context, msg *tgbotapi.Message)) {
	c.resolvers = resolvers
}

func (c *CommandResolver) SetUpdatesChannel(tgUpdates tgbotapi.UpdatesChannel) {
	c.tgUpdates = tgUpdates
}

func (c *CommandResolver) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Print("return from CommandResolver Start, ctx canceled")
			return
		case update := <-c.tgUpdates:
			if update.Message == nil {
				continue
			}

			msg := update.Message

			// грязный хак
			if msg.Location != nil {
				go c.resolvers[sendLocation](ctx, msg)

				continue
			}

			resolver, exists := c.resolvers[msg.Text]
			if !exists {
				go c.resolvers["/unknown_command"](ctx, msg)
				continue
			}

			go resolver(ctx, msg)
		}
	}
}
