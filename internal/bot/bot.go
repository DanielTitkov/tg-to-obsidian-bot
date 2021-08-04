package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielTitkov/tg-to-obsidian-bot/internal/markdown"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Bot struct {
	Telebot *tb.Bot
	Timeout int
}

func NewBot(
	token string,
	deleteTimeout int,
) (*Bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	return &Bot{
		Telebot: b,
		Timeout: deleteTimeout,
	}, nil
}

func (b *Bot) MessageToObsidianHandler(m *tb.Message) {
	noteText, err := markdown.WrapWithMarkdown(m.Text)
	if err != nil {
		b.Telebot.Send(m.Sender, fmt.Sprintf("failed to convert message to note: %s", err))
	}
	log.Printf("generated markdown:\n%s\n", noteText)
	reply, err := b.Telebot.Send(m.Sender, fmt.Sprintf("Сообщение обработано! Удалим через %d секунд", b.Timeout))
	if err != nil {
		log.Println("failed to send reply", err)
	}

	go func() {
		time.Sleep(time.Second * time.Duration(b.Timeout))
		b.Telebot.Delete(m)
		b.Telebot.Delete(reply)
	}()
}
