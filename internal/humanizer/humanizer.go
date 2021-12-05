package humanizer

import (
	"fmt"
	"math/big"

	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
	"github.com/VxVxN/market_analyzer/pkg/tools"
)

type Humanizer struct {
	rawMarketData *marketanalyzer.RawMarketData
	marketData    *marketanalyzer.CalculatedMarketData

	precision int
}

type ReadyData struct {
	Headers []string
	Rows    [][]string
}

func Init(marketData *marketanalyzer.CalculatedMarketData, rawMarketData *marketanalyzer.RawMarketData) *Humanizer {
	return &Humanizer{
		marketData:    marketData,
		rawMarketData: rawMarketData,
		precision:     0,
	}
}

func (humanizer *Humanizer) Humanize() *ReadyData {
	data := &ReadyData{
		Headers: []string{"#"},
	}
	for _, quarter := range humanizer.marketData.Quarters {
		data.Headers = append(data.Headers, fmt.Sprint(quarter.Year, "/", quarter.Quarter))
	}

	for name, records := range humanizer.marketData.Data {
		row := []string{string(name)}
		for i, record := range records {
			switch {
			case record.Sign() == 0:
				row = append(row, "-")
			case record.IsInt():
				str := record.Text('g', 100)
				str = tools.HumanizeNumber(str)
				row = append(row, str)
			default:
				rawData := humanizer.rawMarketData.Data[name][i]
				sign := "+"
				if record.Cmp(big.NewFloat(1)) == -1 {
					sign = "-"
				}
				result := new(big.Float).Mul(record, new(big.Float).SetInt64(100))

				rawDataStr := rawData.String()
				rawDataStr = tools.HumanizeNumber(rawDataStr)

				str := fmt.Sprintf("%s(%s%s%s)", rawDataStr, sign, result.Text('f', humanizer.precision), "%")
				row = append(row, str)
			}
		}
		data.Rows = append(data.Rows, row)
	}
	return data
}

func (humanizer *Humanizer) SetPrecision(precision int) {
	humanizer.precision = precision
}
