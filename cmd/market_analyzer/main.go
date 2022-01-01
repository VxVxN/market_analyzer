package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/VxVxN/market_analyzer/internal/consts"
	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/internal/parser/myfileparser"
	preparerpkg "github.com/VxVxN/market_analyzer/internal/preparer"
	"github.com/VxVxN/market_analyzer/internal/saver/csvsaver"
)

func main() {
	precisionFlag := flag.Int("precision", 2, "sets the number of digits after the dot for percentages")
	periodFlag := flag.String("period", marketanalyzer.NormalMode.String(), "sets the number of digits after the dot for percentages")
	numberFlag := flag.String("number", hum.NumbersWithPercentagesMode.String(), "sets the number mode")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("enter the command")
		os.Exit(1)
	}
	command := os.Args[1]

	switch command {
	case "list":
		filepath.WalkDir("data/emitters", func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			emitter := strings.Replace(filepath.Base(path), consts.CsvFileExtension, "", 1)
			fmt.Println(emitter)
			return nil
		})
		os.Exit(0)
	case "report":
		if len(os.Args) < 3 {
			fmt.Println("enter a name emitter")
			os.Exit(1)
		}
		emitter := os.Args[2]
		report(emitter, periodFlag, numberFlag, precisionFlag)
	default:
		fmt.Println("command doesn't exist")
		os.Exit(1)
	}
}

func report(emitter string, periodFlag *string, numberFlag *string, precisionFlag *int) {
	parser := myfileparser.Init("data/emitters/" + emitter + consts.CsvFileExtension)

	rawData, err := parser.Parse()
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	preparer := preparerpkg.Init(rawData)
	rawMarketData, err := preparer.Prepare()
	if err != nil {
		log.Fatalf("Failed to prepare data: %v", err)
	}

	periodMode, err := marketanalyzer.PeriodModeFromString(*periodFlag)
	if err != nil {
		log.Fatalf("Failed to parse period mode from string: %v", err)
	}

	analyzer := marketanalyzer.Init(rawMarketData)
	analyzer.SetPeriodMode(periodMode)
	marketData, err := analyzer.Calculate()
	if err != nil {
		log.Fatalf("Failed to calculate market model: %v", err)
	}

	numberMode, err := hum.NumberModeFromString(*numberFlag)
	if err != nil {
		log.Fatalf("Failed to parse number mode from string: %v", err)
	}

	humanizer := hum.Init(marketData)
	humanizer.SetPrecision(*precisionFlag)
	humanizer.SetNumbersMode(numberMode)
	humanizer.SetFieldsForDisplay(
		[]marketanalyzer.RowName{
			// marketanalyzer.Sales,
			// marketanalyzer.Earnings,
		},
	)
	data := humanizer.Humanize()

	saver := csvsaver.Init("data/saved_reports/humanize_data.csv", data.Headers, data.Rows)
	if err = saver.Save(); err != nil {
		log.Fatalf("Failed to save file: %v", err)
	}

	// printer := p.Init()
	// printer.Print(data)
}
