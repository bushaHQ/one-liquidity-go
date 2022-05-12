package liquidity

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

type WebhookData struct {
	Webhook string `json:"webhook"`
}

// RegisterIntegrator allows an integrator register with the system
func (cl *Client) RegisterIntegrator(data RegisterIntegratorData) (IntegratorResp, error) {
	var res IntegratorResp
	err := cl.post("/integrator/v1/register", data, &res)
	return res, err
}

type WebhookResp struct {
	Message string `json:"message"`
}

// UpdateWebhook allows an integrator to update their webhook URL
func (cl *Client) UpdateWebhook(webhook string) (WebhookResp, error) {
	var res WebhookResp
	err := cl.patch("/integrator/v1/webhook", WebhookData{webhook}, &res)
	return res, err
}
