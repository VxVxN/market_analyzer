package main

import (
	"log"

	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/internal/parser/myfileparser"
	p "github.com/VxVxN/market_analyzer/internal/printer"
	"github.com/VxVxN/market_analyzer/internal/saver/csvsaver"
)

func main() {
	parser := myfileparser.Init("data/fixp.csv")

	rawMarketData, err := parser.Parse()
	if err != nil {
		log.Fatalln(err)
	}
	analyzer := marketanalyzer.Init(rawMarketData)
	analyzer.SetPeriodMode(marketanalyzer.YearMode)
	marketData := analyzer.Calculate()

	humanizer := hum.Init(marketData)
	humanizer.SetPrecision(2)
	humanizer.SetNumbersMode(hum.NumbersWithPercentages)
	humanizer.SetFieldsForDisplay([]marketanalyzer.RowName{
		// marketanalyzer.Sales,
		// marketanalyz
		// er.Earnings,
	})
	data := humanizer.Humanize()

	saver := csvsaver.Init(data)
	if err = saver.Save("data/save/humanize_data.csv"); err != nil {
		log.Fatalln(err)
	}

	printer := p.Init()
	printer.Print(data)
}
