package parser

import "github.com/VxVxN/market_analyzer/internal/marketanalyzer"

type Parser interface {
	Parse() error
	GetData() *marketanalyzer.RawMarketData
}
