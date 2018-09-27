package main

type balances struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		Currency      string `json:"Currency"`
		Balance       string `json:"Balance"`
		Available     string `json:"Available"`
		Pending       string `json:"Pending"`
		CryptoAddress string `json:"CryptoAddress"`
		IsActive      string `json:"IsActive"`
		AllowDeposit  string `json:"AllowDeposit"`
		AllowWithdraw string `json:"AllowWithdraw"`
	} `json:"result"`
}

type balance struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  struct {
		Currency      string `json:"Currency"`
		Balance       string `json:"Balance"`
		Available     string `json:"Available"`
		Pending       int    `json:"Pending"`
		CryptoAddress string `json:"CryptoAddress"`
		IsActive      string `json:"IsActive"`
		AllowDeposit  string `json:"AllowDeposit"`
		AllowWithdraw string `json:"AllowWithdraw"`
	} `json:"result"`
}

type orderbook struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  struct {
		Buy []struct {
			Quantity string `json:"Quantity"`
			Rate     string `json:"Rate"`
		} `json:"buy"`
		Sell []struct {
			Quantity string `json:"Quantity"`
			Rate     string `json:"Rate"`
		} `json:"sell"`
	} `json:"result"`
}

type currencies struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		Currency        string  `json:"Currency"`
		CurrencyLong    string  `json:"CurrencyLong"`
		MinConfirmation int     `json:"MinConfirmation"`
		TxFee           float64 `json:"TxFee"`
		CoinType        string  `json:"CoinType"`
		IsActive        string  `json:"IsActive"`
		MaintenanceMode string  `json:"MaintenanceMode"`
	} `json:"result"`
}

type markets struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		MarketName         string  `json:"MarketName"`
		MarketCurrency     string  `json:"MarketCurrency"`
		BaseCurrency       string  `json:"BaseCurrency"`
		MarketCurrencyLong string  `json:"MarketCurrencyLong"`
		BaseCurrencyLong   string  `json:"BaseCurrencyLong"`
		IsActive           string  `json:"IsActive"`
		MinTradeSize       float64 `json:"MinTradeSize"`
	} `json:"result"`
}

type ticker struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		Market string `json:"Market"`
		Bid    string `json:"Bid"`
		Ask    string `json:"Ask"`
		Last   string `json:"Last"`
	} `json:"result"`
}

type summaries struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		TimeStamp      string `json:"TimeStamp"`
		MarketName     string `json:"MarketName"`
		MarketCurrency string `json:"MarketCurrency"`
		BaseCurrency   string `json:"BaseCurrency"`
		PrevDay        string `json:"PrevDay"`
		High           string `json:"High"`
		Low            string `json:"Low"`
		Last           string `json:"Last"`
		Average        string `json:"Average"`
		Volume         string `json:"Volume"`
		BaseVolume     string `json:"BaseVolume"`
		Bid            string `json:"Bid"`
		Ask            string `json:"Ask"`
		IsActive       string `json:"IsActive"`
	} `json:"result"`
}

type summary struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		TimeStamp      string  `json:"TimeStamp"`
		MarketName     string  `json:"MarketName"`
		MarketCurrency string  `json:"MarketCurrency"`
		BaseCurrency   string  `json:"BaseCurrency"`
		PrevDay        float64 `json:"PrevDay"`
		High           float64 `json:"High"`
		Low            float64 `json:"Low"`
		Last           float64 `json:"Last"`
		Average        float64 `json:"Average"`
		Volume         float64 `json:"Volume"`
		BaseVolume     float64 `json:"BaseVolume"`
		Bid            float64 `json:"Bid"`
		Ask            float64 `json:"Ask"`
		IsActive       string  `json:"IsActive"`
	} `json:"result"`
}

type markethistory struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		TradeID    string `json:"TradeID"`
		TimeStamp  string `json:"TimeStamp"`
		Quantity   string `json:"Quantity"`
		Price      string `json:"Price"`
		BaseVolume string `json:"BaseVolume"`
		OrderType  string `json:"OrderType"`
		Total      string `json:"Total"`
	} `json:"result"`
}

type candles struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		TimeStamp  string `json:"TimeStamp"`
		Open       string `json:"Open"`
		High       string `json:"High"`
		Low        string `json:"Low"`
		Close      string `json:"Close"`
		Volume     string `json:"Volume"`
		BaseVolume string `json:"BaseVolume"`
	} `json:"result"`
}

type marketsellbuy struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  struct {
		Orderid string `json:"orderid"`
	} `json:"result"`
}

type openorders struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		OrderID            string `json:"OrderId"`
		Exchange           string `json:"Exchange"`
		Type               string `json:"Type"`
		Quantity           string `json:"Quantity"`
		QuantityRemaining  string `json:"QuantityRemaining"`
		QuantityBaseTraded string `json:"QuantityBaseTraded"`
		Price              string `json:"Price"`
		Status             string `json:"Status"`
		Created            string `json:"Created"`
		Comments           string `json:"Comments"`
	} `json:"result"`
}

type depositaddress struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  struct {
		Currency     string `json:"Currency"`
		CurrencyLong string `json:"CurrencyLong"`
		Address      string `json:"Address"`
	} `json:"result"`
}

type deposithistory struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		ID        string `json:"Id"`
		Coin      string `json:"Coin"`
		Amount    string `json:"Amount"`
		TimeStamp string `json:"TimeStamp"`
		Label     string `json:"Label"`
	} `json:"result"`
}

type withdrawhistory struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		ID            string `json:"Id"`
		TimeStamp     string `json:"TimeStamp"`
		Coin          string `json:"Coin"`
		Amount        string `json:"Amount"`
		Label         string `json:"Label"`
		TransactionID string `json:"TransactionId"`
	} `json:"result"`
}

type commonreturn struct {
	Success string        `json:"success"`
	Message string        `json:"message"`
	Result  []interface{} `json:"result"`
}

type ordertrue struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  struct {
		OrderID            string `json:"OrderId"`
		Exchange           string `json:"Exchange"`
		Type               string `json:"Type"`
		Quantity           string `json:"Quantity"`
		QuantityRemaining  string `json:"QuantityRemaining"`
		Price              string `json:"Price"`
		Status             string `json:"Status"`
		Created            string `json:"Created"`
		Comments           string `json:"Comments"`
		QuantityBaseTraded string `json:"QuantityBaseTraded"`
	} `json:"result"`
}

type ordererror struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		OrderID            string `json:"OrderId"`
		Exchange           string `json:"Exchange"`
		Type               string `json:"Type"`
		Quantity           string `json:"Quantity"`
		QuantityRemaining  string `json:"QuantityRemaining"`
		Price              string `json:"Price"`
		Status             string `json:"Status"`
		Created            string `json:"Created"`
		Comments           string `json:"Comments"`
		QuantityBaseTraded string `json:"QuantityBaseTraded"`
	} `json:"result"`
}

type orders struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		OrderID            string `json:"OrderId"`
		Exchange           string `json:"Exchange"`
		Type               string `json:"Type"`
		Quantity           string `json:"Quantity"`
		QuantityRemaining  string `json:"QuantityRemaining"`
		QuantityBaseTraded string `json:"QuantityBaseTraded"`
		Price              string `json:"Price"`
		Status             string `json:"Status"`
		Created            string `json:"Created"`
		Comments           string `json:"Comments"`
	} `json:"result"`
}

type orderhistory struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		OrderID            string `json:"OrderId"`
		Exchange           string `json:"Exchange"`
		Type               string `json:"Type"`
		Quantity           string `json:"Quantity"`
		QuantityRemaining  string `json:"QuantityRemaining"`
		Price              string `json:"Price"`
		Status             string `json:"Status"`
		Created            string `json:"Created"`
		Comments           string `json:"Comments"`
		QuantityBaseTraded string `json:"QuantityBaseTraded"`
	} `json:"result"`
}
