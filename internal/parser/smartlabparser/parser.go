// Package smartlabparser parse report from https://smart-lab.ru/
package smartlabparser

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/VxVxN/market_analyzer/internal/humanizer"
	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
)

type Parser struct {
	filePath    string
	fileData    []byte
	isLTMHeader bool
	readyData   *humanizer.ReadyData
}

func Init() *Parser {
	return &Parser{
		readyData: new(humanizer.ReadyData),
	}
}

var parsedRows = map[string]marketanalyzer.RowName{
	"Выручка, млрд руб":        marketanalyzer.Sales,
	"Чистая прибыль, млрд руб": marketanalyzer.Earnings,
	"Долг, млрд руб":           marketanalyzer.Debts,
	"Капитализация, млрд руб":  marketanalyzer.MarketCap,
}

func (parser *Parser) Parse() (*humanizer.ReadyData, error) {
	var reader io.Reader

	if len(parser.fileData) != 0 {
		reader = bytes.NewReader(parser.fileData)
	} else {
		fileData, err := os.ReadFile(parser.filePath)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(fileData)
	}

	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'

	headers, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	if err = parser.parseHeaders(headers); err != nil {
		return nil, err
	}

	for {
		record, err := csvReader.Read()
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
	return parser.readyData, nil
}

func (parser *Parser) parseHeaders(headers []string) error {
	readyHeaders := []string{""}
	for i, header := range headers {
		if i == 0 {
			continue // skip empty string
		}
		splitHeader := strings.Split(header, "Q")
		if len(splitHeader) != 2 {
			if header == "LTM" {
				parser.isLTMHeader = true
				continue // TODO LTM
			}
			return fmt.Errorf("invalid header, expected format: [yearQquater], actual: %s", header)
		}
		year, err := strconv.Atoi(splitHeader[0])
		if err != nil {
			return fmt.Errorf("invalid year, expected integer, record: %s", splitHeader[0])
		}
		quarter, err := strconv.Atoi(splitHeader[1])
		if err != nil {
			return fmt.Errorf("invalid quarter, expected integer, record: %s", splitHeader[1])
		}
		readyHeaders = append(readyHeaders, fmt.Sprintf("%d/%d", year, quarter))
	}
	parser.readyData.Headers = readyHeaders
	return nil
}

func (parser *Parser) parseRow(records []string) error {
	rowName, ok := parsedRows[records[0]]
	if !ok {
		return nil // skip excess row
	}

	readyRecords := []string{string(rowName)}

	for i, record := range records {
		if i == 0 {
			continue // skip rowName
		}

		if parser.isLTMHeader && i == len(records)-1 {
			continue // skip LTM column
		}

		if record == "" {
			readyRecords = append(readyRecords, "")
			continue
		}
		record = strings.Replace(record, " ", "", -1)
		recordFloat, err := strconv.ParseFloat(record, 64)
		if err != nil {
			return fmt.Errorf("invalid record %s, expected integer, record: %s", rowName, record)
		}

		recordBigFloat := big.NewFloat(recordFloat)
		recordBigFloat.Mul(recordBigFloat, big.NewFloat(1000000000))
		recordBigInt, _ := recordBigFloat.Int(nil)

		readyRecords = append(readyRecords, recordBigInt.String())
	}
	parser.readyData.Rows = append(parser.readyData.Rows, readyRecords)

	return nil
}

func (parser *Parser) SetFilePath(filePath string) {
	parser.filePath = filePath
}

func (parser *Parser) SetFileData(fileData []byte) {
	parser.fileData = fileData
}
