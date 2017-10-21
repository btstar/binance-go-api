# Golang Binance API
golang wrapper for latest [Binance API](https://www.binance.com/restapipub.html)

## Support latest api

| API function | REST API | Websocket API | Description |
|----------|------|-----------|-----|
| GetTicker | Yes  | No        | Getting latest price of a symbol  |
| GetTickers | Yes  | No        | Getting best price/qty on the order book for all symbols  |
| GetDepth | Yes  | No        | Getting depth of a symbol or maintain a depth cache locally  |
| LimitBuy | Yes  | No        | Placing a LIMIT buy order  |
| LimitSell | Yes  | No        | Placing a LIMIT sell order  |
| MarketBuy | Yes  | No       | Placing a MARKET buy order  |
| MarketSell | Yes  | No       | Placing a MARKET sell order  |
| GetOneOrder | Yes  | No     | Checking an order’s status  |
| CancelOrder     | Yes  | No        | Cancelling an order  |
| GetUnfinishOrders | Yes | No       | Getting list of open orders  |
| GetAccount | Yes | No | Getting list of current position |

## Pre-condition

+ [Install golang](https://golang.org/doc/install)
+ Install binance-api library, 
```
go get github.com/BtStar/binance-go-api
```

## Example

```

package main

import (
    "log"
    "net/http"
    "github.com/BtStar/binance-go-api"
    "github.com/BtStar/binance-go-api/Utils"
)

func main() {
	//accessKey,secretKey can apply in binance website, it's necessary.
	accessKey := "your access key"
	secretKey := "your secret key"

	//New Binance API
	BN := binance.New(http.DefaultClient, accessKey, secretKey)

	//Getting list of current position
	account, err := BN.GetAccount()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Total Asset :%.4f, Net Asset:%.4f", account.Asset, account.NetAsset)
	log.Printf("[BTC Asset] Avaliable:%.4f,Frozen:%.4f", account.SubAccounts[BTC].Amount, account.SubAccounts[BTC].ForzenAmount)
	log.Printf("[LTC Asset] Avaliable:%.4f,Frozen:%.4f", account.SubAccounts[LTC].Amount, account.SubAccounts[LTC].ForzenAmount)
	log.Printf("[USDT Asset] Avaliable:%.4f,Frozen:%.4f", account.SubAccounts[USDT].Amount, account.SubAccounts[USDT].ForzenAmount)

}

```


	


