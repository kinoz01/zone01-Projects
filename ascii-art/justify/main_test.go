package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name      string
		args      []string
		want      string
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
			name: "Test 14",
			args: []string{":;<=>?@", "standard"},
		},
		{
			name: "Test 15",
			args: []string{`\!" #$%&'()*+,-./`, "standard"},
		},
		{
			name: "Test 16",
			args: []string{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "standard"},
		},
		{
			name: "Test 17",
			args: []string{"abcdefghijklmnopqrstuvwxyz", "standard"},
		},
		{
			name: "Test 18",
			args: []string{"hello world", "shadow"},
		},
		{
			name: "Test 19",
			args: []string{"nice 2 meet you", "thinkertoy"},
		},
		{
			name: "Test 20",
			args: []string{"you & me", "standard"},
		},
		{
			name: "Test 21",
			args: []string{"123", "shadow"},
		},
		{
			name: "Test 22",
			args: []string{"/(\")", "thinkertoy"},
		},
		{
			name: "Test 23",
			args: []string{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "shadow"},
		},
		{
			name: "Test 24",
			args: []string{"\"#$%&/()*+,-./", "thinkertoy"},
		},
		{
			name: "Test 25",
			args: []string{"It's Working", "thinkertoy"},
		},
		{
			name: "Test 26 (All standard)",  // go run . ' !"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~'
			args: []string{` !"#$%&'()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`+"`"+`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~`},
		},
		{
			name: "Test 27",
			args: []string{"--align", "right", "something", "standard"},  
		},
		{
			name: "Test 28",
			args: []string{"--output=test00.txt", "First\\nTest", "shadow"},
		},
		{
			name: "Test 29",
			args: []string{"--output=test01.txt", "hello", "standard"},
		},
		{
			name: "Test 30",
			args: []string{"--output=test02.txt", "123 -> #$%", "standard"},
		},
		{
			name: "Test 31",
			args: []string{"--output=test03.txt", "432 -> #$%&@", "shadow"},
		},
		{
			name: "Test 32",
			args: []string{"--output=test04.txt", "There", "shadow"},
		},
		{
			name: "Test 33",
			args: []string{"--output=test05.txt", "123 -> \"#$%@", "thinkertoy"},
		},
		{
			name: "Test 34",
			args: []string{"--output=test06.txt", "2 you", "thinkertoy"},
		},
		{
			name: "Test 35",
			args: []string{"--output=test07.txt", "Testing long output!", "standard"},
		},
		{
			name: "Test 36 (All shadow)",  // go run . ' !"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~'
			args: []string{` !"#$%&'()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`+"`"+`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~`, "shadow"},
		},
		{
			name: "Test 37 (All blocks)",  // go run . ' !"#$%&'\''()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~'
			args: []string{` !"#$%&'()\n*+,-./012345\n6789:;<=>?@AB\nCDEFGHIJK\nLMNOPQRSTUVW\nXYZ[\]^_`+"`"+`abc\ndefghijk\nlmnopqrst\nuvwxyz{|}~`, "blocks"},
		},
	}

	wantFiles := make([]string, len(testCases))
	for i := range testCases {
		wantFiles[i] = fmt.Sprintf("./test_cases/want%d.txt", i+1)
	}

	reOutput := regexp.MustCompile(`\A--output=(\S+.txt)$`)

	for i, tc := range testCases {
		var err error
		got := ""
		tc.want, err = readWantFile(wantFiles[i])
		if err != nil{
			fmt.Println(err)
			continue
		}
		t.Run(tc.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"main.go"}
			os.Args = append(os.Args, tc.args...)
			main()
			w.Close()

			var buf bytes.Buffer
			_, _ = buf.ReadFrom(r)

			r.Close()

			if reOutput.MatchString(tc.args[0]) {
				gotByte, err := os.ReadFile(reOutput.FindStringSubmatch(tc.args[0])[1])
				if err != nil {
					t.Errorf("\n\nCan't find output file for input \x1b[31m%s\x1b[0m.\n\n", strings.Join(tc.args, " "))
				}
				got = string(gotByte)
			} else {
				got = buf.String()
			}

			if got != tc.want {
				t.Errorf("\n\nFor input \x1b[31m%s\x1b[0m\n\nExpected:\n\x1b[36m%s\x1b[0m\nBUT Got:\n%s", strings.Join(tc.args, " "), tc.want, got)
			}
		})
	}
}

func readWantFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
