package liquidity

import (
	"bytes"
	"io/ioutil"
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
				webhook: "https://webhook.site/d8e81cdd-0db9-4b10-82a0-54f8d6be247g",
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

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{
							"message": "Ok",
							"data": {
							  "userId": "69a9a77b-5d8d-5738-80eb-ff0b1fb3846a",
							  "uid": "169c9ff3-4af1-4115-a7eb-363827022625"
							}
						  }`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				userID: "69a9a77b-5d8d-5738-80eb-ff0b1fb3846a",
			},
			want: getUserResp{
				Message: "Ok",
				Data: gur{
					UserID: "69a9a77b-5d8d-5738-80eb-ff0b1fb3846a",
					UID:    "169c9ff3-4af1-4115-a7eb-363827022625",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cl.SetHTTPClient(&tt.mockHttpClient)
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

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{
							"message": "Ok"
						  }`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				UpdateUserAddressData: UpdateUserAddressData{
					UserID:     "69a9a77b-5d8d-5738-80eb-ff0b1fb3846a",
					KycCountry: "NGA",
					Address:    "No 15 bentell gardens estate, lokogoma",
					City:       "FCT",
					PostalCode: "909888",
				},
			},
			want: updateUserAddressResp{
				Message: "Ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cl.SetHTTPClient(&tt.mockHttpClient)
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
	type args struct {
		userID string
	}
	tests := []struct {
		name           string
		mockHttpClient MockHttpClient
		args           args
		want           updateUserAddressResp
		wantErr        bool
	}{
		{
			name: "Get a card user document upload url",
			mockHttpClient: MockHttpClient{

				DoFunc: func(r *http.Request) (*http.Response, error) {
					if r.URL.Path != "card/v1/user/document/url" {
						t.Errorf("Expected to request 'card/v1/user/document/url', got: %s", r.URL.Path)
					}
					if r.Header.Get("Content-Type") != "application/json" {
						t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
					}

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						`{
							"message": "Ok"
						  }`)))

					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				userID: "69a9a77b-5d8d-5738-80eb-ff0b1fb3846a",
			},
			want: updateUserAddressResp{
				Message: "Ok",
			},
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
