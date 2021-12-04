package main

import (
	"log"

	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/internal/parser/myfileparser"
	p "github.com/VxVxN/market_analyzer/internal/printer"
)

func main() {
	var err error
	parser := myfileparser.Init("data/fixp.csv")

	if err = parser.Parse(); err != nil {
		log.Fatalln(err)
	}
	rawMarketData := parser.GetData()

	analyzer := marketanalyzer.Init(rawMarketData)
	marketData := analyzer.Calculate()

	printer := p.Init()
	printer.SetMarketData(marketData)
	printer.Print()
}
