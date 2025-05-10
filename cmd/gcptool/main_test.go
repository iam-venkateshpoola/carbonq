package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
    "os/exec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


// MockExec is a mock type for exec.Command
type MockExec struct {
	mock.Mock
}

func (m *MockExec) Run() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockExec) CombinedOutput() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

// Mocked command for exec.Command
var execCommand = exec.Command

// Reset execCommand for testing
func resetExecCommand() {
	execCommand = exec.Command
}

// SetExecCommand allows you to replace execCommand during testing
func SetExecCommand(command func(string, ...string) *exec.Cmd) {
	execCommand = command
}

// Test main function for "deploy" command
func TestMainDeploy(t *testing.T) {
	// Set up a mock for exec.Command
	mockExec := new(MockExec)
	SetExecCommand(func(name string, arg ...string) *exec.Cmd {
		mockExec.On("Run").Return(nil)
		cmd := exec.Command(name, arg...)
		cmd.Run = mockExec.Run
		return cmd
	})
	defer resetExecCommand()

	// Simulate passing the "deploy" subcommand with arguments
	os.Args = []string{"cmd", "deploy", "-e", "prod", "-v", "2", "-c", "true", "myFunction"}
	main()

	// Assert the "deploy" command was run
	mockExec.AssertExpectations(t)
}

// Test main function for "describe" command
func TestMainDescribe(t *testing.T) {
	// Set up a mock for exec.Command
	mockExec := new(MockExec)
	SetExecCommand(func(name string, arg ...string) *exec.Cmd {
		mockExec.On("CombinedOutput").Return([]byte("Function details"), nil)
		cmd := exec.Command(name, arg...)
		cmd.CombinedOutput = mockExec.CombinedOutput
		return cmd
	})
	defer resetExecCommand()

	// Simulate passing the "describe" subcommand with an argument
	os.Args = []string{"cmd", "describe", "myFunction"}
	main()

	// Assert the "describe" command was run
	mockExec.AssertExpectations(t)
}

// Test deploy function with mock exec
func TestDeploy(t *testing.T) {
	// Set up a mock for exec.Command
	mockExec := new(MockExec)
	SetExecCommand(func(name string, arg ...string) *exec.Cmd {
		mockExec.On("Run").Return(nil)
		cmd := exec.Command(name, arg...)
		cmd.Run = mockExec.Run
		return cmd
	})
	defer resetExecCommand()

	// Run the deploy function with mock arguments
	deploy("myFunction", "dev", "1", false)

	// Assert that exec.Command was called
	mockExec.AssertExpectations(t)
}

// Test describe function with mock exec
func TestDescribe(t *testing.T) {
	// Set up a mock for exec.Command
	mockExec := new(MockExec)
	SetExecCommand(func(name string, arg ...string) *exec.Cmd {
		mockExec.On("CombinedOutput").Return([]byte("Function details"), nil)
		cmd := exec.Command(name, arg...)
		cmd.CombinedOutput = mockExec.CombinedOutput
		return cmd
	})
	defer resetExecCommand()

	// Run the describe function with mock arguments
	describe("myFunction")

	// Assert that exec.Command was called
	mockExec.AssertExpectations(t)
}
