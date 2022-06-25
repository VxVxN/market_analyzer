package commands

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/VxVxN/market_analyzer/internal/consts"
)

func InitListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Print list of emitters",
		Run: func(cmd *cobra.Command, args []string) {
			filepath.WalkDir("data/emitters", func(path string, d fs.DirEntry, err error) error {
				if d.IsDir() {
					return nil
				}
				emitter := strings.Replace(filepath.Base(path), consts.CsvFileExtension, "", 1)
				fmt.Println("List of emitters:")
				fmt.Println(emitter)
				return nil
			})
		},
	}

	return cmd
}
