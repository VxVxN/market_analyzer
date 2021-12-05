package main

import (
	"log"

	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/internal/parser/myfileparser"
	p "github.com/VxVxN/market_analyzer/internal/printer"
)

func main() {
	parser := myfileparser.Init("data/fixp.csv")

	rawMarketData, err := parser.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	analyzer := marketanalyzer.Init(rawMarketData)
	marketData := analyzer.Calculate()

	humanizer := hum.Init(marketData, rawMarketData)
	humanizer.SetPrecision(2)
	data := humanizer.Humanize()

	printer := p.Init()
	printer.Print(data)
}
