package services

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Pair struct to keep track of price, amount, and exchange for each log in the product book
type Pair struct {
	Price    float64
	Amount   string
	Exchange string
}

// Response struct to format the service's returned response to the controller
type Response struct {
	BTCAmount float64  `json:"btcAmount"`
	USDAmount float64  `json:"usdAmount"`
	Exchange  []string `json:"exchange"`
}

// MinHeap data structure to efficiently return the min price to buy at
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Amount < h[j].Amount }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var minHeap MinHeap

// parse through coinbase data for a given symbol - store it in the given min heap
func GetCoinbaseData(symbol string) {
	resp, err := http.Get("https://api.exchange.coinbase.com/products/" + symbol + "/book?level=2")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error:", err)
		return
	}

	asks, ok := data["asks"].([]interface{})
	if !ok {
		fmt.Println("Error: 'asks' field not found or not an array")
		return
	}

	for _, ask := range asks {
		askInfo := ask.([]interface{})
		price := askInfo[0].(string)
		pair_flt_price, err := strconv.ParseFloat(price, 64)
		err = err
		amount := askInfo[1].(string)
		heap.Push(&minHeap, Pair{Price: pair_flt_price, Amount: amount, Exchange: "Coinbase"})
	}
}

// parse through kraken data for a given symbol - store it in the given min heap to merge with coinbase data
func GetKrakenData(symbol string) {
	resp, err := http.Get("https://api.kraken.com/0/public/Depth?pair=" + symbol)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return
	}

	result, ok := data["result"].(map[string]interface{})
	if !ok {
		return
	}

	XXBTZUSD, ok := result["XXBTZUSD"].(map[string]interface{})
	if !ok {
		return
	}

	asks, ok := XXBTZUSD["asks"].([]interface{})
	if !ok {
		return
	}

	for _, ask := range asks {
		askInfo := ask.([]interface{})
		price := askInfo[0].(string)
		pair_flt_price, err := strconv.ParseFloat(price, 64)
		err = err
		amount := askInfo[1].(string)
		heap.Push(&minHeap, Pair{Price: pair_flt_price, Amount: amount, Exchange: "Kraken"})
	}
}

// calculate the weighted average of the buy request for the given passed in amount
func GetAverage(amount string, symbol string) Response {
	heap.Init(&minHeap)
	GetCoinbaseData(symbol)
	GetKrakenData(symbol)

	var totalAmount, totalPrice float64
	float_amt, err := strconv.ParseFloat(amount, 64)
	err = err

	kraken := false
	coinbase := false

	// Pop pairs until the total amount reaches or exceeds the requested amount
	for totalAmount < float_amt && len(minHeap) > 0 {
		pair := heap.Pop(&minHeap).(Pair)
		if pair.Exchange == "Kraken" {
			kraken = true
		} else {
			coinbase = true
		}
		pair_flt_amt, err := strconv.ParseFloat(pair.Amount, 64)
		err = err

		totalAmount += pair_flt_amt
		totalPrice += pair_flt_amt * pair.Price
	}

	weightedAvg := totalPrice / totalAmount

	exchange := []string{}
	if coinbase == true {
		exchange = append(exchange, "coinbase")
	}
	if kraken == true {
		exchange = append(exchange, "kraken")
	}

	resp := Response{
		BTCAmount: float_amt,
		USDAmount: weightedAvg,
		Exchange:  exchange,
	}
	return resp

}
