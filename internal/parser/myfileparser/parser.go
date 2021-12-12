package myfileparser

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
)

type Parser struct {
	filePath   string
	marketData *marketanalyzer.RawMarketData
}

func Init(filePath string) *Parser {
	marketData := new(marketanalyzer.RawMarketData)
	marketData.Data = make(map[marketanalyzer.RowName][]*big.Int)
	return &Parser{
		marketData: marketData,
		filePath:   filePath,
	}
}

func (parser *Parser) Parse() (*marketanalyzer.RawMarketData, error) {
	file, err := os.Open(parser.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read() // quarters
	if err != nil {
		return nil, err
	}
	if err = parser.parseQuarters(headers); err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if err = parser.parseRow(record); err != nil {
			return nil, err
		}
	}
	return parser.marketData, nil
}

func (parser *Parser) parseQuarters(headers []string) error {
	var quarters []marketanalyzer.YearQuarter
	for i, header := range headers {
		if i == 0 {
			continue // skip empty string
		}
		splitHeader := strings.Split(header, "/")
		if len(splitHeader) != 2 {
			return fmt.Errorf("invalid header, expected format: [year/quater], actual: %s", header)
		}
		year, err := strconv.Atoi(splitHeader[0])
		if err != nil {
			return fmt.Errorf("invalid year, expected integer, record: %s", splitHeader[0])
		}
		quarter, err := strconv.Atoi(splitHeader[1])
		if err != nil {
			return fmt.Errorf("invalid quarter, expected integer, record: %s", splitHeader[1])
		}
		quarters = append(quarters, marketanalyzer.YearQuarter{
			Year:    year,
			Quarter: quarter,
		})
	}
	parser.marketData.Quarters = quarters
	return nil
}

func (parser *Parser) parseRow(records []string) error {
	rowName := marketanalyzer.RowName(records[0])
	var data []*big.Int
	for i, record := range records {
		if i == 0 {
			continue // skip rowName
		}
		if record == "" {
			data = append(data, new(big.Int))
			continue
		}
		recNumber, err := strconv.Atoi(record)
		if err != nil {
			return fmt.Errorf("invalid record %s, expected integer, record: %s", rowName, record)
		}

		data = append(data, big.NewInt(int64(recNumber)))
	}
	parser.marketData.Data[rowName] = data

	return nil
}
