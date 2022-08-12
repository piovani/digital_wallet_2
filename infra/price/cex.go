package price

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	CEX_URL = "https://cex.io/api/last_price"
)

type Cex struct{}

func NewCex() *Cex {
	return &Cex{}
}

type CexBody struct {
	Lprice string `json:"lprice"`
}

func (c *Cex) getValue(param1 string, param2 string) float64 {
	var body CexBody
	response, err := http.Get(fmt.Sprintf("%s/%s/%s", CEX_URL, param1, param2))
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

	price, err := strconv.ParseFloat(body.Lprice, 64)
	if err != nil {
		fmt.Println(body.Lprice)
		return 1
	}

	return price
}
