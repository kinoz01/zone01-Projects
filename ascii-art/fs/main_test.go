package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {

	want1, _ := os.ReadFile("./test_cases/want1.txt")
	want2, _ := os.ReadFile("./test_cases/want2.txt")
	want3, _ := os.ReadFile("./test_cases/want3.txt")
	want4, _ := os.ReadFile("./test_cases/want4.txt")
	want5, _ := os.ReadFile("./test_cases/want5.txt")
	want6, _ := os.ReadFile("./test_cases/want6.txt")
	want7, _ := os.ReadFile("./test_cases/want7.txt")
	want8, _ := os.ReadFile("./test_cases/want8.txt")
	want9, _ := os.ReadFile("./test_cases/want9.txt")
	want10, _ := os.ReadFile("./test_cases/want10.txt")
	want11, _ := os.ReadFile("./test_cases/want11.txt")
	want12, _ := os.ReadFile("./test_cases/want12.txt")
	want13, _ := os.ReadFile("./test_cases/want13.txt")
	want14, _ := os.ReadFile("./test_cases/want14.txt")
	want15, _ := os.ReadFile("./test_cases/want15.txt")
	want16, _ := os.ReadFile("./test_cases/want16.txt")
	want17, _ := os.ReadFile("./test_cases/want17.txt")
	want18, _ := os.ReadFile("./test_cases/want18.txt")
	want19, _ := os.ReadFile("./test_cases/want19.txt")
	want20, _ := os.ReadFile("./test_cases/want20.txt")
	want21, _ := os.ReadFile("./test_cases/want21.txt")
	want22, _ := os.ReadFile("./test_cases/want22.txt")
	want23, _ := os.ReadFile("./test_cases/want23.txt")
	want24, _ := os.ReadFile("./test_cases/want24.txt")
	want25, _ := os.ReadFile("./test_cases/want25.txt")

	// Define test cases
	testCases := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Test 1",
			args: []string{"\\n", "standard"},
			want: "\n",
		},
		{
			name: "Test 2",
			args: []string{`\n\n`, "standard"},
			want: "\n\n",
		},
		{
			name: "Test 3",
			args: []string{"hello", "standard"},
			want: string(want1),
		},
		{
			name: "Test 4",
			args: []string{"HELLO", "standard"},
			want: string(want2),
		},
		{
			name: "Test 5",
			args: []string{"HeLlo HuMaN", "standard"},
			want: string(want3),
		},
		{
			name: "Test 6",
			args: []string{"1Hello 2There", "standard"},
			want: string(want4),
		},
		{
			name: "Test 7",
			args: []string{"Hello\\nThere", "standard"},
			want: string(want5),
		},
		{
			name: "Test 8",
			args: []string{"Hello\\n\\nThere", "standard"},
			want: string(want6),
		},
		{
			name: "Test 9",
			args: []string{"{Hello & There #}", "standard"},
			want: string(want7),
		},
		{
			name: "Test 10",
			args: []string{"hello There 1 to 2!", "standard"},
			want: string(want8),
		},
		{
			name: "Test 11",
			args: []string{"MaD3IrA&LiSboN", "standard"},
			want: string(want9),
		},
		{
			name: "Test 12",
			args: []string{"1a\"#FdwHywR&/()=", "standard"},
			want: string(want10),
		},
		{
			name: "Test 13",
			args: []string{"{|}~", "standard"},
			want: string(want11),
		},
		{
			name: "Test 14",
			args: []string{`[\]^_ 'a`, "standard"},
			want: string(want12),
		},
		{
			name: "Test 15",
			args: []string{"RGB", "standard"},
			want: string(want13),
		},
		{
			name: "Test 16",
			args: []string{":;<=>?@", "standard"},
			want: string(want14),
		},
		{
			name: "Test 17",
			args: []string{`\!" #$%&'()*+,-./`, "standard"},
			want: string(want15),
		},
		{
			name: "Test 18",
			args: []string{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "standard"},
			want: string(want16),
		},
		{
			name: "Test 19",
			args: []string{"abcdefghijklmnopqrstuvwxyz", "standard"},
			want: string(want17),
		},
		{
			name: "Test 20",
			args: []string{"²", "standard"},
			want: "🚨 Found an Invalid Ascii Character.\n",
		},
		{
			name: "Test 21",
			args: []string{"hello world", "shadow"},
			want: string(want18),
		},
		{
			name: "Test 22",
			args: []string{"nice 2 meet you", "thinkertoy"},
			want: string(want19),
		},
		{
			name: "Test 23",
			args: []string{"you & me", "standard"},
			want: string(want20),
		},
		{
			name: "Test 24",
			args: []string{"123", "shadow"},
			want: string(want21),
		},
		{
			name: "Test 25",
			args: []string{"/(\")", "thinkertoy"},
			want: string(want22),
		},
		{
			name: "Test 26",
			args: []string{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "shadow"},
			want: string(want23),
		},
		{
			name: "Test 27",
			args: []string{"\"#$%&/()*+,-./", "thinkertoy"},
			want: string(want24),
		},
		{
			name: "Test 28",
			args: []string{"It's Working", "thinkertoy"},
			want: string(want25),
		},
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
