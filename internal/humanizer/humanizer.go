package humanizer

import (
	"fmt"
	"math/big"

	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/pkg/tools"
)

type Humanizer struct {
	marketData *marketanalyzer.MarketData

	precision        int
	mode             NumberMode
	fieldsForDisplay []marketanalyzer.RowName
	// fromYear, toYear int // for range
}

type ReadyData struct {
	Headers []string
	Rows    [][]string
}

type NumberMode int

const (
	NumbersWithPercentagesMode NumberMode = iota
	NumbersMode
	PercentagesMode
)

func Init(marketData *marketanalyzer.MarketData) *Humanizer {
	return &Humanizer{
		mode:       NumbersWithPercentagesMode,
		marketData: marketData,
		precision:  0,
	}
}

func (humanizer *Humanizer) Humanize() *ReadyData {
	order := []marketanalyzer.RowName{
		marketanalyzer.Sales,
		marketanalyzer.Earnings,
		marketanalyzer.Debts,
	}

	data := &ReadyData{
		Headers: []string{"#"},
	}
	for _, quarter := range humanizer.marketData.Quarters {
		if quarter.Quarter == 0 {
			data.Headers = append(data.Headers, fmt.Sprint(quarter.Year))
		} else {
			data.Headers = append(data.Headers, fmt.Sprint(quarter.Year, "/", quarter.Quarter))
		}
	}

	for _, name := range order {
		skipRow := true
		for _, rowName := range humanizer.fieldsForDisplay {
			if rowName == name {
				skipRow = false
			}
		}
		if skipRow && len(humanizer.fieldsForDisplay) != 0 {
			continue
		}
		records, ok := humanizer.marketData.PercentageChanges[name]
		if !ok {
			// TODO add warning
			continue
		}
		var isFindNumber bool
		row := []string{string(name)}
		for i, record := range records {
			switch {
			case record.Sign() == 0:
				row = append(row, "-")
			case record.IsInt() && !isFindNumber:
				isFindNumber = true
				str := record.Text('g', 100)
				str = tools.HumanizeNumber(str)
				row = append(row, str)
			default:
				rawData := humanizer.marketData.RawData[name][i]

				result := new(big.Float).Mul(record, new(big.Float).SetInt64(100))
				result.Sub(result, big.NewFloat(100))

				sign := "+"
				if result.Cmp(big.NewFloat(1)) == -1 {
					sign = ""
				}

				rawDataStr := rawData.String()
				rawDataStr = tools.HumanizeNumber(rawDataStr)

				var str string
				switch humanizer.mode {
				case NumbersWithPercentagesMode:
					str = fmt.Sprintf("%s(%s%s%s)", rawDataStr, sign, result.Text('f', humanizer.precision), "%")
				case NumbersMode:
					str = fmt.Sprintf("%s", rawDataStr)
				case PercentagesMode:
					str = fmt.Sprintf("%s%s%s", sign, result.Text('f', humanizer.precision), "%")
				}

				row = append(row, str)
			}
		}
		data.Rows = append(data.Rows, row)
	}
	return data
}

// SetPrecision sets the number of digits after the dot for percentages
func (humanizer *Humanizer) SetPrecision(precision int) {
	humanizer.precision = precision
}

func (humanizer *Humanizer) SetNumbersMode(mode NumberMode) {
	humanizer.mode = mode
}

func (humanizer *Humanizer) SetFieldsForDisplay(fieldsForDisplay []marketanalyzer.RowName) {
	humanizer.fieldsForDisplay = fieldsForDisplay
}

// func (humanizer *Humanizer) SetRange(fromYear, toYear int) {
// 	humanizer.fromYear = fromYear
// 	humanizer.toYear = toYear
// }
