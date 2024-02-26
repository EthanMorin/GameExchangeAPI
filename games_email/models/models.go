package models

type UserEmail struct {
	Email *string `json:"email,omitempty"`
}

type EchangeEmails struct {
	TradeeEmail *string `json:"tradee_email,omitempty"`
	TraderEmail *string `json:"trader_email,omitempty"`
}
