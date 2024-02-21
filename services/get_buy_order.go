package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// MarketDataResponse represents the response structure from the exchange API.
type MarketDataResponse struct {
	// Define the structure of the response based on the data you expect to receive
	// For example, you might have fields for current price, order book data, etc.
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	// Add more fields as needed
}

// GetMarketData fetches market data for the specified symbol from the exchange API.
func GetCoinbaseata(symbol string, amount string) {
	// Make HTTP request to the exchange API
	url := "https://api.exchange.coinbase.com/products/" + symbol + "/book?level=3"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	err = err
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	// // Return successful response
	return c.JSON(string(body))

}
