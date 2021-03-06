package liquidity

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var (
	cl *Client
)

type MockHttpClient struct {
	DoFunc  func(req *http.Request) (*http.Response, error)
	Timeout time.Duration
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func init() {
	err := godotenv.Load("./.env")
	if err != nil && os.Getenv("ENV") == "" {
		panic(err)
	}
	cl = NewClient()
}

func TestClient_RegisterIntegrator(t *testing.T) {
	type args struct {
		data RegisterIntegratorData
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           IntegratorResp
		wantErr        bool
	}{
		{
			name: "allows an integrator register with the system",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/integrator/v1/register" {
						t.Errorf("Expected to request '/integrator/v1/register', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok","data":{"integratorId":"863494e2-40b3-44dc-8acf-5ee520097f75"}}`)))

					return &http.Response{
						StatusCode: 201,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				data: RegisterIntegratorData{
					FloatCurrencies:    []string{"USD", "BTC"},
					FirstName:          "Olusola",
					LastName:           "Alao",
					Country:            "NGA",
					BusinessName:       "Algo Math",
					RegistrationNumber: "12345678",
					BusinessAddress:    "Lekki Ikate",
					Domain:             "olusola.tech",
					Email:              "justiceoyin@gmail.com",
					WebhookUrl:         "https://webhook.site/d8e81cdd-0db9-4b10-82a0-54f8d6be247f",
					ContactNumber:      "+2349034345678",
				},
			},
			want: IntegratorResp{
				Message: "Ok",
				Data: D1{
					IntegratorId: "863494e2-40b3-44dc-8acf-5ee520097f75",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.RegisterIntegrator(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterIntegrator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterIntegrator() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateWebhook(t *testing.T) {
	type args struct {
		webhook string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           Resp
		wantErr        bool
	}{
		{
			name: "allows an integrator to update their webhook url",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/integrator/v1/webhook" {
						t.Errorf("Expected to request '/integrator/v1/webhook', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				webhook: "https://webhook.site",
			},
			want: Resp{
				Message: "Ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.UpdateWebhook(tt.args.webhook)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWebhook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateWebhook() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateCard(t *testing.T) {

	type args struct {
		data CreateCardData
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           CardResp
		wantErr        bool
	}{
		{
			name: "can allow an integrator to create a virtual card for their user",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1" {
						t.Errorf("Expected to request '/card/v1', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{
    "message": "Ok",
    "data": {
        "expiry": "2025-04-19T00:00:00.000Z",
        "valid": "04/25",
        "cvv2": "142",
        "last4": "0083",
        "cardNumber": "5368989511270083",
        "trackingNumber": "147203800064758",
        "currency": "USD",
        "balance": 0,
        "singleUse": false,
        "cardId": "c954e4c9-8ca8-4d1d-8ebf-bb374e9f1b1d",
        "cardName": "Sofiyu Soft"
    }
}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				data: CreateCardData{
					UserId:    "e08078bd-9384-5b7e-93c5-76be956380fe",
					Expiry:    "2025-04-18",
					SingleUse: false,
				},
			},
			want: CardResp{
				Message: "Ok",
				Data: D2{
					CardId:         "c954e4c9-8ca8-4d1d-8ebf-bb374e9f1b1d",
					Expiry:         "2025-04-19T00:00:00.000Z",
					Valid:          "04/25",
					Cvv2:           "142",
					CardNumber:     "5368989511270083",
					Last4:          "0083",
					TrackingNumber: "147203800064758",
					Balance:        0,
					Currency:       "USD",
					SingleUse:      false,
					CardName:       "Sofiyu Soft",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.CreateCard(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//e08078bd-9384-5b7e-93c5-76be956380fe

func TestClient_GetCard(t *testing.T) {

	type args struct {
		card           string
		trackingNumber string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           CardResp
		wantErr        bool
	}{
		{
			name: "an integrator can get full details of one card for their user",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1" {
						t.Errorf("Expected to request '/card/v1', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{
  "message": "Ok",
  "data": {
    "cardId": "c954e4c9-8ca8-4d1d-8ebf-bb374e9f1b1d",
    "userId": "84055c2c-7e07-550e-9dec-533f6f52b54d",
    "expiry": "2025-04-19T00:00:00.000Z",
    "valid": "04/25",
    "cvv2": "142",
    "last4": "0083",
    "balance": 0,
    "status": "issued",
    "currency": "USD",
    "singleUse": false,
    "createdAt": "2022-05-27T08:46:37.549Z",
    "cardName": "Sofiyu Soft",
    "cardNumber": "5368989511270083",
    "trackingNumber": "147203800064758"
  }
}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				card:           "c954e4c9-8ca8-4d1d-8ebf-bb374e9f1b1d",
				trackingNumber: "147203800064758",
			},
			want: CardResp{
				Message: "Ok",
				Data: D2{
					CardId:         "c954e4c9-8ca8-4d1d-8ebf-bb374e9f1b1d",
					Expiry:         "2025-04-19T00:00:00.000Z",
					Valid:          "04/25",
					Cvv2:           "142",
					CardNumber:     "5368989511270083",
					Last4:          "0083",
					TrackingNumber: "147203800064758",
					Balance:        0,
					Status:         "issued",
					Currency:       "USD",
					SingleUse:      false,
					CardName:       "Sofiyu Soft",
					CreatedAt:      "2022-05-27T08:46:37.549Z",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetCard(tt.args.card, tt.args.trackingNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetCards(t *testing.T) {

	type args struct {
		p Params
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           CardsResp
		wantErr        bool
	}{
		{
			name: "allows integrators to get a list of cards for their users",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/cards/v1" {
						t.Errorf("Expected to request '/cards/v1', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{
  "message": "Ok",
  "data": [
    {
      "cardId": "922a54cb-0072-429c-9313-045c8fc09c64",
      "userId": "e08078bd-9384-5b7e-93c5-76be956380fe",
      "expiry": "2025-04-19T00:00:00.000Z",
      "valid": "04/25",
      "cvv2": "372",
      "last4": "5761",
      "balance": 0,
      "status": "issued",
      "currency": "USD",
      "cardType": "virtual",
      "singleUse": false,
      "createdAt": "2022-05-27T09:29:54.209Z",
      "cardName": "Chijioke Amanambu",
      "trackingNumber": "986003800064765"
    }
  ]}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				p: Params{
					Id: "e08078bd-9384-5b7e-93c5-76be956380fe",
				},
			},
			want: CardsResp{
				Message: "Ok",
				Data: []D2{
					{
						CardId:         "922a54cb-0072-429c-9313-045c8fc09c64",
						Expiry:         "2025-04-19T00:00:00.000Z",
						Valid:          "04/25",
						Cvv2:           "372",
						Last4:          "5761",
						TrackingNumber: "986003800064765",
						Balance:        0,
						Status:         "issued",
						Currency:       "USD",
						SingleUse:      false,
						CardName:       "Chijioke Amanambu",
						CreatedAt:      "2022-05-27T09:29:54.209Z",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetCards(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCards() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_TopUp(t *testing.T) {

	type args struct {
		cardId string
		amount float64
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           CardResp
		wantErr        bool
	}{
		{
			name: "Allows an integrator to top-up the card balance of a user",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/credit/balance" {
						t.Errorf("Expected to request '/card/v1/credit/balance', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				cardId: "c954e4c9-8ca8-4d1d-8ebf-bb374e9f1b1d",
				amount: 1000,
			},
			want: CardResp{
				Message: "Ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.TopUp(tt.args.cardId, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("TopUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Debit(t *testing.T) {

	type args struct {
		cardId string
		amount float64
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           CardResp
		wantErr        bool
	}{
		{
			name: "can allow an integrator to deduct from the card balance of a user",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/debit/balance" {
						t.Errorf("Expected to request '/card/v1/debit/balance', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				cardId: "c954e4c9-8ca8-4d1d-8ebf-bb374e9f1b1d",
				amount: 10,
			},
			want: CardResp{
				Message: "Ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.Debit(tt.args.cardId, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Debit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Debit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Freeze(t *testing.T) {

	type args struct {
		CardID string `json:"cardId"`
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           Resp
		wantErr        bool
	}{
		{
			name: "integrator or admin can freeze any type of card",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/freeze" {
						t.Errorf("Expected to request '/card/v1/freeze', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				CardID: "aa174033-fe13-4c3a-90b3-f3485a0e9c86",
			},
			want: Resp{
				Message: "Ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.Freeze(tt.args.CardID)
			if (err != nil) != tt.wantErr {
				log.Println(err)
				t.Errorf("Freeze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freeze() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Unfreeze(t *testing.T) {

	type args struct {
		cardId string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           Resp
		wantErr        bool
	}{
		{
			name: "Allow an integrator or admin to unfreeze any type of card",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/unfreeze" {
						t.Errorf("Expected to request '/card/v1/unfreeze', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				cardId: "aa174033-fe13-4c3a-90b3-f3485a0e9c86",
			},
			want: Resp{
				Message: "Ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.Unfreeze(tt.args.cardId)
			if (err != nil) != tt.wantErr {
				log.Println(err)
				t.Errorf("Unfreeze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unfreeze() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StopCard(t *testing.T) {

	type args struct {
		cardId   string
		reasonId int
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           Resp
		wantErr        bool
	}{
		{
			name: "Allows a card to be stopped",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/stop" {
						t.Errorf("Expected to request '/card/v1/stop', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				cardId:   "",
				reasonId: 0,
			},
			want: Resp{
				Message: "Ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.StopCard(tt.args.cardId, tt.args.reasonId)
			if (err != nil) != tt.wantErr {
				t.Errorf("StopCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StopCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetFailedTransaction(t *testing.T) {

	type args struct {
		txnId string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           TransactionResp
		wantErr        bool
	}{
		{
			name: "allows an integrator to update their webhook url",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/transaction/failed" {
						t.Errorf("Expected to request '/card/v1/transaction/failed', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				txnId: "",
			},
			want: TransactionResp{
				Message: "",
				Data: D4{
					TransactionId:            "",
					DebitId:                  "",
					DebitCurrency:            "",
					ConversionRate:           0,
					CreditCurrency:           "",
					TransactionBalanceBefore: 0,
					CardBalanceAfter:         0,
					CardId:                   "",
					Type:                     "",
					Amount:                   0,
					Currency:                 "",
					ErrorDescription:         "",
					CreatedAt:                "",
					Narrative:                "",
					AcquiringInstitutionCode: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetFailedTransaction(tt.args.txnId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFailedTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFailedTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetFailedTransactions(t *testing.T) {

	type args struct {
		p Params
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           TransactionsResp
		wantErr        bool
	}{
		{
			name: "Allows integrators to get a list of all failed transactions for a given card",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/transactions/failed" {
						t.Errorf("Expected to request '/card/v1/transactions/failed', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{
							"message": "Ok",
							"data": []
						  }`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				p: Params{
					Id:        "aa174033-fe13-4c3a-90b3-f3485a0e9c86",
					Type:      "",
					StartDate: "",
					EndDate:   "",
					Limit:     20,
					Lek:       "",
				},
			},
			want: TransactionsResp{
				Message: "Ok",
				Data:    []D4{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetFailedTransactions(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFailedTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFailedTransactions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetIntegratorDeposit(t *testing.T) {

	type args struct {
		depositId string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           DepositResp
		wantErr        bool
	}{
		{
			name: "allows an integrator to update their webhook url",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/integrator/v1/deposit" {
						t.Errorf("Expected to request '/integrator/v1/deposit', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				depositId: "",
			},
			want: DepositResp{
				Message: "",
				Data: D3{
					DepositId:    "",
					U54DepositId: "",
					Amount:       0,
					Currency:     "",
					Status:       "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetIntegratorDeposit(tt.args.depositId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntegratorDeposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntegratorDeposit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetIntegratorFloat(t *testing.T) {

	type args struct {
		currency string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           FloatResp
		wantErr        bool
	}{
		{
			name: "allows an integrator to update their webhook url",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/integrator/v1/float" {
						t.Errorf("Expected to request '/integrator/v1/float', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				currency: "",
			},
			want: FloatResp{
				Message: "",
				Data: D6{
					FloatId:   "",
					UpdatedAt: "",
					Currency:  "",
					Balance:   0,
					IsDefault: false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetIntegratorFloat(tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntegratorFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntegratorFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetIntegratorFloats(t *testing.T) {

	type args struct {
		currencies []string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           FloatsResp
		wantErr        bool
	}{
		{
			name: "allows an integrator to update their webhook url",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/integrator/v1/floats" {
						t.Errorf("Expected to request '/integrator/v1/floats', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{"message":"Ok"}`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				currencies: []string{},
			},
			want: FloatsResp{
				Message: "",
				Data:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetIntegratorFloats(tt.args.currencies)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntegratorFloats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntegratorFloats() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateUser(t *testing.T) {
	type args struct {
		CreateUserData
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           createUserResp
		wantErr        bool
	}{
		{
			name: "Create a card user",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/user" {
						t.Errorf("Expected to request '/card/v1/user', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{
							"message": "Ok",
							"data": {
							  "userId": "69a9a77b-5d8d-5738-80eb-ff0b1fb3846a"
							}
						  }`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				CreateUserData{
					FirstName:  "kasumu",
					LastName:   "sofiyullahi",
					KycCountry: "nga",
					UID:        "169c9ff3-4af1-4115-a7eb-363827022625",
					Address:    "3 Misratah Street, Wuse 2",
					City:       "fct",
					PostalCode: "900888",
				},
			},
			want: createUserResp{
				Message: "Ok",
				Data: userResp{
					UserID: "69a9a77b-5d8d-5738-80eb-ff0b1fb3846a",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.CreateUser(tt.args.CreateUserData)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUser(t *testing.T) {
	respJSON := `
	{
		"message": "Ok",
		"data": {
		  "createdAt": "2022-05-26T14:47:07.089Z",
		  "updatedAt": "2022-05-26T14:47:09.744Z",
		  "firstName": "Chijioke",
		  "lastName": "Amanambu",
		  "uid": "b@gmail.com",
		  "kycCountry": "NGA",
		  "address": "3 Misratah Street, Wuse 2",
		  "city": "FCT",
		  "postalCode": "900888",
		  "physicalCardCount": 0,
		  "virtualCardCount": 0,
		  "selfieUploaded": true,
		  "idUploaded": true,
		  "ofacChecked": true,
		  "ofacFail": false,
		  "active": true
		}
	  }
	`
	var resp getUserResp
	_ = json.Unmarshal([]byte(respJSON), &resp)
	type args struct {
		userID string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           getUserResp
		wantErr        bool
	}{
		{
			name: "Get a card user",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/user" {
						t.Errorf("Expected to request '/card/v1/user', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(respJSON)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				userID: "d01a03bd-4c83-5b08-b458-1b4a2be535bf",
			},
			want:    resp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateUserAddress(t *testing.T) {
	respJSON := `
	{
		"message": "Ok",
		"data": {
		  "message": "Ok"
		}
	  }
	`
	var resp updateUserAddressResp
	_ = json.Unmarshal([]byte(respJSON), &resp)
	type args struct {
		UpdateUserAddressData
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           updateUserAddressResp
		wantErr        bool
	}{
		{
			name: "Allows integrators to update the address, city, postalCode and kycCountry of their user",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/user/address" {
						t.Errorf("Expected to request '/card/v1/user/address', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(respJSON)))
					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				UpdateUserAddressData: UpdateUserAddressData{
					UserID:     "d01a03bd-4c83-5b08-b458-1b4a2be535bf",
					KycCountry: "NGA",
					Address:    "No 56 bentell gardens estate, lokogoma",
					City:       "KJHGVB",
					PostalCode: "90988VB8",
				},
			},
			want:    resp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.UpdateUserAddress(tt.args.UpdateUserAddressData)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUserAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetCardUserDocURL(t *testing.T) {
	respJSON := `{
		"message": "Ok",
		"data": {
		  "selfieUploadUrl": "https://u54-ci-api-sandbox-kyc-selfie-store.s3.eu-central-1.amazonaws.com/9a2d3488-363b-4fa3-be36-e483a09a8fbf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=ASIA4DZ3BAEQU42QCVPW%2F20220526%2Feu-central-1%2Fs3%2Faws4_request&X-Amz-Date=20220526T165215Z&X-Amz-Expires=3600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEMH%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaDGV1LWNlbnRyYWwtMSJHMEUCIQDMiT0YAZmMF7m9juhqv4AOn%2BRRELhjJwMcEKNlmYNZQAIgHEXnznXPiwQJiCM%2BlG95kmFdS2GLBwPoayE%2FTZluntYqvwIIqv%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FARABGgw4MzI4MTA1ODIzMDUiDB%2Bnw2zUoCSbvsT%2FySqTAlkW2qmk9R9uWIykn6BZjiYWotTlwBEdqzW%2FP%2F624gng2L3VnJMxX3qjaSmEV%2FYcUul0FLoQjqa%2B6%2BL%2FhjFXcjnwmBDqd%2F164R25bdWlIU%2Bzq6i6kogkPQ5%2BAo%2Bmt6rJ%2F8BJ7M%2FN%2BfW0I69xcqNHpagvHbwnvHXUUJQ5NuKA2Q5ss6eLL09hNTM6VRkHZTNxrUExtjzsp9O%2F6YAHpg%2BxXAiyn5D85NP8HIyHIEksoGNyqFAp%2BWBEcmA9Oevyjrax90%2FFfwQm8qLjM8nFE744%2F%2BA%2FLCf0E0TBsBO%2FKzqvRccZh%2B67VL31xnmUGusqkAe%2FsVfmvfmmP0Q8jD5fo3TXejY2aSiWxYKjRzzZ0XGS%2FHVsfuN3MLTgvpQGOpoBv7L8tT4kmofM9BwGtv%2BuTKoCrZUXwBn%2F0PSgz8bcXEbKVTCKu0U9ncno5F9C55ge2U6ukHQXvtdxhhkX%2BGxcG1llPvaYyo4C6qCj535LxhwfwxiifcvSu0ilihAmoY46DZhIs6RGWyt%2FdPuakzbA3PgEb%2BwGo8mR3GN77kACpJxQUjef7CrQeyjOVwBkQO5qobzyRbzTQw8dJQ%3D%3D&X-Amz-Signature=9c8b28d1d32189b4406ddeedd39f948e6dcf50478a74fa66998d53bd7ece2e65&X-Amz-SignedHeaders=host",
		  "idUploadUrl": "https://u54-ci-api-sandbox-kyc-id-store.s3.eu-central-1.amazonaws.com/9a2d3488-363b-4fa3-be36-e483a09a8fbf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=ASIA4DZ3BAEQU42QCVPW%2F20220526%2Feu-central-1%2Fs3%2Faws4_request&X-Amz-Date=20220526T165215Z&X-Amz-Expires=3600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEMH%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaDGV1LWNlbnRyYWwtMSJHMEUCIQDMiT0YAZmMF7m9juhqv4AOn%2BRRELhjJwMcEKNlmYNZQAIgHEXnznXPiwQJiCM%2BlG95kmFdS2GLBwPoayE%2FTZluntYqvwIIqv%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FARABGgw4MzI4MTA1ODIzMDUiDB%2Bnw2zUoCSbvsT%2FySqTAlkW2qmk9R9uWIykn6BZjiYWotTlwBEdqzW%2FP%2F624gng2L3VnJMxX3qjaSmEV%2FYcUul0FLoQjqa%2B6%2BL%2FhjFXcjnwmBDqd%2F164R25bdWlIU%2Bzq6i6kogkPQ5%2BAo%2Bmt6rJ%2F8BJ7M%2FN%2BfW0I69xcqNHpagvHbwnvHXUUJQ5NuKA2Q5ss6eLL09hNTM6VRkHZTNxrUExtjzsp9O%2F6YAHpg%2BxXAiyn5D85NP8HIyHIEksoGNyqFAp%2BWBEcmA9Oevyjrax90%2FFfwQm8qLjM8nFE744%2F%2BA%2FLCf0E0TBsBO%2FKzqvRccZh%2B67VL31xnmUGusqkAe%2FsVfmvfmmP0Q8jD5fo3TXejY2aSiWxYKjRzzZ0XGS%2FHVsfuN3MLTgvpQGOpoBv7L8tT4kmofM9BwGtv%2BuTKoCrZUXwBn%2F0PSgz8bcXEbKVTCKu0U9ncno5F9C55ge2U6ukHQXvtdxhhkX%2BGxcG1llPvaYyo4C6qCj535LxhwfwxiifcvSu0ilihAmoY46DZhIs6RGWyt%2FdPuakzbA3PgEb%2BwGo8mR3GN77kACpJxQUjef7CrQeyjOVwBkQO5qobzyRbzTQw8dJQ%3D%3D&X-Amz-Signature=2cf1ed1568c866071f9377950730fc9a984a14847446685a1f290b1a0f28fd54&X-Amz-SignedHeaders=host",
		  "uid": "b@gmail.com"
		}
	  	}
	  `
	var resp getCardUserDocURLResp
	_ = json.Unmarshal([]byte(respJSON), &resp)

	type args struct {
		userID string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           getCardUserDocURLResp
		wantErr        bool
	}{
		{
			name: "Get a card user document upload url",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/user/documentation/urls" {
						t.Errorf("Expected to request 'card/v1/user/document/url', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(respJSON)))
					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				userID: "d01a03bd-4c83-5b08-b458-1b4a2be535bf",
			},
			want:    resp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetCardUserDocURL(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardUserDocURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCardUserDocURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PostIntegratorDeposit(t *testing.T) {

	type args struct {
		amount   int
		currency string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           PostDepositResp
		wantErr        bool
	}{
		{
			name: "allows an integrator to update their webhook url",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/integrator/v1/deposit" {
						t.Errorf("Expected to request '/integrator/v1/deposit', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`
						{
							"message": "Ok",
							"data": {
							  "depositId": "265bee19-f533-4f6c-8076-4189950efeb2",
							  "amount": 1000,
							  "currency": "USD",
							  "createdAt": "2022-06-01T13:04:46.362Z",
							  "usd": {
								"accountNumber": "Nan",
								"accountName": "Nan",
								"bankName": "Nan",
								"bankAddress": "Nan",
								"branchCode": "Nan",
								"swiftCode": "Nan"
							  }
							}
						}
						`,
					)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				amount:   0,
				currency: "",
			},
			want: PostDepositResp{
				Message: "Ok",
				Data: D5{
					U54DepositId: "",
					DepositId:    "",
					Amount:       0,
					Currency:     "",
					CreatedAt:    "",
					Usd: Usd{
						AccountNumber: "",
						AccountName:   "",
						BankName:      "",
						BankAddress:   "",
						BranchCode:    "",
						SwiftCode:     "",
					},
					Btc: Coin{
						"",
					},
					Eth: Coin{
						"",
					},
					Busd: Coin{
						"",
					},
					Usdc: Coin{
						"",
					},
					Usdt: Coin{
						"",
					},
				},
			},
			wantErr: false,
		},
		// TODO add test cases for different currencies
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.PostIntegratorDeposit(tt.args.amount, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostIntegratorDeposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostIntegratorDeposit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTransaction(t *testing.T) {

	type args struct {
		cardId string
		p      Params
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           TransactionsResp
		wantErr        bool
	}{
		{
			name: "Allows integrators to get a list of all transactions for a given card",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/transactions" {
						t.Errorf("Expected to request '/card/v1/transactions', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{
							"message": "Ok",
							"data": [
							  {
								"transactionId": "55ac5531-ca87-4ed4-bce0-e70b44a44b02",
								"createdAt": "2022-06-02T14:02:37.313Z",
								"debitId": "aa174033-fe13-4c3a-90b3-f3485a0e9c86",
								"debitCurrency": "USD",
								"conversionRate": 1,
								"type": "debit",
								"amount": 250
							  },
							  {
								"transactionId": "e891d291-f76b-4e51-affa-8bd9c3e1d1b5",
								"createdAt": "2022-06-02T11:18:18.788Z",
								"conversionRate": 1,
								"creditCurrency": "USD",
								"type": "credit",
								"amount": 1000
							  }
							]
						  }`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				cardId: "aa174033-fe13-4c3a-90b3-f3485a0e9c86",
				p: Params{
					StartDate: "",
					EndDate:   "",
					Limit:     20,
					Lek:       "",
				},
			},
			want: TransactionsResp{
				Message: "Ok",
				Data: []D4{
					{
						TransactionId:  "55ac5531-ca87-4ed4-bce0-e70b44a44b02",
						CreatedAt:      "2022-06-02T14:02:37.313Z",
						DebitId:        "aa174033-fe13-4c3a-90b3-f3485a0e9c86",
						DebitCurrency:  "USD",
						ConversionRate: 1,
						Type:           "debit",
						Amount:         250,
					},
					{
						TransactionId:  "e891d291-f76b-4e51-affa-8bd9c3e1d1b5",
						CreatedAt:      "2022-06-02T11:18:18.788Z",
						ConversionRate: 1,
						CreditCurrency: "USD",
						Type:           "credit",
						Amount:         1000,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.GetTransaction(tt.args.cardId, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
} // TestClient_GetTransaction

// TestClient_CreateCard
// TestClient_GetCard
// TestClient_GetCards
// TestClient_TopUp
// TestClient_Debit
// TestClient_Freeze
// TestClient_Unfreeze
// TestClient_StopCard
// TestClient_GetFailedTransactions
// TestClient_CreateUser
// TestClient_GetIntegratorDeposit
// TestClient_GetIntegratorFloat
// TestClient_GetIntegratorFloats
// TestClient_GetTransaction
// TestClient_UpdateUserAddress
// TestClient_PostIntegratorDeposit
// TestClient_GetCardUserDocURL
