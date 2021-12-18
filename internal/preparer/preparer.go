package preparer

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
)

type Preparer struct {
	data *humanizer.ReadyData
}

func Init(data *humanizer.ReadyData) *Preparer {
	return &Preparer{data: data}
}

func (preparer *Preparer) Prepare() (*marketanalyzer.RawMarketData, error) {
	rawData := new(marketanalyzer.RawMarketData)

	for i, header := range preparer.data.Headers {
		if i == 0 {
			continue // skip empty header
		}
		splitHeader := strings.Split(header, "/")
		if len(splitHeader) != 2 {
			return nil, fmt.Errorf("invalid header, expected format: [year/quater], actual: %s", header)
		}
		year, err := strconv.Atoi(splitHeader[0])
		if err != nil {
			return nil, fmt.Errorf("invalid year, expected integer, record: %s", splitHeader[0])
		}
		quarter, err := strconv.Atoi(splitHeader[1])
		if err != nil {
			return nil, fmt.Errorf("invalid quarter, expected integer, record: %s", splitHeader[1])
		}
		rawData.YearQuarters = append(rawData.YearQuarters, marketanalyzer.YearQuarter{
			Year:    year,
			Quarter: quarter,
		})
	}
	rawData.Data = make(map[marketanalyzer.RowName][]*big.Int)
	var ok bool
	for _, row := range preparer.data.Rows {
		rowName := marketanalyzer.RowName(row[0])

		var data []*big.Int
		for i, record := range row {
			if i == 0 {
				continue
			}
			var value *big.Int
			if record != "" {
				value, ok = new(big.Int).SetString(record, 0)
				if !ok {
					return nil, fmt.Errorf("invalid record, expected integer, record: %s", record)
				}
			} else {
				value = big.NewInt(0)
			}
			data = append(data, value)
		}
		rawData.Data[rowName] = data
	}
	return rawData, nil
}
