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
			givenAskIsInstalled()
			actualOutput, actualError := whenCommandIsRun("ask", testCase.args)
			thenItExitsWithStatusCode(t, actualError, testCase.expectedExitStatus)
			thenItOutputsAValidAskResult(t, actualOutput, testCase.expectedOutput)
		})
	}
}

func givenAskIsInstalled() {
	// TODO: Install ask command
}

func whenCommandIsRun(command string, args []string) (bytes.Buffer, error) {
	var output bytes.Buffer

	askCommand := exec.Command(command, args...)
	askCommand.Stdout = &output

	err := askCommand.Run()

	return output, err
}

func thenItExitsWithStatusCode(t *testing.T, actualError error, expectedExitStatus int) {
	assert.EqualError(t, actualError, fmt.Sprintf("exit status %d", expectedExitStatus))
}

func thenItOutputsAValidAskResult(t *testing.T, actualOutput bytes.Buffer, expectedOutput string) {
	// TODO: Actually verify if output is valid
	assert.Equal(t, expectedOutput, strings.Trim(actualOutput.String(), "\n"))
}
