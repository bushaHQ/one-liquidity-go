package liquidity

type IntegratorResp struct {
	Message string `json:"message"`
	Data    D1     `json:"data"`
}

type D1 struct {
	IntegratorId string `json:"integratorId"`
}

type WebhookResp struct {
	Message string `json:"message"`
}
