// telegramBot.go

package main

import (
	tele "gopkg.in/telebot.v3"
	"os"
	"strconv"
	"strings"
	"time"
)

type TeleBot struct {
	Bot   *tele.Bot
	Chats []*tele.Chat
}

func (t *TeleBot) sendMsg(msg string) error {
	for _, chat := range t.Chats {
		_, err := t.Bot.Send(chat, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewTeleBot() (*TeleBot, error) {
	pref := tele.Settings{
		Token:  os.Getenv("TELEBOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	chatsRaw := os.Getenv("TELEBOT_CHATS")
	chatIDs := strings.Split(chatsRaw, ",")
	chats := make([]*tele.Chat, len(chatIDs))

	for i, rawID := range chatIDs {
		id, err := strconv.ParseInt(rawID, 10, 64)
		if err != nil {
			return nil, err
		}
		chats[i] = &tele.Chat{ID: id}
	}

	return &TeleBot{
		Bot:   b,
		Chats: chats,
	}, nil
}
