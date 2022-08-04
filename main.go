package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type QuoteEstimate struct {
	EstimatedAmountIn      string `json:"estimatedAmountIn"`
	EstimatedAmountOut     string `json:"estimatedAmountOut"`
	EstimatedEndTickIndex  int    `json:"estimatedEndTickIndex"`
	EstimatedEndSqrtPrice  string `json:"estimatedEndSqrtPrice"`
	EstimatedFeeAmount     string `json:"estimatedFeeAmount"`
	Amount                 string `json:"amount"`
	AmountSpecifiedIsInput bool   `json:"amountSpecifiedIsInput"`
	AToB                   bool   `json:"aToB"`
	OtherAmountThreshold   string `json:"otherAmountThreshold"`
	SqrtPriceLimit         string `json:"sqrtPriceLimit"`
	TickArray0             string `json:"tickArray0"`
	TickArray1             string `json:"tickArray1"`
	TickArray2             string `json:"tickArray2"`
}

func main() {
	command := "ts-node scripts/quote.ts" +
		" CRR7huZnXaiBjGGMAU6iVeQU9b2g71NXiLHA6g29DeYN" +
		" 57K3gMtUMctYGYUpm9PjzYQeiCV8BeRkSuuBFGkuWAdt" +
		" Dphoc5nPvC5eadUP79McRB36hgKcetgJ7BRG5Zv6QeYp" +
		" 57K3gMtUMctYGYUpm9PjzYQeiCV8BeRkSuuBFGkuWAdt"
	parts := strings.Fields(command)
	data, err := exec.Command(parts[0], parts[1:]...).Output()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var quote QuoteEstimate
	json.Unmarshal(data, &quote)
	fmt.Println(quote)
}
