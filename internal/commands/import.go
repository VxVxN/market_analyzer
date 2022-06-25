package commands

import (
	"log"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/VxVxN/market_analyzer/internal/consts"
	"github.com/VxVxN/market_analyzer/internal/parser/smartlabparser"
	"github.com/VxVxN/market_analyzer/internal/saver/csvsaver"
)

func InitImportCmd() *cobra.Command {
	var nameFlag string

	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import smart lab file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			parser := smartlabparser.Init()
			parser.SetFilePath(args[0])
			data, err := parser.Parse()
			if err != nil {
				log.Fatalln(err)
			}
			if nameFlag == "" {
				nameFlag = filepath.Base(args[0])
			}
			saver := csvsaver.Init("data/emitters/"+nameFlag+consts.CsvFileExtension, data.Headers, data.Rows)
			if err = saver.Save(); err != nil {
				log.Fatalln(err)
			}
		},
	}

	cmd.Flags().StringVarP(&nameFlag, "name", "n", "", "Sets name of emitter")

	return cmd
}
