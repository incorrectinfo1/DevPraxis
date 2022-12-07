package main

import (
	"fmt"
	"os/exec"
)

func runBashScript(script string) error {
	// Run the script using the bash command
	cmd := exec.Command("bash", "-c", script)

	// Get the output from the command
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	// Print the output to the console
	fmt.Println(string(output))

	return nil
}

func main() {
	// Define the script to be run
	script := "echo 'Hello, World!'"

	// Run the script
	err := runBashScript(script)
	if err != nil {
		fmt.Println(err)
	}
}
