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
	titleFormat    = "Telegram Note %s"
)

type Bot struct {
	Telebot      *tb.Bot
	Timeout      int
	NotesPath    string
	TemplatePath string
	Debug        bool
}

func NewBot(
	token string,
	deleteTimeout int,
	notesPath string,
	templatePath string,
	debug bool,
) (*Bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	return &Bot{
		Telebot:      b,
		Timeout:      deleteTimeout,
		NotesPath:    notesPath,
		TemplatePath: templatePath,
		Debug:        debug,
	}, nil
}

func (b *Bot) MessageToObsidianHandler(m *tb.Message) {
	// this is run every time for in case template is changed
	template, err := ioutil.ReadFile(b.TemplatePath)
	if err != nil {
		errMsg := fmt.Sprintf("failed to load template: %s", err)
		b.Telebot.Send(m.Sender, errMsg)
		log.Println(errMsg)
		return
	}

	text, title, err := markdown.ExtractTitle(m.Text)
	if err != nil {
		errMsg := fmt.Sprintf("failed to extract title: %s", err)
		b.Telebot.Send(m.Sender, errMsg)
		log.Println(errMsg)
		return
	}

	if title == "" {
		title = fmt.Sprintf(titleFormat, time.Now().Format(datetimeFormat))
	}

	noteText, err := markdown.WrapWithMarkdown(text, string(template), title)
	if err != nil {
		errMsg := fmt.Sprintf("failed to convert message to note: %s", err)
		b.Telebot.Send(m.Sender, errMsg)
		log.Println(errMsg)
		return
	}

	log.Printf("\ngenerated markdown:\n---\n%s\n---\n", noteText)

	filePath := b.NotesPath + fmt.Sprintf("%s.md", title)
	if !b.Debug {
		err = ioutil.WriteFile(filePath, []byte(noteText), 0644)
		if err != nil {
			errMsg := fmt.Sprintf("failed to save file: %s", err)
			b.Telebot.Send(m.Sender, errMsg)
			log.Println(errMsg)
			return
		}

		log.Printf("saved to file: %s", filePath)
	} else {
		log.Printf("running in debug mode, saving omitted, generated path: %s", filePath)
	}

	reply, err := b.Telebot.Send(
		m.Sender,
		fmt.Sprintf("Сообщение обработано! Удалим через %d секунд.\n\nСоздана заметка:\n%s", b.Timeout, noteText),
	)
	if err != nil {
		log.Println("failed to send reply", err)
	}

	go func() {
		time.Sleep(time.Second * time.Duration(b.Timeout))
		b.Telebot.Delete(m)
		b.Telebot.Delete(reply)
	}()
}
