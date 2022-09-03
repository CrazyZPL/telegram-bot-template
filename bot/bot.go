package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/telegram-bot-template/pkg/config"
	"github.com/telegram-bot-template/pkg/db"
	"github.com/telegram-bot-template/pkg/db/ent"
)

type MyBot struct {
	Bot         tgbotapi.BotAPI
	MysqlClient *ent.Client
}

// ConnectMyBot is a function to connect to our telegram bot
func ConnectMyBot(configPath string, needMySQL bool) {
	config, err := config.InitConfigFile(configPath)
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.Bot.BotToken)
	if err != nil {
		log.Panic(err)
	}

	myBot := MyBot{
		Bot: *bot,
	}

	if needMySQL {
		myBot.MysqlClient, err = db.InitMysqlClient(config.Mysql)
		if err != nil {
			log.Panic(err)
		}
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		var newMessage tgbotapi.MessageConfig
		var needSend bool
		newMessage.ParseMode = "MarkdownV2"
		if update.Message.IsCommand() {
			needSend = myBot.WhenUpdateIsCommand(update, &newMessage)
		} else if update.CallbackQuery != nil {
			needSend = myBot.WhenUpdateIsCallbackQuery(update, &newMessage)
		} else if update.Message != nil {
			needSend = myBot.WhenUpdateIsMessage(update, &newMessage)
		} else {
			continue
		}

		if needSend {
			if _, err := bot.Send(newMessage); err != nil {
				log.Printf("Send message error: %v", err)
			}
		}

		if _, err := bot.Send(newMessage); err != nil {
			panic(err)
		}
	}
}

func (bot *MyBot) WhenUpdateIsMessage(update tgbotapi.Update, newMessage *tgbotapi.MessageConfig) bool {
	return false
}

func (bot *MyBot) WhenUpdateIsCommand(update tgbotapi.Update, newMessage *tgbotapi.MessageConfig) bool {
	return false
}

func (bot *MyBot) WhenUpdateIsCallbackQuery(update tgbotapi.Update, newMessage *tgbotapi.MessageConfig) bool {
	return false
}

func (bot *MyBot) WhenUpdateIsKeyWord(update tgbotapi.Update, newMessage *tgbotapi.MessageConfig) bool {
	return false
}
