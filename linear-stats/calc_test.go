package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

// ANSI color codes
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

// Run a command and capture return its output
func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return stdout.String(), nil
}

func TestOutput(t *testing.T) {
	// Run the binary executable.
	// docker command: "./stat-bin-dockerized/stat-bin/run.sh", "linear-stats"
	binaryOutput, err := runCommand("./linear-stats")
	if err != nil {
		t.Fatalf(Red+"Error running binary executable: %v"+Reset, err)
	}

	// Run the Python program.
	pythonOutput, err := runCommand("go", "run", ".", "data.txt")
	if err != nil {
		t.Fatalf(Red+"Error running Python program: %v"+Reset, err)
	}

	// Compare the outputs
	if pythonOutput != binaryOutput {
		fmt.Printf(Yellow + "Outputs do not match!\n" + Reset)
		fmt.Printf(Red+"Python Output:\n%s\n"+Reset, pythonOutput)
		fmt.Printf(Red+"Binary Output:\n%s\n"+Reset, binaryOutput)
		t.Errorf(Red + "Test failed: outputs are different" + Reset)
	} else {
		fmt.Printf(Green + "Outputs match!\n" + Reset)
		fmt.Printf(Green+"Python Output: %s\n"+Reset, pythonOutput)
		fmt.Printf(Green+"Binary Output: %s\n"+Reset, binaryOutput)
	}
}
