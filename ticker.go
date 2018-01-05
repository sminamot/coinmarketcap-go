package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://api.coinmarketcap.com/v1/ticker/?convert=JPY&limit=0"

func GetAllCoinData() (Coins, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(string(body))
	}

	c := &[]Coin{}
	err = json.Unmarshal(body, c)
	if err != nil {
		return nil, err
	}

	co := Coins{}
	for _, v := range *c {
		co[v.Symbol] = v
	}

	return co, nil
}
