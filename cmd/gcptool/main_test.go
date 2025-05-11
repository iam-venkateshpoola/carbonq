package main

import (
	"os"
	"os/exec"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


var commandRunner = exec.Command // Used for mocking exec.Command in tests

// Helper function to run the CLI in tests
func runCLI(args []string) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	os.Args = append([]string{"gcptool"}, args...) // Set os.Args for the CLI
	main()
}

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

// Test deploy command
func TestDeploy(t *testing.T) {
	// Set up a mock for exec.Command
	mockExec := new(MockExec)
	commandRunner = func(name string, args ...string) *exec.Cmd {
		mockExec.On("Run").Return(nil) // Simulate successful run
		cmd := exec.Command(name, args...)
		cmd.Run = mockExec.Run
		return cmd
	}
	defer func() { commandRunner = exec.Command }() // Reset after the test

	// Run the deploy CLI test
	runCLI([]string{"deploy", "-e", "prod", "-v", "2", "-c", "true", "myFunction"})

	// Assert that exec.Command was called
	mockExec.AssertExpectations(t)
}

// Test describe command
func TestDescribe(t *testing.T) {
	// Set up a mock for exec.Command
	mockExec := new(MockExec)
	commandRunner = func(name string, args ...string) *exec.Cmd {
		mockExec.On("CombinedOutput").Return([]byte("Function details"), nil) // Simulate output
		cmd := exec.Command(name, args...)
		cmd.CombinedOutput = mockExec.CombinedOutput
		return cmd
	}
	defer func() { commandRunner = exec.Command }() // Reset after the test

	// Run the describe CLI test
	runCLI([]string{"describe", "myFunction"})

	// Assert that exec.Command was called
	mockExec.AssertExpectations(t)
}

// Test deploy function with mock exec
func TestDeployFunction(t *testing.T) {
	// Set up a mock for exec.Command
	mockExec := new(MockExec)
	commandRunner = func(name string, args ...string) *exec.Cmd {
		mockExec.On("Run").Return(nil)
		cmd := exec.Command(name, args...)
		cmd.Run = mockExec.Run
		return cmd
	}
	defer func() { commandRunner = exec.Command }()

	// Run the deploy function with mock arguments
	deploy("myFunction", "dev", "1", false)

	// Assert that exec.Command was called
	mockExec.AssertExpectations(t)
}

// Test describe function with mock exec
func TestDescribeFunction(t *testing.T) {
	// Set up a mock for exec.Command
	mockExec := new(MockExec)
	commandRunner = func(name string, args ...string) *exec.Cmd {
		mockExec.On("CombinedOutput").Return([]byte("Function details"), nil)
		cmd := exec.Command(name, args...)
		cmd.CombinedOutput = mockExec.CombinedOutput
		return cmd
	}
	defer func() { commandRunner = exec.Command }()

	// Run the describe function with mock arguments
	describe("myFunction")

	// Assert that exec.Command was called
	mockExec.AssertExpectations(t)
}
