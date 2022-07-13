# one-liquidity-go

## Introduction
This is a Go wrapper around the [API](https://docs.oneliquidity.technology/docs) for [Cards Integrators by OneLiquidity](https://docs.oneliquidity.technology/).

## Installation
To install, run

``` go get github.com/bushaHQ/one-liquidity-go```

## Import Package
The base class for this package is 'one-liquidity-go'. To use this class, add:

```
import (
  "github.com/bushaHQ/one-liquidity-go"
)
 ```

## Initialization

To use One Liquidity, instantiate one-liquidity-go with your public key. We recommend that you store your secret key in an environment variable named, ```LIQUIDITY_PRIVATE_KEY```. See example below.
 ```
err := godotenv.Load("./.env")
if err != nil {
  log.Fatal("Error loading .env file")
}
client := liquidity.NewClient()
 ```
You can override the default settings by passing in the following parameters:

```
client.SetBaseURL("your-base-url")
client.SetDebug(false)
client.SetHTTPClient(&http.Client{
  Timeout: your-timeout,
})
client.SetAuth(os.Getenv({LIQUIDITY_PRIVATE_KEY}))
```

# Card Integration Methods
This is the documentation for all of the components of card Integrator

**Methods Included:**

* ```.RegisterIntegrator```

* ```.UpdateWebhook```

* ```.CreateCard```

* ```.GetCard```

* ```.GetCards```

* ```.TopUp```

* ```.Debit```

* ```.GetCardDeposit```

* ```.PostCardDeposit```

* ```.Freeze```

* ```.Unfreeze```

* ```.StopCard```

* ```.GetFailedTransaction```

* ```.GetFailedTransactions```

* ```.GetTransaction```

* ```.GetIntegratorDeposit```

* ```.PostIntegratorDeposit```

* ```.GetIntegratorFloats```

* ```.GetIntegratorFloat```

* ```.UpdateFloatDefault```

* ```.GetUser```

* ```.CreateUser```

* ```.UpdateUserAddress```

* ```.GetCardUserDocURL```

### ```.RegisterIntegrator(data RegisterAccountData) (AccountResp, error)```
This is called to allow an integrator register with the system. The payload should be of type ```liquidity.RegisterIntegratorData```. See below for  ```liquidity.RegisterIntegratorData``` definition

```
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
```
A sample register call is:

```
payload := liquidity.RegisterIntegratorData{
  FloatCurrencies:    []string{"USD"},
  FirstName:          "Olusola",
  LastName:           "Alao",
  Country:            "NG",
  BusinessName:       "Algoro",
  RegistrationNumber: "RC-5467898",
  BusinessAddress:    "Ajah",
  Domain:             "https://ajalekoko.com",
  Email:              "ajalenkoko@gmail.com",
  WebhookUrl:         "https://webhook.site/185f138247c1",
  ContactNumber:      "+2349034384669",
}

response, err := client.RegisterAccount(payload)
if err != nil {
  panic(err)
}
    
fmt.Println(response)
```
#### Sample Response

```
    {Ok {863494e2-40b3-44dc-8acf-5ee520097f75}}
```

### ```.UpdateWebhook(webhook, string) (Resp, error)```
This is called to allow an integrator to update their webhook URL.

A sample valid call is:

```
response, err := client.UpdateWebhook("https://webhook.site/4491ce3b0b0e")
if err != nil {
  panic(err)
}
fmt.Println(response)
```
#### Sample Response

```
{Ok}
```

### ```.CreateCard(data CreateCardData) (CardResp, error)```
This is called to allow an integrator to create a virtual card for their user. The payload should be of type ```spend-juice-go.RegisterUserData```. See below for  ```spend-juice-go.RegisterUserData``` definition
```
type CreateCardData struct {
  UserId    string    `json:"userId"`
  Expiry    time.Time `json:"expiry"`
  SingleUse bool      `json:"singleUse"`
}
```
UserId is the unique OneLiquidity user id that will own the card
Expiry is a date string in the format YYYY-MM-DD representing the expiry time for the ordered card. Date must be in the future. Note, card will expire at midnight of following day.
SingleUse defines whether the card is single use or not - currently only accepts false


A sample CreateCard call is:

```
payload := juice.CreateCardData{
  UserId: "d01a03bd-4c83-5b08-b458-1b4a2be535bf",
  Expiry: time.Now(),
  SingleUse:  false,
}
        
response, err := client.RegisterUser(payload)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
#### Sample Response

```
{Ok { aa174033-fe13-4c3a-90b3-f3485a0e9c86 d01a03bd-4c83-5b08-b458-1b4a2be535bf 2025-04-19T00:00:00.000Z 04/25 204 2178 0 issued USD false 2022-05-27T09:57:16.597Z Chijioke Amanambu 5368988938002178  214103800064766}}
```

### ```.GetCard(card, trackingNumber string) (CardResp, error)```
This is called to allows an integrator to get full details of one card for their user.
card is the id identifying the card to return details for
trackingNumber is the tracking number of the card


A sample list users call is:

```
response, err := client.GetCard("aa174033-fe13-4c3a-90b3-f3485a0e9c86", "ande3ge3")
if err != nil {
  panic(err)
}
fmt.Println(response)
```
#### Sample Response

```
 {Ok { aa174033-fe13-4c3a-90b3-f3485a0e9c86 d01a03bd-4c83-5b08-b458-1b4a2be535bf 2025-04-19T00:00:00.000Z 04/25 204 2178 0 issued USD false 2022-05-27T09:57:16.597Z Chijioke Amanambu 5368988938002178  214103800064766}}
```





### ```.GetCards(card Params) (CardsResp, error)```
This is called to allow an integrator to get all cards for their user

```
type Params struct {
  Id        string
  Type      string
  StartDate time.Time
  EndDate   time.Time
  Limit     int
  Lek       string
}
```
Id is the id identifying a card user
Type is the card type to return in the response
StartDate is the start date of chargebacks to include in the response
EndDate is the end date of chargebacks to include in the response
Limit is number of cards to fetch per page
Lek is the value provided in the last call for paginated responses

A sample list users call is:

```
response, err := client.GetCard("aa174033-fe13-4c3a-90b3-f3485a0e9c86", "ande3ge3")
if err != nil {
  panic(err)
}
fmt.Println(response)
```
#### Sample Response

```
 {Ok { aa174033-fe13-4c3a-90b3-f3485a0e9c86 d01a03bd-4c83-5b08-b458-1b4a2be535bf 2025-04-19T00:00:00.000Z 04/25 204 2178 0 issued USD false 2022-05-27T09:57:16.597Z Chijioke Amanambu 5368988938002178  214103800064766}}
```