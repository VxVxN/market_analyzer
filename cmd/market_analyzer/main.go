package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/VxVxN/market_analyzer/internal/commands"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "app",
	}

	rootCmd.AddCommand(commands.InitListCmd())
	rootCmd.AddCommand(commands.InitReportCmd())
	rootCmd.AddCommand(commands.InitImportCmd())
	rootCmd.AddCommand(commands.InitWebCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
