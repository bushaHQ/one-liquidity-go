package liquidity

import "time"

type IntegratorResp struct {
	Message string `json:"message"`
	Data    D1     `json:"data"`
}

type Resp struct {
	Message string `json:"message"`
}

type CardResp struct {
	Message string `json:"message"`
	Data    D2     `json:"data"`
}

type CardsResp struct {
	Message string `json:"message"`
	Data    []D2   `json:"data"`
}

type DepositResp struct {
	Message string `json:"message"`
	Data    D3     `json:"data"`
}

type TransactionResp struct {
	Message string `json:"message"`
	Data    D4     `json:"data"`
}

type TransactionsResp struct {
	Message string `json:"message"`
	Data    []D4   `json:"data"`
}

type D1 struct {
	IntegratorId string `json:"integratorId"`
}

type D2 struct {
	CardId         string    `json:"cardId"`
	Expiry         time.Time `json:"expiry"`
	Valid          string    `json:"valid"`
	Cvv2           string    `json:"cvv2"`
	CardNumber     string    `json:"cardNumber"`
	Last4          string    `json:"last4"`
	TrackingNumber string    `json:"trackingNumber"`
	Balance        int       `json:"balance"`
	Status         string    `json:"status,omitempty"`
	Currency       string    `json:"currency"`
	SingleUse      bool      `json:"singleUse"`
	CardName       string    `json:"cardName"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}

type D3 struct {
	U54DepositId string `json:"u54DepositId"`
	Amount       int    `json:"amount"`
	Currency     string `json:"currency"`
	Status       string `json:"status"`
}

type D4 struct {
	TransactionId            string `json:"transactionId"`
	DebitId                  string `json:"debitId"`
	DebitCurrency            string `json:"debitCurrency"`
	ConversionRate           int    `json:"conversionRate"`
	CreditCurrency           string `json:"creditCurrency"`
	TransactionBalanceBefore int    `json:"transactionBalanceBefore,omitempty"`
	CardBalanceAfter         int    `json:"cardBalanceAfter,omitempty"`
	CardId                   string `json:"cardId,omitempty"`
	Type                     string `json:"type"`
	Amount                   int    `json:"amount"`
	Currency                 string `json:"currency,omitempty"`
	ErrorDescription         string `json:"errorDescription,omitempty"`
	CreatedAt                string `json:"createdAt"`
	Narrative                string `json:"narrative"`
	AcquiringInstitutionCode string `json:"acquiringInstitutionCode"`
}

type t struct {
	CardId string
	Amount float64
}

type d struct {
	Amount   int
	Currency string
}

type f struct {
	CardId string
}

type s struct {
	CardId   string
	ReasonId int
}

type w struct {
	Webhook string
}
