package handlers

import (
	"SnipSnapBot/internal/models"
	"SnipSnapBot/internal/services"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"gopkg.in/telebot.v4"
)

/// /save url <tags>
///росмотр ссылок
//o	/list — показать все ссылки пользователя с тегами
//o	/list <тег> — показать ссылки по конкретному тегу
//o	/get — получить случайную ссылку
//4.	Редактирование / удаление ссылок
//o	/delete <url> — удалить ссылку
//o	/edit <ID> <новый URL> <новые теги> — редактировать ссылку

type Handlers struct {
	bot     telebot.Bot
	service services.ILinkService
}

func (h *Handlers) SaveHandler(c telebot.Context) error {
	userID := int(c.Sender().ID)
	msg := c.Message().Text

	parts := h.parseMsg(msg)

	if len(parts) < 3 {
		return c.Reply("неверная команда. Пример: /save <url> <tags>")
	}

	url := parts[1] // /save is part[0]
	tags := parts[2:]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := h.service.Save(ctx, userID, url, tags); err != nil {
		return c.Reply("ошибка при сохранении!")
	}

	return c.Reply("Ссылка сохранена!✅")
}

func (h *Handlers) ListHandler(c telebot.Context) error {
	userID := int(c.Sender().ID)
	msg := c.Message().Text

	parts := h.parseMsg(msg)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// just /list
	if len(parts) == 1 {
		links, err := h.service.List(ctx, userID)
		if err != nil {
			log.Printf("list error: %v", err)
			return c.Reply("Ошибка при загрузке!")
		}
		if links == nil {
			return c.Reply("Ссылок нет :(")
		}

		return c.Reply(h.formatLinks(links))
	}
	// /list with tag
	if len(parts) == 2 {
		links, err := h.service.GetByTag(ctx, userID, parts[1])
		if err != nil {
			log.Printf("list error: %v", err)
			c.Reply("Ошибка при загрузке")
		}

		return c.Reply(h.formatLinks(links))
	}

	return c.Reply("Неверный формат команды.\nПравильно:\n/list\n/list <tag>")
}

func (h *Handlers) RemoveHandler(c telebot.Context) error {
	userID := int(c.Sender().ID)
	msg := c.Message().Text

	parts := h.parseMsg(msg)
	if len(parts) != 2 {
		return c.Reply("Неверный формат командыю \nПравильно:\n/remove\n/<url>")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := parts[1]

	err := h.service.Remove(ctx, userID, url)
	if err != nil {
		log.Printf("remove error: %v", err)
		c.Reply("Ошибка при удалении.")
	}

	return c.Reply("Ссылка удалена!")
}

func (h *Handlers) GetHandler(c telebot.Context) error {
	userID := int(c.Sender().ID)
	msg := c.Message().Text

	parts := h.parseMsg(msg)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if len(parts) == 1 {
		link, err := h.service.GetRandom(ctx, userID)
		if err != nil {
			log.Printf("get random error: %v", err)
			return c.Reply("Ошибка при загрузке!")
		}
		if link == nil {
			return c.Reply("ссылок нет :(")
		}

		c.Reply(h.formatLink(*link))
	}
	return c.Reply("Неверный формат команды.\nПравильно:\n/get\n/")
}

func (h *Handlers) parseMsg(msg string) []string {
	return strings.Fields(msg)
}

func (h *Handlers) formatLink(link models.Link) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(
		"URL: %sТеги: %s\n",
		link.URL,
		strings.Join(link.Tags, ","),
	))
	return b.String()
}

func (h *Handlers) formatLinks(links []models.Link) string {
	if len(links) == 0 {
		return "ничего не найдено"
	}

	var b strings.Builder

	for _, l := range links {
		b.WriteString(fmt.Sprintf(
			"URL: %s\nТеги: %s\n",
			l.URL,
			strings.Join(l.Tags, ","),
		))
	}

	return b.String()
}
