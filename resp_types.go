package liquidity

type IntegratorResp struct {
	Message string `json:"message"`
	Data    D1     `json:"data"`
}

type D1 struct {
	IntegratorId string `json:"integratorId"`
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

type D2 struct {
	CardId         string `json:"cardId"`
	Expiry         string `json:"expiry"`
	Valid          string `json:"valid"`
	Cvv2           string `json:"cvv2"`
	CardNumber     string `json:"cardNumber,omitempty"`
	Last4          string `json:"last4"`
	TrackingNumber string `json:"trackingNumber"`
	Balance        int    `json:"balance"`
	Status         string `json:"status,omitempty"`
	Currency       string `json:"currency"`
	SingleUse      bool   `json:"singleUse"`
	CardName       string `json:"cardName"`
	CreatedAt      string `json:"createdAt,omitempty"`
}

type DepositResp struct {
	Message string `json:"message"`
	Data    D3     `json:"data"`
}

type D3 struct {
	DepositId    string `json:"depositId,omitempty"`
	U54DepositId string `json:"u54DepositId,omitempty"`
	Amount       int    `json:"amount"`
	Currency     string `json:"currency"`
	Status       string `json:"status"`
}

type TransactionResp struct {
	Message string `json:"message"`
	Data    D4     `json:"data"`
}

type TransactionsResp struct {
	Message string `json:"message"`
	Data    []D4   `json:"data"`
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

type PostDepositResp struct {
	Message string `json:"message"`
	Data    D5     `json:"data"`
}

type D5 struct {
	U54DepositId string `json:"u54DepositId,omitempty"`
	DepositId    string `json:"depositId,omitempty"`
	Amount       int    `json:"amount"`
	Currency     string `json:"currency"`
	CreatedAt    string `json:"createdAt"`
	Usd          Usd    `json:"usd"`
	Btc          Coin   `json:"btc"`
	Eth          Coin   `json:"eth"`
	Busd         Coin   `json:"busd"`
	Usdc         Coin   `json:"usdc"`
	Usdt         Coin   `json:"usdt"`
}

type Usd struct {
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	BankName      string `json:"bankName"`
	BankAddress   string `json:"bankAddress"`
	BranchCode    string `json:"branchCode"`
	SwiftCode     string `json:"swiftCode"`
}

type Coin struct {
	WalletAddress string `json:"walletAddress"`
}

type FloatsResp struct {
	Message string `json:"message"`
	Data    []D6   `json:"data"`
}

type FloatResp struct {
	Message string `json:"message"`
	Data    D6     `json:"data"`
}

type D6 struct {
	FloatId   string `json:"floatId"`
	UpdatedAt string `json:"updatedAt"`
	Currency  string `json:"currency"`
	Balance   int    `json:"balance"`
	IsDefault bool   `json:"isDefault"`
}

type t struct {
	CardId string  `json:"cardId"`
	Amount float64 `json:"amount"`
}

type d struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type f struct {
	Id string `json:"id"`
}

type s struct {
	CardId   string `json:"cardId"`
	ReasonId int    `json:"reasonId"`
}

type w struct {
	Webhook string `json:"webhook"`
}
