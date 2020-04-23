package main

import "os"
import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

/**
{
	"trade_id": 25944746,
	"price": "6595.47",
	"size": "0.00829901",
	"time": "2020-04-23T05:04:47.351175Z",
	"bid": "6589.58",
	"ask": "6591.88",
	"volume": "2292.20935984"
}
*/

// Define data structure
type Response struct {
	TradeID int
	Price   string
	Size    string
	Bid     string
	Ask     string
	Volume  string
	Time    string
}

func get_content() {
	// json data
	url := "https://api.gdax.com/products/BTC-EUR/ticker"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data Response

	// unmarshall
	json.Unmarshal(body, &data)

	// fmt.Printf("Results: %v\n", data)
	// {0 6592.99 0.24015291 6587.01 6591.69 2292.08098398 2020-04-23T05:06:09.559705Z}

	// print values of the object
	fmt.Printf("Price: $%s\n", data.Price)
	fmt.Printf("Bid: $%s\n", data.Bid)
	fmt.Printf("Ask: $%s\n", data.Ask)

	os.Exit(0)

	/**
	{
		"trade_id": 25944746,
		"price": "6595.47",
		"size": "0.00829901",
		"time": "2020-04-23T05:04:47.351175Z",
		"bid": "6589.58",
		"ask": "6591.88",
		"volume": "2292.20935984"
	}
	*/
}

func main() {
	get_content()
}
