package csvsaver

import (
	"encoding/csv"
	"os"
)

type Saver struct {
	fileName string
	headers  []string
	data     [][]string
}

func Init(fileName string, headers []string, data [][]string) *Saver {
	return &Saver{
		fileName: fileName,
		headers:  headers,
		data:     data,
	}
}

func (saver *Saver) Save() error {
	csvFile, err := os.Create(saver.fileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	if err = csvWriter.Write(saver.headers); err != nil {
		return err
	}
	for _, row := range saver.data {
		if err = csvWriter.Write(row); err != nil {
			return err
		}
	}
	csvWriter.Flush()
	return nil
}
