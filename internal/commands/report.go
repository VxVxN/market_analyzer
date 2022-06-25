package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/VxVxN/market_analyzer/internal/consts"
	hum "github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/internal/parser/myfileparser"
	preparerpkg "github.com/VxVxN/market_analyzer/internal/preparer"
	"github.com/VxVxN/market_analyzer/internal/saver/csvsaver"
)

func InitReportCmd() *cobra.Command {
	var precisionFlag int
	var periodFlag, numberFlag string
	var reportCmd = &cobra.Command{
		Use:   "report [name of emitter]",
		Short: "Generate report about emitter",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			report(args[0], &periodFlag, &numberFlag, &precisionFlag)
		},
	}
	reportCmd.Flags().IntVarP(&precisionFlag, "precision", "", 2, "Sets the number of digits after the dot for percentages")
	reportCmd.Flags().StringVarP(&periodFlag, "period", "", marketanalyzer.NormalMode.String(), "Sets the period mode")
	reportCmd.Flags().StringVarP(&numberFlag, "number", "", hum.NumbersWithPercentagesMode.String(), "Sets the number mode")

	return reportCmd
}

func report(emitter string, periodFlag *string, numberFlag *string, precisionFlag *int) {
	parser := myfileparser.Init("data/emitters/" + emitter + consts.CsvFileExtension)

	rawData, err := parser.Parse()
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s emmiter not found\n", emitter)
			fmt.Println("Use the list command to view the list of emitters")
			os.Exit(1)
		} else {
			log.Fatalf("Failed to parse file: %v", err)
		}
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

	fileName := fmt.Sprintf("data/saved_reports/%s.csv", emitter)
	saver := csvsaver.Init(fileName, data.Headers, data.Rows)
	if err = saver.Save(); err != nil {
		log.Fatalf("Failed to save file: %v", err)
	}
	fmt.Printf("Report is generated: %s\n", fileName)

	// printer := p.Init()
	// printer.Print(data)
}
