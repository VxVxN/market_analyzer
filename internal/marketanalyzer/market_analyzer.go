package marketanalyzer

import (
	"math/big"

	"github.com/VxVxN/market_analyzer/pkg/tools"
)

type MarketAnalyzer struct {
	data *RawMarketData

	// quarter analyze by quarter
	quarter Quarter
}

type Quarter int

const (
	FirstQuarter Quarter = iota
	SecondQuarter
	ThirdQuarter
	FourthQuarter
	AllQuarters
)

func Init(data *RawMarketData) *MarketAnalyzer {
	return &MarketAnalyzer{
		data:    data,
		quarter: AllQuarters,
	}
}

func (analyzer *MarketAnalyzer) Calculate() *MarketData {
	calculatedData := &MarketData{
		Quarters: analyzer.data.Quarters,
	}

	changedData := analyzer.changeDataByQuarter()
	calculatedData.Quarters = changedData.Quarters
	calculatedData.RawData = changedData.Data

	data := make(map[RowName][]*big.Float)
	for name, records := range changedData.Data {
		var calculatedData []*big.Float
		lastRecord := new(big.Float)
		for _, record := range records {
			result := new(big.Float).SetInt(record)
			if lastRecord.Sign() == 0 {
				lastRecord.SetInt(record)
			} else if record.Sign() != 0 {
				result = new(big.Float).Quo(result, lastRecord)
				lastRecord.SetInt(record)
			}
			calculatedData = append(calculatedData, result)
		}
		data[name] = calculatedData
	}
	calculatedData.PercentageChanges = data
	return calculatedData
}

func (analyzer *MarketAnalyzer) SetQuarter(quarter Quarter) {
	analyzer.quarter = quarter
}

func (analyzer *MarketAnalyzer) changeDataByQuarter() *RawMarketData {
	if analyzer.quarter == AllQuarters {
		return analyzer.data
	}
	changedData := new(RawMarketData)
	var resultQuarters []YearQuarter
	var skipIndex []int

	for i, quarter := range analyzer.data.Quarters {
		switch analyzer.quarter {
		case FirstQuarter:
			if quarter.Quarter == 1 {
				resultQuarters = append(resultQuarters, quarter)
			} else {
				skipIndex = append(skipIndex, i)
			}
		case SecondQuarter:
			if quarter.Quarter == 2 {
				resultQuarters = append(resultQuarters, quarter)
			} else {
				skipIndex = append(skipIndex, i)
			}
		case ThirdQuarter:
			if quarter.Quarter == 3 {
				resultQuarters = append(resultQuarters, quarter)
			} else {
				skipIndex = append(skipIndex, i)
			}
		case FourthQuarter:
			if quarter.Quarter == 4 {
				resultQuarters = append(resultQuarters, quarter)
			} else {
				skipIndex = append(skipIndex, i)
			}
		}
	}

	changedData.Quarters = resultQuarters
	data := make(map[RowName][]*big.Int)
	for name, records := range analyzer.data.Data {
		var recordsCurrentRow []*big.Int
		for i, record := range records {
			if !tools.ContainNumberInSlice(i, skipIndex) {
				recordsCurrentRow = append(recordsCurrentRow, record)
			}
		}
		data[name] = recordsCurrentRow
	}
	changedData.Data = data
	return changedData
}
