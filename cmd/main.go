package main

import (
	bot "github.com/aAmer0neee/Bot_Study-Helper/internal/TGbot"
	"github.com/aAmer0neee/Bot_Study-Helper/internal/TGbot/handlers"
	"github.com/aAmer0neee/Bot_Study-Helper/internal/cache"
	"github.com/aAmer0neee/Bot_Study-Helper/internal/config"
	"github.com/aAmer0neee/Bot_Study-Helper/internal/logger"
	"github.com/aAmer0neee/Bot_Study-Helper/internal/service"
)

func main() {

	cfg := config.LoadConfig()

	logger := logger.ConfigureLogger(cfg.Env)

	Cache := cache.NewRedisClient(cfg.CachePort)

	service := service.InitService(Cache)

	handler := handlers.InitHandlers(service)

	cfgBot := bot.ConfugureBot(cfg, logger, handler)

	cfgBot.ListenAndHandleUpdates()
}
