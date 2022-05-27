package liquidity

import (
	"bytes"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
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
					if r.URL.Path != "/integrator/v1/card" {
						t.Errorf("Expected to request '/integrator/v1/card', got: %s", r.URL.Path)
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
			//cl.SetHTTPClient(&tt.mockHttpClient)
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
				cardId: "",
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
			got, err := cl.Freeze(tt.args.cardId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Freeze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freeze() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetCardDeposit(t *testing.T) {

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
			name: "can get 1Liquidity Union54 float deposit with the deposit ID",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/service/deposit" {
						t.Errorf("Expected to request '/card/v1/service/deposit', got: %s", r.URL.Path)
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
			got, err := cl.GetCardDeposit(tt.args.depositId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardDeposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCardDeposit() got = %v, want %v", got, tt.want)
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
			name: "allows an integrator to update their webhook url",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "/card/v1/transactions/failed" {
						t.Errorf("Expected to request '/card/v1/transactions/failed', got: %s", r.URL.Path)
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
				p: Params{
					Id:        "",
					Type:      "",
					StartDate: "",
					EndDate:   "",
					Limit:     0,
					Lek:       "",
				},
			},
			want: TransactionsResp{
				Message: "",
				Data:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
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
				cardId: "",
				p: Params{
					Id:        "",
					Type:      "",
					StartDate: "",
					EndDate:   "",
					Limit:     0,
					Lek:       "",
				},
			},
			want: TransactionsResp{
				Message: "",
				Data:    nil,
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
}

func TestClient_PostCardDeposit(t *testing.T) {

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
					if r.URL.Path != "/card/v1/service/deposit" {
						t.Errorf("Expected to request '/card/v1/service/deposit', got: %s", r.URL.Path)
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
				amount:   0,
				currency: "",
			},
			want: PostDepositResp{
				Message: "",
				Data: D5{
					U54DepositId: "",
					DepositId:    "",
					Amount:       0,
					Currency:     "",
					CreatedAt:    "",
					Usd:          Usd{},
					Btc:          Coin{},
					Eth:          Coin{},
					Busd:         Coin{},
					Usdc:         Coin{},
					Usdt:         Coin{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl.SetHTTPClient(&tt.mockHttpClient)
			got, err := cl.PostCardDeposit(tt.args.amount, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostCardDeposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostCardDeposit() got = %v, want %v", got, tt.want)
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
						`{"message":"Ok"}`)))

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
				Message: "",
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
			name: "allows an integrator to update their webhook url",
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
			name: "allows an integrator to update their webhook url",
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
				cardId: "",
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
			got, err := cl.Unfreeze(tt.args.cardId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unfreeze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unfreeze() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateFloatDefault(t *testing.T) {

	type args struct {
		floatId string
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
					if r.URL.Path != "/integrator/v1/float/default" {
						t.Errorf("Expected to request '/integrator/v1/float/default', got: %s", r.URL.Path)
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
				floatId: "",
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
			got, err := cl.UpdateFloatDefault(tt.args.floatId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateFloatDefault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateFloatDefault() got = %v, want %v", got, tt.want)
			}
		})
	}
}
