package liquidity

import (
	"bytes"
	"encoding/json"
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
	respJSON := `{"message":"Ok","data":{"createdAt":"2022-05-26T14:47:07.089Z","updatedAt":"2022-05-27T10:03:43.208Z","firstName":"Chijioke","lastName":"Amanambu","uid":"b@gmail.com","kycCountry":"NGA","address":"No 56 bentell gardens estate, lokogoma","city":"KJHGVB","postalCode":"90988VB8","physicalCardCount":0,"virtualCardCount":4,"selfieUploaded":true,"idUploaded":true,"ofacChecked":true,"ofacFail":false,"active":true}}
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

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						respJSON)))
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

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						respJSON)))
					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			},
			args: args{
				UpdateUserAddressData: UpdateUserAddressData{
					UserID:     "d01a03bd-4c83-5b08-b458-1b4a2be535bf",
					KycCountry: "USA",
					Address:    "No 58 bentell gardens estate, lokogomb",
					City:       "KJHGVBb",
					PostalCode: "90988VAA",
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

					responseBody := ioutil.NopCloser(bytes.NewReader([]byte(
						respJSON)))
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
