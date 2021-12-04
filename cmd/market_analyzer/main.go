package main

import (
	"fmt"
	"log"

	"market_analyzer/internal/parser/myfileparser"
)

func main() {
	var err error
	parser := myfileparser.Init("data/fixp.csv")

	if err = parser.Parse(); err != nil {
		log.Fatalln(err)
	}
	marketData := parser.GetData()

	for _, quarter := range marketData.Quarters {
		fmt.Print(quarter.Year, "/", quarter.Quarter, "\t")
	}
	fmt.Print("\n")
	for _, records := range marketData.Data {
		for _, record := range records {
			fmt.Print(record, "\t")
		}
		fmt.Print("\n")
	}
}
