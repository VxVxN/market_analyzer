package printer

import (
	"fmt"

	"github.com/alexeyco/simpletable"

	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
)

type Printer struct {
	marketData *marketanalyzer.RawMarketData
}

func Init() *Printer {
	return &Printer{}
}

func (printer *Printer) SetMarketData(data *marketanalyzer.RawMarketData) {
	printer.marketData = data
}

func (printer *Printer) Print() {
	table := simpletable.New()

	headers := []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Text: "#"},
	}

	for _, quarter := range printer.marketData.Quarters {
		headers = append(headers, &simpletable.Cell{
			Align: simpletable.AlignCenter,
			Text:  fmt.Sprint(quarter.Year, "/", quarter.Quarter),
		})
	}
	table.Header = &simpletable.Header{
		Cells: headers,
	}

	for rowName, row := range printer.marketData.Data {
		records := []*simpletable.Cell{
			{
				Align: simpletable.AlignRight,
				Text:  string(rowName),
			},
		}
		for _, record := range row {
			records = append(records, &simpletable.Cell{
				Align: simpletable.AlignRight,
				Text:  record.String(),
			})
		}
		table.Body.Cells = append(table.Body.Cells, records)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}
