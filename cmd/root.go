/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "command line application [input CSV path] [output JSON lines path]",
	Short: "Convert housing CSV data to JSON lines",
	Long:  `The program reads a CSV file and converts it to a file with JSON lines. Only 2 arguments are required: CSV input and JSON lines file output`,
	Args:  cobra.ExactArgs(2),

	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath := args[0]

		//open the input file here
		file, err := os.Open(inputPath)

		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}

		defer file.Close()

		return nil
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.command-line-application.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
