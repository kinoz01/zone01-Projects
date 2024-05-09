package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Test 1",
			args: []string{"hello", "standard"},
		},
		{
			name: "Test 2",
			args: []string{"HELLO", "standard"},
		},
		{
			name: "Test 3",
			args: []string{"HeLlo HuMaN", "standard"},
		},
		{
			name: "Test 4",
			args: []string{"1Hello 2There", "standard"},
		},
		{
			name: "Test 5",
			args: []string{"Hello\\nThere", "standard"},
		},
		{
			name: "Test 6",
			args: []string{"Hello\\n\\nThere", "standard"},
		},
		{
			name: "Test 7",
			args: []string{"{Hello & There #}", "standard"},
		},
		{
			name: "Test 8",
			args: []string{"hello There 1 to 2!", "standard"},
		},
		{
			name: "Test 9",
			args: []string{"MaD3IrA&LiSboN", "standard"},
		},
		{
			name: "Test 10",
			args: []string{"1a\"#FdwHywR&/()=", "standard"},
		},
		{
			name: "Test 11",
			args: []string{"{|}~", "standard"},
		},
		{
			name: "Test 12",
			args: []string{`[\]^_ 'a`, "standard"},
		},
		{
			name: "Test 13",
			args: []string{"RGB", "standard"},
		},
		{
			name: "Test 16",
			args: []string{":;<=>?@", "standard"},
		},
		{
			name: "Test 17",
			args: []string{`\!" #$%&'()*+,-./`, "standard"},
		},
		{
			name: "Test 18",
			args: []string{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "standard"},
		},
		{
			name: "Test 19",
			args: []string{"abcdefghijklmnopqrstuvwxyz", "standard"},
		},
		{
			name: "Test 21",
			args: []string{"hello world", "shadow"},
		},
		{
			name: "Test 22",
			args: []string{"nice 2 meet you", "thinkertoy"},
		},
		{
			name: "Test 23",
			args: []string{"you & me", "standard"},
		},
		{
			name: "Test 24",
			args: []string{"123", "shadow"},
		},
		{
			name: "Test 25",
			args: []string{"/(\")", "thinkertoy"},
		},
		{
			name: "Test 26",
			args: []string{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "shadow"},
		},
		{
			name: "Test 27",
			args: []string{"\"#$%&/()*+,-./", "thinkertoy"},
		},
		{
			name: "Test 28",
			args: []string{"It's Working", "thinkertoy"},
		},

	}

	wantFiles := make([]string, len(testCases))
	for i := range testCases {
		wantFiles[i] = fmt.Sprintf("./test_cases/want%d.txt", i+1)
	}

	for i, tc := range testCases {
		
		tc.want = readWantFile(wantFiles[i])
		
		t.Run(tc.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			go func() {
				os.Args = []string{"main.go"}
				os.Args = append(os.Args, tc.args...)
				main()
				w.Close()
			}()

			var buf bytes.Buffer
			_, _ = buf.ReadFrom(r)
			r.Close()
			os.Stdout = old

			got := buf.String()

			if got != tc.want {
				t.Errorf("\n\nFor input \x1b[31m%s\x1b[0m\n\nExpected:\n\x1b[36m%s\x1b[0m\nBUT Got:\n%s", strings.Join(tc.args, " "), tc.want, got)
			}
		})
	}
}

func readWantFile(filename string) string {
	content, _ := os.ReadFile(filename)
	return string(content)
}
