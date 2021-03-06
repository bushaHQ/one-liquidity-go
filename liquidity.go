package liquidity

import (
	"fmt"
	"strings"
)

// url paths to the various endpoints
const (
	userEndpoint   = "/card/v1/user"
	getUserAddress = userEndpoint + "/address"
	getUserDoc     = userEndpoint + "/documentation/urls"
)

type RegisterIntegratorData struct {
	FloatCurrencies    []string `json:"floatCurrencies"`
	FirstName          string   `json:"firstName"`
	LastName           string   `json:"lastName"`
	Country            string   `json:"country"`
	BusinessName       string   `json:"businessName"`
	RegistrationNumber string   `json:"registrationNumber"`
	BusinessAddress    string   `json:"businessAddress"`
	Domain             string   `json:"domain"`
	Email              string   `json:"email"`
	WebhookUrl         string   `json:"webhookUrl"`
	ContactNumber      string   `json:"contactNumber"`
}

type CreateCardData struct {
	UserId    string `json:"userId"`
	Expiry    string `json:"expiry"`
	SingleUse bool   `json:"singleUse"`
}

type Params struct {
	Id        string
	Type      string
	StartDate string
	EndDate   string
	Limit     int
	Lek       string
}

type CreateUserData struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	KycCountry string `json:"kycCountry"`
	UID        string `json:"uid"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
}

type UpdateUserAddressData struct {
	UserID     string `json:"userId"`
	KycCountry string `json:"kycCountry"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
}

// RegisterIntegrator allows an integrator register with the system
func (cl *Client) RegisterIntegrator(data RegisterIntegratorData) (IntegratorResp, error) {
	var res IntegratorResp
	err := cl.post("/integrator/v1/register", data, &res)
	return res, err
}

// UpdateWebhook allows an integrator to update their webhook URL
func (cl *Client) UpdateWebhook(webhook string) (Resp, error) {
	var res Resp
	err := cl.patch("/integrator/v1/webhook", w{webhook}, &res)
	return res, err
}

//CreateCard allows an integrator to create a virtual card for their user
func (cl *Client) CreateCard(data CreateCardData) (CardResp, error) {
	var res CardResp
	err := cl.post("/card/v1", data, &res)
	return res, err
}

// GetCard allows an integrator to get full details of one card for their user
func (cl *Client) GetCard(card string, trackingNumber string) (CardResp, error) {
	var res CardResp
	err := cl.get(fmt.Sprintf("/card/v1?card=%s&trackingNumber=%s", card, trackingNumber), nil, &res)
	return res, err
}

// GetCards allows an integrator to get all cards for their user
func (cl *Client) GetCards(p Params) (CardsResp, error) {
	var res CardsResp
	err := cl.get(fmt.Sprintf("/cards/v1?user=%s&type=%s&startDate=%s&endDate=%s&limit=%d&lek=%s", p.Id, p.Type, p.StartDate, p.EndDate, p.Limit, p.Lek), nil, &res)
	return res, err
}

// TopUp allows an integrator to top up the card balance of a user
func (cl *Client) TopUp(cardId string, amount float64) (CardResp, error) {
	var res CardResp
	err := cl.patch("/card/v1/credit/balance", t{cardId, amount}, &res)
	return res, err
}

// Debit allows an integrator to deduct from the card balance of a user
func (cl *Client) Debit(cardId string, amount float64) (CardResp, error) {
	var res CardResp
	err := cl.patch("/card/v1/debit/balance", t{cardId, amount}, &res)
	return res, err
}

// Freeze allows an integrator or admin to freeze any type of card
func (cl *Client) Freeze(cardId string) (Resp, error) {
	var res Resp
	err := cl.patch("/card/v1/freeze", s{CardId: cardId}, &res)
	return res, err
}

// Unfreeze allows an integrator or admin to unfreeze any type of card
func (cl *Client) Unfreeze(cardId string) (Resp, error) {
	var res Resp
	err := cl.patch("/card/v1/unfreeze", s{CardId: cardId}, &res)
	return res, err
}

// StopCard allows an integrator to stop a card
func (cl *Client) StopCard(cardId string, reasonId int) (Resp, error) {
	var res Resp
	err := cl.patch("/card/v1/stop", s{cardId, reasonId}, &res)
	return res, err
}

// GetFailedTransaction returns the details of a failed transaction
func (cl *Client) GetFailedTransaction(txnId string) (TransactionResp, error) {
	var res TransactionResp
	err := cl.get(fmt.Sprintf("/card/v1/transaction/failed?transaction=%s", txnId), nil, &res)
	return res, err
}

// GetFailedTransactions returns the details of all failed transactions
func (cl *Client) GetFailedTransactions(p Params) (TransactionsResp, error) {
	var res TransactionsResp
	err := cl.get(fmt.Sprintf("/card/v1/transactions/failed?card=%s&startDate=%s&endDate=%s&limit=%d&lek=%s", p.Id, p.StartDate, p.EndDate, p.Limit, p.Lek), nil, &res)
	return res, err
}

// GetTransaction allows integrators to get a list of all transactions for a given card
func (cl *Client) GetTransaction(cardId string, p Params) (TransactionsResp, error) {
	var res TransactionsResp
	err := cl.get(fmt.Sprintf("/card/v1/transactions?card=%s&startDate=%s&endDate=%s&limit=%d&lek=%s", cardId, p.StartDate, p.EndDate, p.Limit, p.Lek), nil, &res)
	return res, err
}

// GetIntegratorDeposit allows an integrator retrieve a deposit
func (cl *Client) GetIntegratorDeposit(depositId string) (DepositResp, error) {
	var res DepositResp
	err := cl.get(fmt.Sprintf("/integrator/v1/deposit?deposit=%s", depositId), nil, &res)
	return res, err
}

// PostIntegratorDeposit allows an admin to update an integrator's deposit
func (cl *Client) PostIntegratorDeposit(amount int, currency string) (PostDepositResp, error) {
	var res PostDepositResp
	err := cl.post("/integrator/v1/deposit", d{amount, currency}, &res)
	return res, err
}

// GetIntegratorFloats retrieves an integrators list of float account balances for given array of currencies
func (cl *Client) GetIntegratorFloats(currencies []string) (FloatsResp, error) {
	var res FloatsResp
	bd := strings.Builder{}
	for idx, currency := range currencies {

		if idx == len(currencies)-1 {
			bd.WriteString("currencies=" + currency)
			break
		}

		bd.WriteString("currencies=" + currency + "&")
	}
	err := cl.get(fmt.Sprintf("/integrator/v1/floats?%s", bd.String()), nil, &res)
	return res, err
}

// GetIntegratorFloat retrieves an integrator's float account balance for a given currency
func (cl *Client) GetIntegratorFloat(currency string) (FloatResp, error) {
	var res FloatResp
	err := cl.get(fmt.Sprintf("/integrator/v1/float?currency=%s", currency), nil, &res)
	return res, err
}

// UpdateFloatDefault allows an integrator to update their default float
func (cl *Client) UpdateFloatDefault(floatId string) (Resp, error) {
	var res Resp
	err := cl.patch("/integrator/v1/float/default", f{floatId}, &res)
	return res, err
}

// GetUser Users allows an integrator to create a user
func (cl *Client) GetUser(userID string) (getUserResp, error) {
	var res getUserResp
	err := cl.get(fmt.Sprintf("%s?userId=%s", userEndpoint, userID), nil, &res)
	return res, err
}

// CreateUser Users allows an integrator to create a user
func (cl *Client) CreateUser(userData CreateUserData) (createUserResp, error) {
	var res createUserResp
	err := cl.post(userEndpoint, userData, &res)
	return res, err
}

// UpdateUserAdress allows an integrator to update address, postal code and KYC country
func (cl *Client) UpdateUserAddress(updateData UpdateUserAddressData) (updateUserAddressResp, error) {
	var res updateUserAddressResp
	err := cl.patch(getUserAddress, updateData, &res)
	return res, err
}

// GetCardUserDocURL allows an integrator to update address, postal code and KYC country
func (cl *Client) GetCardUserDocURL(userID string) (getCardUserDocURLResp, error) {
	var res getCardUserDocURLResp
	err := cl.get(fmt.Sprintf("%s?user=%s", getUserDoc, userID), nil, &res)
	return res, err
}
