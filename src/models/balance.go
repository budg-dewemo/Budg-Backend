package models

type IBalance interface {
	GetBalance(date string) []Balance
}

type Balance struct {
	Date   string `json:"date"`
	Income int    `json:"income"`

	Transactions []Transaction `json:"transactions"`
}
