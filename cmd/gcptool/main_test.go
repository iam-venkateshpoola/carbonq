package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

// execCommand is a variable to allow mocking exec.Command in tests
var execCommand = exec.Command

// Test deploy function without clean
func TestDeploy(t *testing.T) {
	called := false

	// Mock exec.Command
	execCommand = func(name string, args ...string) *exec.Cmd {
		called = true
		assert.Equal(t, "gcloud", name)
		return fakeExecCommand()
	}
	defer func() { execCommand = exec.Command }()

	deploy("myFunction", "dev", "1", false)
	assert.True(t, called, "execCommand was not called")
}

// Test describe function
func TestDescribe(t *testing.T) {
	called := false

	// Mock exec.Command
	execCommand = func(name string, args ...string) *exec.Cmd {
		called = true
		assert.Equal(t, "gcloud", name)
		return fakeExecCommandWithOutput("Function details\n")
	}
	defer func() { execCommand = exec.Command }()

	describe("myFunction")
	assert.True(t, called, "execCommand was not called")
}

// --- Helpers ---

// fakeExecCommand returns a dummy *exec.Cmd that runs `echo "mock"`
func fakeExecCommand() *exec.Cmd {
	return exec.Command("echo", "mock")
}

// fakeExecCommandWithOutput returns a *exec.Cmd that outputs mock data
func fakeExecCommandWithOutput(output string) *exec.Cmd {
	return exec.Command("echo", output)
}
