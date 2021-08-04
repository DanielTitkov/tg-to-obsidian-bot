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
	flag.Parse()

	if *token == "" {
		log.Fatal("telegram token not provided")
	}

	tgBot, err := bot.NewBot(
		*token,
		*deleteTimeout,
	)
	if err != nil {
		log.Fatalf("failed to create bot: %s", err)
	}

	tgBot.Telebot.Handle(tb.OnText, tgBot.MessageToObsidianHandler)
	log.Println("Starting bot...")
	tgBot.Telebot.Start()
}
