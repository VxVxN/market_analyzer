package parser

import (
	"github.com/VxVxN/market_analyzer/internal/humanizer"
)

type Parser interface {
	Parse() error
	GetData() *humanizer.ReadyData
}
