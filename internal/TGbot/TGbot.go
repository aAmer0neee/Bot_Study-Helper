package TGbot

import (
	"github.com/aAmer0neee/Bot_Study-Helper/internal/TGbot/handlers"
	cfg "github.com/aAmer0neee/Bot_Study-Helper/internal/config"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

var (
	envLocal = "local"
)

type (
	BotWithConfig struct {
		Bot    *tgbotapi.BotAPI
		Config tgbotapi.UpdateConfig

		Handler        *handlers.Handler
		ButtonHandlers map[string]func(u tgbotapi.Update) tgbotapi.Chattable
		commands       map[string]func(u tgbotapi.Update) tgbotapi.Chattable
	}
)

func ConfugureBot(cfg cfg.Config, logger *slog.Logger, handler *handlers.Handler) *BotWithConfig {
	b := &BotWithConfig{
		Handler:        handler,
		commands:       make(map[string]func(tgbotapi.Update) tgbotapi.Chattable),
		ButtonHandlers: make(map[string]func(tgbotapi.Update) tgbotapi.Chattable),
	}
	var err error
	if b.Bot, err = tgbotapi.NewBotAPI(cfg.BotToken); err != nil {
		logger.Error("Create bot")
	}
	logger.Info("Authorized on account", "bot_username", b.Bot.Self.UserName)

	if cfg.Env == envLocal {
		b.Bot.Debug = true
		b.Config = tgbotapi.NewUpdate(0)
		b.Config.Timeout = 30
	}

	/* b.configureBothandlers() */

	return b
}

func (b *BotWithConfig) ListenAndHandleUpdates() {

	updates := b.Bot.GetUpdatesChan(b.Config)
	var response tgbotapi.Chattable
	for update := range updates {
		if update.Message != nil {
			action := update.Message.Text
			switch action {
				case "/start":
					response = b.Handler.HandleStart(update)
				
			}

			if update.CallbackQuery != nil {
				action := update.CallbackQuery.Data
				if handler, exists := b.ButtonHandlers[action]; exists {
					response = handler(update)
				}
			}
		}
		if response != nil {
			b.sendResponse(response)
		}
	}
}

/* func (b *BotWithConfig) configureBothandlers() {

	b.ButtonHandlers["➕ Добавить задачу"] = handlers.HandleCreateTask

	b.commands["/start"] = handlers.HandleStart
		"/delete_user"
		"/help"
		"/info"
		"/source"
} */

func (b *BotWithConfig) sendResponse(response tgbotapi.Chattable) {
	b.Bot.Send(response)
}
