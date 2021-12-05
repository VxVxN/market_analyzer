package csvsaver

import (
	"encoding/csv"
	"os"

	"github.com/VxVxN/market_analyzer/internal/humanizer"
)

type Saver struct {
	data *humanizer.ReadyData
}

func Init(data *humanizer.ReadyData) *Saver {
	return &Saver{
		data: data,
	}
}

func (saver *Saver) Save(fileName string) error {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	if err = csvWriter.Write(saver.data.Headers); err != nil {
		return err
	}
	for _, row := range saver.data.Rows {
		if err = csvWriter.Write(row); err != nil {
			return err
		}
	}
	csvWriter.Flush()
	return nil
}
