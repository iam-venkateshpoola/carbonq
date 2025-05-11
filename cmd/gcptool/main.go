package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'deploy' or 'describe' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "deploy":
		deployCmd := flag.NewFlagSet("deploy", flag.ExitOnError)
		env := deployCmd.String("e", "dev", "Environment")
		rev := deployCmd.String("v", "1", "Revision")
		clean := deployCmd.Bool("c", false, "Clean and rebuild before deploying")
		deployCmd.Parse(os.Args[2:])
		deploy(deployCmd.Arg(0), *env, *rev, *clean)
	case "describe":
		describeCmd := flag.NewFlagSet("describe", flag.ExitOnError)
		describeCmd.Parse(os.Args[2:])
		describe(describeCmd.Arg(0))
	default:
		fmt.Println("Expected 'deploy' or 'describe' subcommands")
		os.Exit(1)
	}
}
