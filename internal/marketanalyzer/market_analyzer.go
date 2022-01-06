package marketanalyzer

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/VxVxN/market_analyzer/pkg/tools"
)

type MarketAnalyzer struct {
	data *RawMarketData

	periodMode PeriodMode
}

func Init(data *RawMarketData) *MarketAnalyzer {
	return &MarketAnalyzer{
		data:       data,
		periodMode: NormalMode,
	}
}

func (analyzer *MarketAnalyzer) Calculate() (*MarketData, error) {
	calculatedData := &MarketData{
		Quarters: analyzer.data.YearQuarters,
	}

	data := analyzer.changeDataByPeriodMode()
	calculatedData.Quarters = data.YearQuarters
	calculatedData.RawData = data.Data
	var err error
	calculatedData.Multipliers, err = analyzer.calculateMultipliers(data)
	if err != nil {
		return nil, fmt.Errorf("error calculate multipliers: %v", err)
	}
	analyzer.calculatePercentageChanges(data, calculatedData)
	return calculatedData, nil
}

func (analyzer *MarketAnalyzer) calculateMultipliers(data *RawMarketData) (map[MultiplierName][]*big.Float, error) {
	multipliers := make(map[MultiplierName][]*big.Float)
	marketCaps, ok := data.Data[MarketCap]
	if !ok {
		return nil, errors.New("market capacity not found")
	}
	sales, ok := data.Data[Sales]
	if !ok {
		fmt.Println("warn: sales not found")
	}
	earnings, ok := data.Data[Earnings]
	if !ok {
		fmt.Println("warn: earnings not found")
	}
	var pe []*big.Float
	var ps []*big.Float

	for i, marketCap := range marketCaps {
		marketCapFloat := new(big.Float).SetInt(marketCap)
		earningsFloat := new(big.Float)
		if len(earnings) > i {
			earningsFloat = earningsFloat.SetInt(earnings[i])
		}
		salesFloat := new(big.Float)
		if len(sales) > i {
			salesFloat = salesFloat.SetInt(sales[i])
		}

		if earningsFloat.Sign() == 0 {
			pe = append(pe, new(big.Float))
		} else {
			pe = append(pe, new(big.Float).Quo(marketCapFloat, earningsFloat))
		}

		if salesFloat.Sign() == 0 {
			ps = append(ps, new(big.Float))
		} else {
			ps = append(ps, new(big.Float).Quo(marketCapFloat, salesFloat))
		}
	}
	multipliers[PE] = pe
	multipliers[PS] = ps

	return multipliers, nil
}

func (analyzer *MarketAnalyzer) calculatePercentageChanges(changedData *RawMarketData, calculatedData *MarketData) {
	data := make(map[RowName][]*big.Float)
	for name, records := range changedData.Data {
		var calculatedData []*big.Float
		lastRecord := new(big.Float)
		for _, record := range records {
			currentRecord := new(big.Float).SetInt(record)
			result := new(big.Float)
			if lastRecord.Sign() == 0 {
				lastRecord.SetInt(record)
				result.Set(lastRecord)
			} else if record.Sign() != 0 {
				changePrice := new(big.Float).Sub(currentRecord, lastRecord)
				result = new(big.Float).Quo(changePrice, lastRecord)
				if (currentRecord.Sign() == -1 && lastRecord.Sign() == -1) || (currentRecord.Sign() == 1 && lastRecord.Sign() == -1) {
					result.Mul(result, big.NewFloat(-1))
				}

				lastRecord.SetInt(record)
			}
			calculatedData = append(calculatedData, result)
		}
		data[name] = calculatedData
	}
	calculatedData.PercentageChanges = data
}

func (analyzer *MarketAnalyzer) SetPeriodMode(quarter PeriodMode) {
	analyzer.periodMode = quarter
}

func (analyzer *MarketAnalyzer) changeDataByPeriodMode() *RawMarketData {
	if analyzer.periodMode == NormalMode {
		return analyzer.data
	}
	changedData := new(RawMarketData)
	var resultQuarters []YearQuarter
	var skipIndex []int
	var yearsWithQuarters []struct {
		year     int
		quarters []int
	}

	for i, yearQuarter := range analyzer.data.YearQuarters {
		switch analyzer.periodMode {
		case FirstQuarterMode:
			if yearQuarter.Quarter == 1 {
				resultQuarters = append(resultQuarters, yearQuarter)
			} else {
				skipIndex = append(skipIndex, i)
			}
		case SecondQuarterMode:
			if yearQuarter.Quarter == 2 {
				resultQuarters = append(resultQuarters, yearQuarter)
			} else {
				skipIndex = append(skipIndex, i)
			}
		case ThirdQuarterMode:
			if yearQuarter.Quarter == 3 {
				resultQuarters = append(resultQuarters, yearQuarter)
			} else {
				skipIndex = append(skipIndex, i)
			}
		case FourthQuarterMode:
			if yearQuarter.Quarter == 4 {
				resultQuarters = append(resultQuarters, yearQuarter)
			} else {
				skipIndex = append(skipIndex, i)
			}
		case YearMode:
			var index int
			for i, yq := range yearsWithQuarters {
				if yq.year == yearQuarter.Year {
					index = i
					break
				}
			}
			if index != 0 {
				yearsWithQuarters[index].quarters = append(yearsWithQuarters[index].quarters, yearQuarter.Quarter)
			} else {
				yearsWithQuarters = append(
					yearsWithQuarters, struct {
						year     int
						quarters []int
					}{year: yearQuarter.Year, quarters: []int{yearQuarter.Quarter}})
			}
		}
	}

	if analyzer.periodMode == YearMode {
		analyzer.groupByYears(changedData, yearsWithQuarters)
		return changedData
	}

	changedData.YearQuarters = resultQuarters
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

func (analyzer *MarketAnalyzer) groupByYears(
	changedData *RawMarketData, yearsWithQuarters []struct {
		year     int
		quarters []int
	},
) {
	changedData.YearQuarters = []YearQuarter{}
	for _, yearQuarter := range yearsWithQuarters {
		changedData.YearQuarters = append(
			changedData.YearQuarters, YearQuarter{
				Year: yearQuarter.year,
			},
		)
	}
	data := make(map[RowName][]*big.Int)
	var startSegment, endSegment int
	for _, quarters := range yearsWithQuarters {
		endSegment = startSegment + len(quarters.quarters)
		for name, records := range analyzer.data.Data {
			values := records[startSegment:endSegment]
			sumQuartersValue := new(big.Int)

			for _, value := range values {
				sumQuartersValue.Add(sumQuartersValue, value)
			}

			data[name] = append(data[name], sumQuartersValue)
		}
		startSegment = endSegment

	}
	changedData.Data = data
}
