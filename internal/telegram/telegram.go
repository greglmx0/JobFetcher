package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
    bot *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (*TelegramBot, error) {
    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        return nil, err
    }
    bot.Debug = true
    return &TelegramBot{bot: bot}, nil
}

func (tb *TelegramBot) SendMessage(chatID int64, message string) error {
				log.Printf("Envoi du message: %s", message)
    msg := tgbotapi.NewMessage(chatID, message)
    _, err := tb.bot.Send(msg)

				if err != nil {
					log.Fatalf("Erreur lors de l'envoi du message: %v", err)
								return err
				}

				log.Printf("Message envoyé: %s", message)


    return err
}
