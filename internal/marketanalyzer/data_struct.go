package marketanalyzer

import "math/big"

type RawMarketData struct {
	Quarters []Quarter
	Data     map[RowName][]*big.Int
}

type CalculatedMarketData struct {
	Quarters []Quarter
	Data     map[RowName][]*big.Float
}

type Quarter struct {
	Year    int
	Quarter int
}

type RowName string

const (
	Sales    RowName = "sales"
	Earnings         = "earnings"
)
