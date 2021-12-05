package printer

import (
	"fmt"

	"github.com/alexeyco/simpletable"

	"github.com/VxVxN/market_analyzer/internal/humanizer"
)

type Printer struct {
}

func Init() *Printer {
	return &Printer{}
}

func (printer *Printer) Print(data *humanizer.ReadyData) {
	table := simpletable.New()

	headers := make([]*simpletable.Cell, 0, len(data.Headers))

	for _, header := range data.Headers {
		headers = append(headers, &simpletable.Cell{
			Align: simpletable.AlignCenter,
			Text:  header,
		})
	}
	table.Header = &simpletable.Header{
		Cells: headers,
	}

	for _, row := range data.Rows {
		records := make([]*simpletable.Cell, 0, len(row))
		for _, record := range row {
			records = append(records, &simpletable.Cell{
				Align: simpletable.AlignRight,
				Text:  record,
			})
		}
		table.Body.Cells = append(table.Body.Cells, records)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}
