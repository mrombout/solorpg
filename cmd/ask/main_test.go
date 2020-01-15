package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/mrombout/solorpg/assert"
)

func TestMain_GeneratesResponse(t *testing.T) {
	testCases := map[string]struct {
		args []string
	}{
		"no args":           {args: []string{}},
		"positive modifier": {args: []string{"+1"}},
		"negative modifier": {args: []string{"-1"}},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			// Arrange
			var output bytes.Buffer

			askCommand := exec.Command("ask", testCase.args...)
			askCommand.Stdout = &output

			// Act
			err := askCommand.Run()

			// Assert
			assert.Nil(t, err)
			assert.NotEmpty(t, output.String())
		})
	}
}

func TestMain_HandlesInvalidArguments(t *testing.T) {
	testCases := map[string]struct {
		args               []string
		expectedOutput     string
		expectedExitStatus int
	}{
		"no +/- on modifier":    {args: []string{"1"}, expectedOutput: "modifier not in format <+|-><number>", expectedExitStatus: 1},
		"modifier not a number": {args: []string{"forty"}, expectedOutput: "modifier not a valid number", expectedExitStatus: 1},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			// Arrange
			var output bytes.Buffer

			askCommand := exec.Command("ask", testCase.args...)
			askCommand.Stdout = &output

			// Act
			err := askCommand.Run()

			// Assert
			assert.EqualError(t, err, fmt.Sprintf("exit status %d", testCase.expectedExitStatus))
			assert.Equal(t, testCase.expectedOutput, strings.Trim(output.String(), "\n"))
		})
	}
}
