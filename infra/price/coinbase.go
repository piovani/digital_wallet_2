package price

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	COINBASE_URL = "https://api.coinbase.com/v2/exchange-rates"
)

type Coinbase struct{}

func NewCoinbase() *Coinbase {
	return &Coinbase{}
}

type CoinbaseBody struct {
	Data struct {
		Currency string `json:"currency"`
		Rates    struct {
			Eur string `json:"EUR"`
		} `json:"rates"`
	} `json:"data"`
}

func (c *Coinbase) GetDogeEur() float64 {
	var body CoinbaseBody
	response, err := http.Get(fmt.Sprintf("%s?currency=DOGE", COINBASE_URL))
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	price, err := strconv.ParseFloat(body.Data.Rates.Eur, 64)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	return price
}
