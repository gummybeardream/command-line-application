/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// defining the housing struct and add JSON struct tags
type Housing struct {
	Value    int     `json:"value"`
	Income   float64 `json:"income"`
	Age      int     `json:"age"`
	Rooms    int     `json:"rooms"`
	Bedrooms int     `json:"bedrooms"`
	Pop      int     `json:"pop"`
	Hh       int     `json:"hh"`
}

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

		//return your error if file does not open
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}

		defer file.Close()

		//create csv reader
		reader := csv.NewReader(file)
		header, err := reader.Read()

		//return error if header cannot be read
		if err != nil {
			return fmt.Errorf("failed to read CSV: %w", err)
		}

		//print header here
		fmt.Println(header)

		record, err := reader.Read()

		//return error if first row cannot be read
		if err != nil {
			return fmt.Errorf("failed to read first CSV record: %w", err)

		}

		//print first row here
		fmt.Println(record)

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
