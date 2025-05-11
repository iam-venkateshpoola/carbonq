package main

import (
    "fmt"
)

func describe(functionName string) {
    if functionName == "" {
        fmt.Println("Function name is required")
        return
    }

    cmd := execCommand("gcloud", "functions", "describe", functionName, "--region", "us-central1")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Failed to describe function:", err)
    }
    fmt.Println(string(output))
}
