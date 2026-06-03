package main

import (
	"fmt"
)

func main() {
	// liq_price 强平价
	var liq_price int = 136500
	// entry_price 开仓价
	var entry_price int = 91200
	// leverage  实际杠杆
	var leverage int = 1
	// margin 保证金
	var margin int = 5000
	// size_btc 持仓数量btc
	var size_btc float64
	//size_usdt 持仓数量usdt
	var size_usdt int

	//d0 := decimal.NewFromFloat(0)
	decimal.DivisionPrecision = 3
	size_btc = decimal.NewFromFloat(margin).Div(decimal.NewFromFloat(liq_price - entry_price))
	fmt.Println(size_btc)
	fmt.Println(leverage)
	fmt.Println(size_usdt)

}
