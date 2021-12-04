package parser

import "market_analyzer/internal/marketanalyzer"

type Parser interface {
	Parse() error
	GetData() *marketanalyzer.RawMarketData
}
