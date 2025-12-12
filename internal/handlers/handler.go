package handlers

import (
	"SnipSnapBot/internal/services"
	"context"
	"strings"
	"time"

	"gopkg.in/telebot.v4"
)

/// /save url <tags>
///росмотр ссылок
//o	/list — показать все ссылки пользователя с тегами
//o	/list <тег> — показать ссылки по конкретному тегу
//o	/get — получить случайную ссылку
//o	/get <тег> — получить случайную ссылку по тегу
//4.	Редактирование / удаление ссылок
//o	/delete <ID> — удалить ссылку
//o	/edit <ID> <новый URL> <новые теги> — редактировать ссылку

type Handlers struct {
	bot     telebot.Bot
	service services.ILinkService
}

func (h *Handlers) SaveHandler(c telebot.Context) error {
	userID := int(c.Sender().ID)
	msg := c.Message().Text

	parts := strings.Fields(msg)

	if len(parts) < 3 {
		return c.Reply("неверная команда. Пример: /save <url> <tags>")
	}

	url := parts[1] // /save is part[0]
	tags := parts[2:]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := h.service.Save(ctx, userID, url, tags); err != nil {
		return c.Reply("ошибка при сохранении! + ", err.Error())
	}

	return c.Reply("Ссылка сохранена!✅")
}
