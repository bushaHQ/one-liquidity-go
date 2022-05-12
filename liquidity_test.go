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
		want           WebhookResp
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
			want: WebhookResp{
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
