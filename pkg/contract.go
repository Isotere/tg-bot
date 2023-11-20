package tg_bot

import "Isotere/tg-bot/pkg/internal/clients/telegram"

//go:generate mockgen -source=contract.go -destination contract_mocks_test.go -package $GOPACKAGE

type TgClient interface {
	SendMessage(chatID int, text string) error
	Updates(offset int, limit int) ([]telegram.Update, error)
}
