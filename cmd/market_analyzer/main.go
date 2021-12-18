package main

import (
	"log"

	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/internal/parser/myfileparser"
	preparerpkg "github.com/VxVxN/market_analyzer/internal/preparer"
	p "github.com/VxVxN/market_analyzer/internal/printer"
	"github.com/VxVxN/market_analyzer/internal/saver/csvsaver"
)

func main() {
	parser := myfileparser.Init("data/emitters/ozon.csv")

	rawData, err := parser.Parse()
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	preparer := preparerpkg.Init(rawData)
	rawMarketData, err := preparer.Prepare()
	if err != nil {
		log.Fatalf("Failed to prepare data: %v", err)
	}

	analyzer := marketanalyzer.Init(rawMarketData)
	analyzer.SetPeriodMode(marketanalyzer.NormalMode)
	marketData := analyzer.Calculate()

	humanizer := hum.Init(marketData)
	humanizer.SetPrecision(2)
	humanizer.SetNumbersMode(hum.NumbersWithPercentagesMode)
	humanizer.SetFieldsForDisplay([]marketanalyzer.RowName{
		// marketanalyzer.Sales,
		// marketanalyz
		// er.Earnings,
	})
	data := humanizer.Humanize()

	saver := csvsaver.Init("data/saved_reports/humanize_data.csv", data.Headers, data.Rows)
	if err = saver.Save(); err != nil {
		log.Fatalf("Failed to save file: %v", err)
	}

	printer := p.Init()
	printer.Print(data)
}
