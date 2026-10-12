package cmd_test

import (
	"errors"
	"testing"

	"github.com/gummybeardream/command-line-application/cmd"
)

func TestValidateArgs(t *testing.T) {
	t.Run("returns error when input is not a csv", func(t *testing.T) {
		result := cmd.ValidateArgs("test.ccc", "test.jl")

		if !errors.Is(result, cmd.ErrInvalidInputFile) {
			t.Error("accepted a non-CSV file")
		}
	})

	t.Run("returns error when output is not .jl or .json", func(t *testing.T) {
		result := cmd.ValidateArgs("test.csv", "test.jjj")

		if !errors.Is(result, cmd.ErrInvalidOutputFile) {
			t.Error("accepted a non-json or non-jl file")
		}
	})

	t.Run("returns nil when input is a csv and output is .jl", func(t *testing.T) {
		result := cmd.ValidateArgs("test.csv", "test.jl")

		if result != nil {
			t.Error("received unexpected error")
		}
	})

	t.Run("returns nil when input is a csv and output is .json", func(t *testing.T) {
		result := cmd.ValidateArgs("test.csv", "test.json")

		if result != nil {
			t.Error("received unexpected error")
		}
	})
}
