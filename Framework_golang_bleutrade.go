package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var apikey = ""
var apisecret = ""

func getResponse(uri string) []byte {

	resp, err := http.Get(uri)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	//le o html da pagina
	html, err := ioutil.ReadAll(resp.Body)

	return html
}

//-------------------PUBLIC FUNCTIONS-------------------//
//-------------------OK-------------------//
//Get the list of all pairs traded.
func getMarkets() *markets {

	uri := "https://bleutrade.com/api/v2/public/getmarkets"

	html := getResponse(uri)

	var m markets

	jsonStr := json.Unmarshal(html, &m)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if m.Success == "false" {
		fmt.Println("Erro:", m)

	}

	return &m
}

//-------------------OK-------------------//
//Get a list of all coins traded.
func getCurrencies() *currencies {

	uri := "https://bleutrade.com/api/v2/public/getcurrencies"
	html := getResponse(uri)

	var c currencies

	jsonStr := json.Unmarshal(html, &c)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if c.Success == "false" {
		fmt.Println("Erro:", c)

	}

	return &c
}

//-------------------OK-------------------//
//Used to get the current tick values for a market.

//Required parameters:
//market (or markets) (ex.: /public/getticker?market=ETH_BTC or /public/getticker?market=ETH_BTC,HTML5_DOGE,DOGE_LTC)
func getTicker(currency string, market string) *ticker {

	uri := "https://bleutrade.com/api/v2/public/getticker?market="

	values := []string{}
	values = append(values, uri+currency+"_"+market)
	uri = strings.Join(values, "")

	html := getResponse(uri)

	var t ticker
	jsonStr := json.Unmarshal(html, &t)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if t.Success == "false" {
		fmt.Println("Erro:", t)

	}

	return &t

}

//-------------------OK-------------------//
//Used to get the last 24 hour summary of all active markets.
func getMarketSummaries() *summaries {

	uri := "https://bleutrade.com/api/v2/public/getmarketsummaries"

	html := getResponse(uri)

	var s summaries
	jsonStr := json.Unmarshal(html, &s)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if s.Success == "false" {
		fmt.Println("Erro:", s)

	}

	return &s
}

//-------------------OK-------------------//
//Used to get the last 24 hour summary of specific market.
func getMarketSummary(currency string, market string) *summary {

	uri := "https://bleutrade.com/api/v2/public/getmarketsummary?market="

	values := []string{}
	values = append(values, uri+currency+"_"+market)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var s summary

	jsonStr := json.Unmarshal(html, &s)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if s.Success == "false" {
		fmt.Println("Erro:", s)

	}

	return &s
}

//-------------------OK-------------------//
//Loads the book offers specific market.

//Required parameters:
//market
//type (BUY, SELL, ALL)
//depth (optional, default is 20)
func getOrderBook(currency string, market string, ordertype string, depth string) *orderbook {

	uri := "https://bleutrade.com/api/v2/public/getorderbook?market="
	values := []string{}
	values = append(values, uri+currency+"_"+market+"&type="+ordertype)

	if depth != "" {
		values = append(values, "&depth=")
		values = append(values, depth)
	}

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var o orderbook
	jsonStr := json.Unmarshal(html, &o)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if o.Success == "false" {
		fmt.Println("Erro:", o)

	}
	return &o
}

//-------------------OK-------------------//
//Obtains historical trades of a specific market.

//Required parameters:
//market
//count (optional, default: 20, max: 200)
func getMarketHistory(currency string, market string, count string) *markethistory {

	uri := "https://bleutrade.com/api/v2/public/getmarkethistory?market="

	values := []string{}
	values = append(values, uri+currency+"_"+market)

	if count != "" {
		values = append(values, "&count=20")
	} else {
		values = append(values, "&count="+count)
	}

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var m markethistory
	jsonStr := json.Unmarshal(html, &m)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if m.Success == "false" {
		fmt.Println("Erro:", m)
	}
	return &m
}

//-------------------OK-------------------//
//Obtains candles format historical trades of a specific market.

// Required parameters:
// market
// period (15m, 20m, 30m, 1h, 2h, 3h, 4h, 6h, 8h, 12h, 1d)
// count (default: 1000, max: 999999)
// lasthours (default: 24, max: 2160)
func getCandles(currency string, market string, period string, count string, lasthours string) *candles {

	uri := "https://bleutrade.com/api/v2/public/getcandles?market="

	values := []string{}
	values = append(values, uri+currency+"_"+market+"&period="+period+"&count="+count+"&lasthours="+lasthours)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var c candles
	jsonStr := json.Unmarshal(html, &c)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if c.Success == "false" {
		fmt.Println("Erro:", c)
	}
	return &c
}

//-------------------PRIVATE FUNCTIONS-------------------//

func hashHmac512(request string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(request))
	return hex.EncodeToString(h.Sum(nil))
}

//-------------------OK-------------------//
//Use to send BUY orders

//Required parameters:
//market
//rate
//quantity
//comments (optional, up to 128 characters)
func buyLimit(currency string, market string, rate string, quantity string, comments string) *marketsellbuy {

	uri := "https://bleutrade.com/api/v2/market/buylimit?"

	values := []string{}
	values = append(values, uri+"market="+currency+"_"+market+"&rate="+rate+"&quantity="+quantity)

	query := ""
	if comments != "" {
		parameters := url.Values{}
		parameters.Add("comments", comments)
		query = parameters.Encode()
	}

	values = append(values, query+"&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))

	sign := strings.Join(values, "")
	sign = hashHmac512(sign, apisecret)

	values = append(values, "&apisign="+sign)
	uri = strings.Join(values, "")
	html := getResponse(uri)

	var limit marketsellbuy
	jsonStr := json.Unmarshal(html, &limit)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if limit.Success == "false" {
		fmt.Println("Erro:", limit)
	}
	return &limit

}

//-------------------OK-------------------//
//Use to send SELL orders

//Required parameters:
//market
//rate 0.00000037

//quantity
//comments (optional, up to 128 characters)
func sellLimit(currency string, market string, rate string, quantity string, comments string) *marketsellbuy {

	uri := "https://bleutrade.com/api/v2/market/selllimit?"

	values := []string{}
	values = append(values, uri+"market="+currency+"_"+market+"&rate="+rate+"&quantity="+quantity+"&")
	query := ""
	if comments != "" {
		parameters := url.Values{}
		parameters.Add("comments", comments)
		query = parameters.Encode()
	}

	values = append(values, query+"&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))

	sign := strings.Join(values, "")
	sign = hashHmac512(sign, apisecret)

	values = append(values, "&apisign="+sign)
	uri = strings.Join(values, "")

	html := getResponse(uri)

	var limit marketsellbuy
	jsonStr := json.Unmarshal(html, &limit)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if limit.Success == "false" {
		fmt.Println("Erro:", limit)
	}
	return &limit

}

//-------------------OK-------------------//
//Use to cancel an order

//Required parameters:
//orderid
func cancelOrder(orderid string) *commonreturn {

	uri := "https://bleutrade.com/api/v2/market/cancel?orderid="

	values := []string{}
	values = append(values, uri+orderid)

	values = append(values, "&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))
	sign := strings.Join(values, "")

	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var c commonreturn
	jsonStr := json.Unmarshal(html, &c)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if c.Success == "false" {
		fmt.Println("Erro:", c)
	}

	return &c
}

//-------------------OK-------------------//
//Use to list your open orders
func getOpenOrders() *openorders {

	uri := "https://bleutrade.com/api/v2/market/getopenorders?apikey="

	values := []string{}
	values = append(values, uri+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))

	sign := strings.Join(values, "")

	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var o openorders

	jsonStr := json.Unmarshal(html, &o)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if o.Success == "false" {
		fmt.Println("Erro:", o)
	}

	return &o
}

//-------------------OK-------------------//
//Use to get the balance of all your coins

//Required parameters:
//currencies (optional, default=ALL) eg.: /account/getbalances?currencies=DOGE;BTC
func getBalances(currencie string) *balances {

	uri := "https://bleutrade.com/api/v2/account/getbalances?"

	values := []string{}
	values = append(values, uri)

	if currencie != "" {
		values = append(values, "currencies="+currencie)
	}

	values = append(values, "&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))
	sign := strings.Join(values, "")
	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var b balances

	jsonStr := json.Unmarshal(html, &b)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if b.Success == "false" {
		fmt.Println("Erro:", b)
	}

	return &b

}

//-------------------OK-------------------//
//Use to get the balance of a specific currency

//Required parameters:
//currency
func getBalance(currency string) *balance {

	uri := "https://bleutrade.com/api/v2/account/getbalance?currency="

	values := []string{}
	values = append(values, uri+currency)

	values = append(values, "&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))
	sign := strings.Join(values, "")
	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var b balance

	jsonStr := json.Unmarshal(html, &b)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if b.Success == "false" {
		fmt.Println("Erro:", b)
	}

	return &b
}

//-------------------OK-------------------//
//Use to get the deposit address of specific coin.

//Required parameters:
//currency
func getDepositAdress(currency string) *depositaddress {

	uri := "https://bleutrade.com/api/v2/account/getdepositaddress?currency="

	values := []string{}
	values = append(values, uri+currency)

	values = append(values, "&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))
	sign := strings.Join(values, "")
	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var d depositaddress

	jsonStr := json.Unmarshal(html, &d)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if d.Success == "false" {
		fmt.Println("Erro:", d)
	}

	return &d
}

//-------------------OK-------------------//
//Use to withdraw their currencies to another wallet.

//Required parameters:
//currency
//quantity
//address
//comments (optional, up to 128 characters)
func withdraw(currency string, quantity string, address string, comments string) *commonreturn {

	uri := "https://bleutrade.com/api/v2/account/withdraw?currency="

	values := []string{}
	values = append(values, uri+currency+"&quantity="+quantity+"&address="+address)

	if comments != "" {
		values = append(values, "&comments="+comments)
	}

	values = append(values, "&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))
	sign := strings.Join(values, "")
	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var c commonreturn

	jsonStr := json.Unmarshal(html, &c)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if c.Success == "false" {
		fmt.Println("Erro:", c)
	}

	return &c
}

//-------------------OK-------------------//
//Use to get the data given order

//Required parameters:
//orderid
func getOrder(orderid string) (*ordertrue, *ordererror) {

	uri := "https://bleutrade.com/api/v2/account/getorder?orderid="

	values := []string{}
	values = append(values, uri+orderid)

	values = append(values, "&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))
	sign := strings.Join(values, "")
	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var otrue ordertrue
	var oerror ordererror

	erro := json.Unmarshal((html), &otrue)
	//fmt.Println(string(html))

	if otrue.Success == "false" {
		erro = json.Unmarshal((html), &oerror)
	}
	if erro != nil {
		fmt.Println("Erro:", erro)
	}

	return &otrue, &oerror
}

//Use to list your orders

//Required parameters:
//market (DIVIDEND_DIVISOR or ALL)
//orderstatus (ALL, OK, OPEN, CANCELED)
//ordertype (ALL, BUY, SELL)
//depth (optional, default is 500, max is 20000)

//example use:
// getOrders("ALL", "", "ALL", "BUY", "")
// getOrders("", "ALL", "ALL", "SELL", "600")
func getOrders(currency string, market string, orderstatus string, ordertype string, depth string) *orders {

	uri := "https://bleutrade.com/api/v2/account/getorders?market="

	values := []string{}
	if currency != "" && market != "" {
		values = append(values, uri+currency+"_"+market+"&orderstatus="+orderstatus+"&ordertype="+ordertype)
	} else {
		values = append(values, uri+"ALL"+"&orderstatus="+orderstatus+"&ordertype="+ordertype)
	}

	if depth != "" {
		values = append(values, "&depth="+depth)
	}

	values = append(values, "&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))
	sign := strings.Join(values, "")
	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var o orders

	jsonStr := json.Unmarshal(html, &o)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if o.Success == "false" {
		fmt.Println("Erro:", o)
	}

	return &o
}

//-------------------OK-------------------//
//Use for historical trades of a given order.

//Required parameters:
//orderid
func getOrderHistory(orderid string) *orderhistory {

	uri := "https://bleutrade.com/api/v2/account/getorderhistory?orderid="

	values := []string{}
	values = append(values, uri+orderid)

	values = append(values, "&apikey="+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))
	sign := strings.Join(values, "")
	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var o orderhistory

	jsonStr := json.Unmarshal(html, &o)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if o.Success == "false" {
		fmt.Println("Erro:", o)
	}

	return &o
}

//-------------------OK-------------------//
//Use for historical deposits and received direct transfers.
func getDepositHistory() *deposithistory {

	uri := "https://bleutrade.com/api/v2/account/getdeposithistory?apikey="

	values := []string{}
	values = append(values, uri+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))

	sign := strings.Join(values, "")
	//gera o hash e inclui na url
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var d deposithistory

	jsonStr := json.Unmarshal(html, &d)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if d.Success == "false" {
		fmt.Println("Erro:", d)
	}

	return &d
}

//-------------------OK-------------------//
//Use for historical withdraw and sent direct transfers.
func getWithdrawHistory() *withdrawhistory {

	uri := "https://bleutrade.com/api/v2/account/getwithdrawhistory?apikey="

	values := []string{}
	values = append(values, uri+apikey+"&nonce="+strconv.Itoa(int(time.Now().Unix())))

	sign := strings.Join(values, "")
	sign = hashHmac512(sign, apisecret)
	values = append(values, "&apisign="+sign)

	uri = strings.Join(values, "")
	html := getResponse(uri)

	var w withdrawhistory

	jsonStr := json.Unmarshal(html, &w)
	if jsonStr != nil {
		fmt.Println("Erro:", jsonStr)
	}
	if w.Success == "false" {
		fmt.Println("Erro:", w)
	}

	return &w
}
