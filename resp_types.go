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

type D1 struct {
	IntegratorId string `json:"integratorId"`
}

type CardsResp struct {
	Message string `json:"message"`
	Data    []D2   `json:"data"`
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
	U54DepositId string    `json:"u54DepositId,omitempty"`
	DepositId    string    `json:"depositId,omitempty"`
	Amount       int       `json:"amount"`
	Currency     string    `json:"currency"`
	CreatedAt    time.Time `json:"createdAt"`
	Usd          Usd       `json:"usd"`
	Btc          Coin      `json:"btc"`
	Eth          Coin      `json:"eth"`
	Busd         Coin      `json:"busd"`
	Usdc         Coin      `json:"usdc"`
	Usdt         Coin      `json:"usdt"`
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
	FloatId   string    `json:"floatId"`
	UpdatedAt time.Time `json:"updatedAt"`
	Currency  string    `json:"currency"`
	Balance   int       `json:"balance"`
	IsDefault bool      `json:"isDefault"`
}

type t struct {
	cardId string
	amount float64
}

type d struct {
	amount   int
	currency string
}

type f struct {
	id string
}

type s struct {
	cardId   string
	reasonId int
}

type w struct {
	webhook string
}

type getUserResp struct {
	Message string `json:"message"`
	Data    gur    `json:"data"`
}

type gur struct {
	CreatedAt         time.Time `json:"createdAt,string"`
	UpdatedAt         time.Time `json:"updatedAt,string"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	UID               string    `json:"uid"`
	KycCountry        string    `json:"kycCountry"`
	Address           string    `json:"address"`
	City              string    `json:"city"`
	PostalCode        string    `json:"postalCode"`
	PhysicalCardCount int       `json:"physicalCardCount"`
	VirtualCardCount  int       `json:"virtualCardCount"`
	SelfieUploaded    bool      `json:"selfieUploaded"`
	IDUploaded        bool      `json:"idUploaded"`
	OfacChecked       bool      `json:"ofacChecked"`
	OfacFail          bool      `json:"ofacFail"`
	Active            bool      `json:"active"`
}

type createUserResp struct {
	Message string   `json:"message"`
	Data    userResp `json:"data"`
}
type userResp struct {
	UserID string `json:"userId"`
}

type getCardUserDocURLResp struct {
	Message string     `json:"message"`
	Data    docURLResp `json:"data"`
}

type updateUserAddressResp struct {
	Message string `json:"message"`
	Data    struct {
		Message string `json:"message"`
	} `json:"data"`
}
type validationError struct {
	Code     string   `json:"code"`
	Keys     []string `json:"keys"`
	Expected string   `json:"expected"`
	Received string   `json:"received"`
	Path     []string `json:"path"`
	Message  string   `json:"message"`
}

type docURLResp struct {
	SelfieUploadURL string `json:"selfieUploadUrl"`
	IDUploadURL     string `json:"idUploadUrl"`
	UID             string `json:"uid"`
}
