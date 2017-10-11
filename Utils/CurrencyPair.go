package utils

import "strings"

type Currency struct {
	Symbol string
	Desc   string
}

func (c Currency) String() string {
	return c.Symbol
}

// A->B(A兑换为B)
type CurrencyPair struct {
	CurrencyA Currency
	CurrencyB Currency
}

var (
	UNKNOWN = Currency{"UNKNOWN", ""}
	CNY     = Currency{"CNY", "rmb （China Yuan)"}
	USD     = Currency{"USD", "USA dollar"}
	USDT    = Currency{"USDT", "Tether dollar"}
	EUR     = Currency{"EUR", ""}
	BTC     = Currency{"BTC", "bitcoin.org"}
	BCC     = Currency{"BCC", "bitcoin-abc"}
	BCH     = Currency{"BCH", "bitcoin-abc"}
	LTC     = Currency{"LTC", "litecoin.org"}
	ETH     = Currency{"ETH", ""}
	ETC     = Currency{"ETC", ""}

	//currency pair
	BTC_USDT = CurrencyPair{BTC, USDT}
	LTC_USDT = CurrencyPair{LTC, USDT}
	ETH_USDT = CurrencyPair{ETH, USDT}
	ETC_USDT = CurrencyPair{ETC, USDT}
	BCH_USDT = CurrencyPair{BCH, USDT}
	BCC_USDT = CurrencyPair{BCC, USDT}

	LTC_BTC = CurrencyPair{LTC, BTC}
	ETH_BTC = CurrencyPair{ETH, BTC}
	ETC_BTC = CurrencyPair{ETC, BTC}
	BCC_BTC = CurrencyPair{BCC, BTC}
	BCH_BTC = CurrencyPair{BCH, BTC}

	ETC_ETH      = CurrencyPair{ETC, ETH}
	UNKNOWN_PAIR = CurrencyPair{UNKNOWN, UNKNOWN}

	CurrencyArray     = []Currency{CNY, USD, USDT, EUR, BTC, BCC, BCH, LTC, ETH, ETC}
	CurrencyPairArray = []CurrencyPair{BTC_USDT, LTC_USDT, ETH_USDT, ETC_USDT, BCH_USDT, BCC_USDT, LTC_BTC, ETH_BTC, ETC_BTC, BCC_BTC, BCH_BTC, ETC_ETH}
)

func (c CurrencyPair) String() string {
	return c.ToSymbol("_")
}

func NewCurrency(symbol, desc string) Currency {
	return Currency{symbol, desc}
}

func NewCurrencyPair(currencyA Currency, currencyB Currency) CurrencyPair {
	return CurrencyPair{currencyA, currencyB}
}

func NewCurrencyPair2(currencyPairSymbol string) CurrencyPair {
	currencys := strings.Split(currencyPairSymbol, "_")
	if len(currencys) == 2 {
		return CurrencyPair{NewCurrency(currencys[0], ""),
			NewCurrency(currencys[1], "")}
	}
	return UNKNOWN_PAIR
}

func (pair CurrencyPair) ToSymbol(joinChar string) string {
	return strings.Join([]string{pair.CurrencyA.Symbol, pair.CurrencyB.Symbol}, joinChar)
}

func (pair CurrencyPair) ToSymbol2(joinChar string) string {
	return strings.Join([]string{pair.CurrencyB.Symbol, pair.CurrencyA.Symbol}, joinChar)
}
