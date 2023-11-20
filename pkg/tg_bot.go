package tg_bot

import "Isotere/tg-bot/pkg/internal/clients/telegram"

type TgBot struct {
	tgClient TgClient
}

func New(options Options) *TgBot {
	tgClient := telegram.New(options.Host, options.Token)

	return &TgBot{
		tgClient: tgClient,
	}
}
