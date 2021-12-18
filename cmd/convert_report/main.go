package main

import (
	"log"
	"os"

	"github.com/VxVxN/market_analyzer/internal/parser/smartlabparser"
	"github.com/VxVxN/market_analyzer/internal/saver/csvsaver"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("two arguments were expected")
	}
	parser := smartlabparser.Init(os.Args[1])
	data, err := parser.Parse()
	if err != nil {
		log.Fatalln(err)
	}
	saver := csvsaver.Init(os.Args[2], data.Headers, data.Rows)
	if err = saver.Save(); err != nil {
		log.Fatalln(err)
	}
}
