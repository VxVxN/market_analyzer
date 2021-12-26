package myfileparser

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/VxVxN/market_analyzer/internal/humanizer"
)

type Parser struct {
	filePath  string
	readyData *humanizer.ReadyData
}

func Init(filePath string) *Parser {
	return &Parser{
		readyData: new(humanizer.ReadyData),
		filePath:  filePath,
	}
}

func (parser *Parser) Parse() (*humanizer.ReadyData, error) {
	file, err := os.Open(parser.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}
	if err = parser.parseHeaders(headers); err != nil {
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

		parser.readyData.Rows = append(parser.readyData.Rows, record)
	}
	return parser.readyData, nil
}

func (parser *Parser) parseHeaders(headers []string) error {
	var readyHeaders []string
	for _, header := range headers {
		readyHeaders = append(readyHeaders, header)
	}
	parser.readyData.Headers = readyHeaders
	return nil
}
