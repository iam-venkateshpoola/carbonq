package main

import "os/exec"

// Override in tests
var execCommand = exec.Command
