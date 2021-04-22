package jobs

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func Start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			if update.ChannelPost != nil {
				fileNames := GetFiles()
				if update.ChannelPost.Text == "sent" || update.ChannelPost.Text == "Sent" {
					for _, fileName := range fileNames {
						msg := tgbotapi.NewDocumentUpload(update.ChannelPost.Chat.ID, "../bot_files/"+fileName)
						bot.Send(msg)
					}
				}
			}
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bot work at the channel"+os.Getenv("CHANNEL_NAME"))
		bot.Send(msg)
	}
}
