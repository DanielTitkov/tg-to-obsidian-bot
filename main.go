package main

import (
	"flag"
	"log"

	"github.com/DanielTitkov/tg-to-obsidian-bot/internal/bot"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var token = flag.String("token", "", "telegram bot token")
	var deleteTimeout = flag.Int("timeout", 180, "delete processed messages after n seconds")
	var notesPath = flag.String("path", "", "path to save notes")
	var templatePath = flag.String("template", "", "path to template folder")
	var templateFile = flag.String("file", "telegram.md", "template file")
	var debug = flag.Bool("debug", false, "run in debug mode (without saving files)")
	flag.Parse()

	if *token == "" {
		log.Fatal("telegram token not provided")
	}

	if *notesPath == "" {
		log.Fatal("notes path not provided")
	}

	tgBot, err := bot.NewBot(
		*token,
		*deleteTimeout,
		*notesPath,
		*templatePath+*templateFile,
		*debug,
	)
	if err != nil {
		log.Fatalf("failed to create bot: %s", err)
	}

	tgBot.Telebot.Handle(tb.OnText, tgBot.MessageToObsidianHandler)
	// tbBot.Telebot.Handle(tb.On)
	log.Println("Starting bot...")
	tgBot.Telebot.Start()
}
