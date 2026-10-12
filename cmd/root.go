/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	ErrInvalidInputFile  = errors.New("input file is not a CSV")
	ErrInvalidOutputFile = errors.New("output file is not a .jl or .json file")
)

// defining the housing struct and add JSON struct tags
type Houses struct {
	Records []House `json:"houses"`
}

type House struct {
	Value    float64 `json:"value"`
	Income   float64 `json:"income"`
	Age      int64   `json:"age"`
	Rooms    int64   `json:"rooms"`
	Bedrooms int64   `json:"bedrooms"`
	Pop      int64   `json:"pop"`
	Hh       int64   `json:"hh"`
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
		outputPath := args[1]

		if err := ValidateArgs(inputPath, outputPath); err != nil {
			return fmt.Errorf("input validation failed: %w", err)
		}

		houses := Houses{
			Records: []House{},
		}

		//open the input file here
		file, err := os.Open(inputPath)

		//return your error if file does not open
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}

		defer file.Close()

		//create csv reader and read header
		reader := csv.NewReader(file)

		_, err = reader.Read()

		//return error if header cannot be read
		if err != nil {
			return fmt.Errorf("failed to read CSV: %w", err)
		}

		//read first record
		for {
			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			//return error if first record cannot be read
			if err != nil {
				return fmt.Errorf("failed to read first CSV record: %w", err)
			}

			value, err := strconv.ParseFloat(record[0], 64)

			if err != nil {
				fmt.Println(record)
				return fmt.Errorf("value failed to convert from string to float: %w", err)
			}

			//convert income to float64
			income, err := strconv.ParseFloat(record[1], 64)

			if err != nil {
				return fmt.Errorf("income failed to convert from string to float64: %w", err)
			}
			//convert age to int
			age, err := strconv.ParseInt(record[2], 10, 64)

			if err != nil {
				return fmt.Errorf("age failed to convert from string to int: %w", err)
			}
			//convert rooms to int
			rooms, err := strconv.ParseInt(record[3], 10, 64)

			if err != nil {
				return fmt.Errorf("rooms failed to convert from string to int: %w", err)
			}

			//convert bedrooms to int
			bedrooms, err := strconv.ParseInt(record[4], 10, 64)

			if err != nil {
				return fmt.Errorf(" bedrooms failed to convert from string to int: %w", err)
			}

			//convert pop to int
			pop, err := strconv.ParseInt(record[5], 10, 64)

			if err != nil {
				return fmt.Errorf("pop failed to convert from string to int: %w", err)
			}

			//convert hh to int
			hh, err := strconv.ParseInt(record[6], 10, 64)

			if err != nil {
				return fmt.Errorf("hh failed to convert from string to int: %w", err)
			}

			house := House{
				Value:    value,
				Income:   income,
				Age:      age,
				Rooms:    rooms,
				Bedrooms: bedrooms,
				Pop:      pop,
				Hh:       hh,
			}

			houses.Records = append(houses.Records, house)

		}

		//create JSON output file
		outputFile, err := os.Create(outputPath)

		//return your error if file does not create
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}

		defer outputFile.Close()

		//create encoder JSON file
		encoder := json.NewEncoder(outputFile)
		err = encoder.Encode(houses)

		//return error if jsonrecord cannot be encoded
		if err != nil {
			return fmt.Errorf("failed to encode file: %w", err)
		}

		fmt.Printf("Sucess! Created %s\n", outputPath)
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

func ValidateArgs(input, output string) error {
	inputExtension := filepath.Ext(input)
	if inputExtension != ".csv" {
		return ErrInvalidInputFile
	}

	outputExtension := filepath.Ext(output)
	if outputExtension != ".jl" && outputExtension != ".json" {
		return ErrInvalidOutputFile
	}
	return nil
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
