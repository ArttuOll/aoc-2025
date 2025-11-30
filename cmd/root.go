/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/ArttuOll/aoc-2025/internal/solution"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc-2025 <day-number> <a|b>",
	Short: "Program for solving Advent of Code 2025",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		return solution.Run(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
