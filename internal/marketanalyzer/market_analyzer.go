package marketanalyzer

import "math/big"

type MarketAnalyzer struct {
	data *RawMarketData
}

func Init(data *RawMarketData) *MarketAnalyzer {
	return &MarketAnalyzer{
		data: data,
	}
}

func (analyzer *MarketAnalyzer) Calculate() *CalculatedMarketData {
	calculatedData := new(CalculatedMarketData)
	calculatedData.Quarters = analyzer.data.Quarters

	data := make(map[RowName][]*big.Float)
	for name, records := range analyzer.data.Data {
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
	calculatedData.Data = data
	return calculatedData
}
