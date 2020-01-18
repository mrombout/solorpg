package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"testing"

	"github.com/mrombout/solorpg/assert"
)

var validAskResponseRegexp = `(Yes|No).*(and|but)?\n`
var validAskResponse = regexp.MustCompile(validAskResponseRegexp)

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
			givenAskIsInstalled()
			actualOutput, actualError := whenCommandIsRun("ask", testCase.args)
			thenItExitsWithStatusCode(t, actualError, 0)
			thenItOutputsAValidAskResult(t, actualOutput)
		})
	}
}

func TestMain_HandlesInvalidArguments(t *testing.T) {
	testCases := map[string]struct {
		args               []string
		expectedOutput     string
		expectedExitStatus int
	}{
		"no +/- on modifier":    {args: []string{"1"}, expectedOutput: "modifier not in format <+|-><number>\n", expectedExitStatus: 1},
		"modifier not a number": {args: []string{"forty"}, expectedOutput: "modifier not a valid number\n", expectedExitStatus: 1},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			givenAskIsInstalled()
			actualOutput, actualError := whenCommandIsRun("ask", testCase.args)
			thenItExitsWithStatusCode(t, actualError, testCase.expectedExitStatus)
			thenItOutputsAnErrorMessage(t, actualOutput, testCase.expectedOutput)
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
	t.Helper()

	if expectedExitStatus == 0 {
		assert.Nil(t, actualError)
	} else {
		assert.EqualError(t, actualError, fmt.Sprintf("exit status %d", expectedExitStatus))
	}
}

func thenItOutputsAValidAskResult(t *testing.T, actualOutput bytes.Buffer) {
	t.Helper()

	actualOutputBytes := actualOutput.Bytes()
	if !validAskResponse.Match(actualOutputBytes) {
		t.Errorf("expected '%v' to match regular expression '%v', but it didn't", string(actualOutputBytes), validAskResponseRegexp)
	}
}

func thenItOutputsAnErrorMessage(t *testing.T, actualOutput bytes.Buffer, expectedOutput string) {
	t.Helper()

	assert.Equal(t, expectedOutput, actualOutput.String())
}
