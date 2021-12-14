package marketanalyzer

import "math/big"

type RawMarketData struct {
	YearQuarters []YearQuarter
	Data         map[RowName][]*big.Int
}

type MarketData struct {
	Quarters          []YearQuarter
	PercentageChanges map[RowName][]*big.Float
	RawData           map[RowName][]*big.Int
}

type YearQuarter struct {
	Year    int
	Quarter int
}

type RowName string

const (
	Sales    RowName = "sales"
	Earnings         = "earnings"
	Debts            = "debts"
)
