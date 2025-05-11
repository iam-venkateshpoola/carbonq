package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCLI_Deploy(t *testing.T) {
	called := false
	execCommand = func(name string, args ...string) *exec.Cmd {
		called = true
		assert.Equal(t, "gcloud", name)
		return fakeExecCommand()
	}
	defer func() { execCommand = exec.Command }()

	runCLI([]string{"deploy", "-e", "dev", "-v", "1", "myFunction"})
	assert.True(t, called, "deploy command was not triggered")
}

func TestRunCLI_Describe(t *testing.T) {
	called := false
	execCommand = func(name string, args ...string) *exec.Cmd {
		called = true
		assert.Equal(t, "gcloud", name)
		return fakeExecCommandWithOutput("Mocked output")
	}
	defer func() { execCommand = exec.Command }()

	runCLI([]string{"describe", "myFunction"})
	assert.True(t, called, "describe command was not triggered")
}

// Helper mocks
func fakeExecCommand() *exec.Cmd {
	return exec.Command("echo", "mock")
}

func fakeExecCommandWithOutput(output string) *exec.Cmd {
	return exec.Command("echo", output)
}
