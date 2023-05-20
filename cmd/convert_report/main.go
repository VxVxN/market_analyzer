package main

import (
	"github.com/VxVxN/market_analyzer/internal/saver/csvsaver"
	"log"
	"os"
	"path/filepath"

	"github.com/VxVxN/market_analyzer/internal/parser/smartlabparser"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("two arguments were expected")
	}
	inputFile := os.Args[1]
	parser := smartlabparser.Init()
	parser.SetFilePath(inputFile)
	data, err := parser.Parse()
	if err != nil {
		log.Fatalln(err)
	}
	fileOutput := filepath.Join(os.Args[2], filepath.Base(inputFile))
	saver := csvsaver.Init(fileOutput, data.Headers, data.Rows)
	if err = saver.Save(); err != nil {
		log.Fatalln(err)
	}
}
