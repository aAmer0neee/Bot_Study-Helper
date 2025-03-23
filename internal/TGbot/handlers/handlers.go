package handlers

import (
	_ "log/slog"

	"github.com/aAmer0neee/Bot_Study-Helper/internal/service"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	MenuButtons = [][]string{
		{"📋 Список задач", "➕ Добавить задачу"},
		{"⚙ Настройки", "❓ Помощь"},
	}

	TaskMenu = [][]string{
		{"✏️ Редактировать", "❌ Удалить"},
	}
)

type (
	MessageContent struct {
		Text   string
		Markup tgbotapi.InlineKeyboardMarkup
	}
	Handler struct {
		Service *service.Service
	}
)

func Commandhandler(){
	
}

/* func UserState(u tgbotapi.Update){
	h.Service.
}

func (h *Handler) HandleCreateTask(u tgbotapi.Update) tgbotapi.Chattable {
	h.Service.CreateTaskService(u.Message.From.ID)
	return
} */

func (h *Handler) HandleStart(u tgbotapi.Update) tgbotapi.Chattable {

	h.Service.StartService(u.Message.From.ID)
	return createMessage(u.Message.Chat.ID, MessageContent{
		Text: "Hello",
		Markup:createInlineKeyboard(MenuButtons),})
}

func InitHandlers(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func createMessage(chatID int64, c MessageContent) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, c.Text)
	msg.ReplyMarkup = c.Markup
	return msg
}

func editMessage(chatID int64, messageID int, c MessageContent) tgbotapi.EditMessageTextConfig {
	return tgbotapi.NewEditMessageTextAndMarkup(
		chatID,
		messageID,
		c.Text,
		c.Markup,
	)
}

func createInlineKeyboard(buttons [][]string) tgbotapi.InlineKeyboardMarkup {
	var markup [][]tgbotapi.InlineKeyboardButton
	for _, row := range buttons {
		var markupRow []tgbotapi.InlineKeyboardButton
		for _, botton := range row {
			markupRow = append(markupRow, tgbotapi.NewInlineKeyboardButtonData(botton, botton))
		}
		markup = append(markup, markupRow)
	}
	return tgbotapi.NewInlineKeyboardMarkup(markup...)
}
