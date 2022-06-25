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

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
