package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploy_Success(t *testing.T) {
	called := false
	execCommand = func(name string, args ...string) *exec.Cmd {
		called = true
		return fakeExecCommand(name, args...)
	}
	defer func() { execCommand = exec.Command }()

	// Set up minimal required env for the function
	os.MkdirAll("./function", 0755)
	defer os.RemoveAll("./function")

	deploy("test-func", "prod", "1", false)
	assert.True(t, called, "execCommand should be called")
}

func TestDeploy_MissingName(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		deploy("", "dev", "1", false)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestDeploy_MissingName")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	assert.NotNil(t, err, "deploy with empty function name should exit")
}

func TestDescribe_Success(t *testing.T) {
	called := false
	execCommand = func(name string, args ...string) *exec.Cmd {
		called = true
		return fakeExecCommand(name, args...)
	}
	defer func() { execCommand = exec.Command }()

	describe("test-func")
	assert.True(t, called, "execCommand should be called")
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	os.Exit(0)
}
