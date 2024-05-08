package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {

	want1, _ := os.ReadFile("./test/cases/want1.txt")

	// Define test cases
	testCases := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Test 1",
			args: []string{"hello", "standard"},
			want: string(want1),
		},
		{
			name: "Test with empty arguments",
			args: []string{},
			want: "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n",
		},
		{
			name: "Test with empty arguments",
			args: []string{},
			want: "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n",
		},
		// Add more test cases as needed
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Redirect stdout to capture output
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the main function with simulated arguments
			go func() {
				os.Args = []string{"main.go"}
				os.Args = append(os.Args, tc.args...)
				main()
				w.Close()
			}()

			// Read output from stdout
			var buf bytes.Buffer
			_, _ = buf.ReadFrom(r)
			r.Close()
			os.Stdout = old

			// Check the output
			got := buf.String()

			if got != tc.want {
				t.Errorf("\n\nFor input \x1b[31m%s\x1b[0m\n\nExpected:\n\x1b[36m%s\x1b[0m\nBUT Got:\n%s", strings.Join(tc.args, " "), tc.want, got)
			}
		})
	}
}
