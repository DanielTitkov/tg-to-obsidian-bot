package bot

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/DanielTitkov/tg-to-obsidian-bot/internal/markdown"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	datetimeFormat = "2006-01-02 15:04:05"
)

type Bot struct {
	Telebot   *tb.Bot
	Timeout   int
	NotesPath string
}

func NewBot(
	token string,
	deleteTimeout int,
	notesPath string,
) (*Bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	return &Bot{
		Telebot:   b,
		Timeout:   deleteTimeout,
		NotesPath: notesPath,
	}, nil
}

func (b *Bot) MessageToObsidianHandler(m *tb.Message) {
	noteText, err := markdown.WrapWithMarkdown(m.Text)
	if err != nil {
		errMsg := fmt.Sprintf("failed to convert message to note: %s", err)
		b.Telebot.Send(m.Sender, errMsg)
		log.Println(errMsg)
		return
	}

	log.Printf("generated markdown:\n%s\n", noteText)

	filePath := b.NotesPath + fmt.Sprintf("telegram note %s.md", time.Now().Format(datetimeFormat))
	err = ioutil.WriteFile(filePath, []byte(noteText), 0644)
	if err != nil {
		errMsg := fmt.Sprintf("failed to save file: %s", err)
		b.Telebot.Send(m.Sender, errMsg)
		log.Println(errMsg)
		return
	}
	log.Printf("saved to file: %s", filePath)

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
