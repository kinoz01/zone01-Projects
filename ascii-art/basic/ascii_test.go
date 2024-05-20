package main

import (
	"bytes"
	"os"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	// Define test cases
	want1, _ := os.ReadFile("./cases/want1.txt")
	want2, _ := os.ReadFile("./cases/want2.txt")
	want3, _ := os.ReadFile("./cases/want3.txt")
	want4, _ := os.ReadFile("./cases/want4.txt")
	want5, _ := os.ReadFile("./cases/want5.txt")
	want6, _ := os.ReadFile("./cases/want6.txt")
	want7, _ := os.ReadFile("./cases/want7.txt")
	want8, _ := os.ReadFile("./cases/want8.txt")
	want9, _ := os.ReadFile("./cases/want9.txt")
	want10, _ := os.ReadFile("./cases/want10.txt")
	want11, _ := os.ReadFile("./cases/want11.txt")
	want12, _ := os.ReadFile("./cases/want12.txt")
	want13, _ := os.ReadFile("./cases/want13.txt")
	want14, _ := os.ReadFile("./cases/want14.txt")
	want15, _ := os.ReadFile("./cases/want15.txt")
	want16, _ := os.ReadFile("./cases/want16.txt")
	want17, _ := os.ReadFile("./cases/want17.txt")

	tests := []struct {
		name string // Name of the test case
		arg  string // Input text
		want string // Expected output using a regexp pattern
	}{
		{"Test 1", `\n`, "\n"},
		{"Test 2", `\n\n`, "\n\n"},
		{"Test 3", "hello", string(want1)},
		{"Test 4", "HELLO", string(want2)},
		{"Test 5", "HeLlo HuMaN", string(want3)},
		{"Test 6", "1Hello 2There", string(want4)},
		{"Test 7", `Hello\nThere`, string(want5)},
		{"Test 8", `Hello\n\nThere`, string(want6)},
		{"Test 9", `{Hello & There #}`, string(want7)},
		{"Test 10", `hello There 1 to 2!`, string(want8)},
		{"Test 11", `MaD3IrA&LiSboN`, string(want9)},
		{"Test 12", "1a\"#FdwHywR&/()=", string(want10)},
		{"Test 13", "{|}~", string(want11)},
		{"Test 14", `[\]^_ 'a`, string(want12)},
		{"Test 15", "RGB", string(want13)},
		{"Test 16", ":;<=>?@", string(want14)},
		{"Test 17", `\!" #$%&` + "'" + `()*+,-./`, string(want15)},
		{"Test 18", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(want16)},
		{"Test 19", "abcdefghijklmnopqrstuvwxyz", string(want17)},
	}

	for _, cas := range tests {
		got := ""
		t.Run(cas.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"main.go"}
			os.Args = append(os.Args, cas.arg)
			main()
			w.Close()

			var buf bytes.Buffer
			_, _ = buf.ReadFrom(r)

			r.Close()
			got = buf.String()

			if got != cas.want {
				t.Errorf("\n\nFor input \x1b[31m%s\x1b[0m\n\nExpected:\n\x1b[36m%s\x1b[0m\nBUT Got:\n%s", cas.arg, cas.want, got)
			}
		})
	}
}
