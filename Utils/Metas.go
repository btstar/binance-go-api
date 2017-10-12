package utils

import "net/http"

type Order struct {
	Price,
	Amount,
	AvgPrice,
	DealAmount,
	Fee float64
	OrderID   int
	OrderTime int
	Status    TradeStatus
	Currency  CurrencyPair
	Side      TradeSide
}

type Trade struct {
	Tid    int64   `json:"tid"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount,string"`
	Price  float64 `json:"price,string"`
	Date   int64   `json:"date_ms"`
}

type SubAccount struct {
	Currency Currency
	Amount,  //available
	ForzenAmount float64 //forzen
}

type Account struct {
	Exchange    string
	Asset       float64 //total asset
	NetAsset    float64 //net asset
	SubAccounts map[Currency]SubAccount
}

type Ticker struct {
	Last float64 `json:"last"`
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
	High float64 `json:"high"`
	Low  float64 `json:"low"`
	Vol  float64 `json:"vol"`
	Date uint64  `json:"date"`
}

type Tickers struct {
	//Last float64 `json:"last"`
	Buy  float64 `json:"buy"`
	BuyQty  float64 `json:"buyqty"`
	Sell float64 `json:"sell"`
	SellQty float64 `json:"sellqty"`
	Currency string `json:currency`
	Date uint64  `json:"date"`
}

type DepthRecord struct {
	Price,
	Amount float64
}

type DepthRecords []DepthRecord

func (dr DepthRecords) Len() int {
	return len(dr)
}

func (dr DepthRecords) Swap(i, j int) {
	dr[i], dr[j] = dr[j], dr[i]
}

func (dr DepthRecords) Less(i, j int) bool {
	return dr[i].Price < dr[j].Price
}

type Depth struct {
	AskList,
	BidList DepthRecords
}

type APIConfig struct {
	HttpClient *http.Client
	ApiUrl,
	AccessKey,
	SecretKey string
}

type Kline struct {
	Timestamp int64
	Open,
	Close,
	High,
	Low,
	Vol float64
}
