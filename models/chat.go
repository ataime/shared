package models

type Chat struct {
	ChatId string `json:"chat_id"`
	From   string `json:"from"`
	To     string `json:"to"`
}
