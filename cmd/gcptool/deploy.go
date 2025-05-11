package main

import (
    "fmt"
    "os"
)

func deploy(functionName, environment, revision string, clean bool) {
    if functionName == "" {
        fmt.Println("Function name is required")
        os.Exit(1)
    }

    if clean {
        cmd := execCommand("go", "clean")
        cmd.Dir = "../../function"
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Run()
    }

    cmd := execCommand("gcloud", "functions", "deploy", functionName,
        "--runtime", "go121",
        "--trigger-http",
        "--allow-unauthenticated",
        "--region", "us-central1",
        "--entry-point", "Carbonquest",
    )

    cmd.Dir = "./function"
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    fmt.Printf("Deploying function %s [env: %s, rev: %s]...\n", functionName, environment, revision)
    err := cmd.Run()
    if err != nil {
        fmt.Println("Deployment failed:", err)
    }
}
