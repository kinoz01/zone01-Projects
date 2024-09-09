package format

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFlags(t *testing.T) {
	// Define test cases
	var tests = []struct {
		name string // Name of the test case
		text string // Input text
		want string // Expected output using a regexp pattern
	}{
		{"Test Flag", "LOW (low, 10(bin))", `low`},
	}

	for _, cas := range tests {
		t.Run(cas.name, func(t *testing.T) {
			// Compile the expected regexp.
			want := regexp.MustCompile(cas.want)

			// Call the function.
			msg := Flags(cas.text)

			// Check if the output matches the expected regexp.
			if !want.MatchString(msg) {
				t.Fatalf(`%q, want match for %#q, nil`, msg, want)
			}
		})
	}
}

func TestPunctuation(t *testing.T) {
	// Define test cases
	var tests = []struct {
		name string // Name of the test case
		text string // Input text
		want string // Expected output using a regexp pattern
	}{
		{"Test 1", "iwent .out  .,,!!! awesome", `iwent. out.,,!!! awesome`},
	}

	for _, cas := range tests {
		t.Run(cas.name, func(t *testing.T) {
			// Compile the expected regexp.
			// Call the function.
			got := Punctuation(cas.text)
			// Check if the output matches the expected regexp.
			if got != cas.want {
				t.Errorf("For input %q, expected match for %#q but got %q", cas.text, cas.want, got)
			}
		})
	}
}

// Test all the functions running in the expected order.
func TestFormat(t *testing.T) {
	// Define test cases
	var tests = []struct {
		name string // Name of the test case
		text string // Input text
		want string // Expected output using a regexp pattern
	}{
		{"Standard Test 1", "1E (hex) files were added", "30 files were added"},
		{"Standard Test 2", "It has been 10 (bin) years", "It has been 2 years"},
		{"Standard Test 3", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"Standard Test 4", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"Standard Test 5", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"Standard Test 6", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"Standard Test 7", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"Standard Test 8", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"Standard Test 9", "'I am the most well-known homosexual in the world'", "'I am the most well-known homosexual in the world'"},
		{"Standard Test 10", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"Standard Test 11", "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.", "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."},
		{"Standard Test 12", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"Standard Test 13", "There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"Standard Test 14", "Punctuation tests are ... kinda boring ,don't you think !?", "Punctuation tests are... kinda boring, don't you think!?"},
		{"Standard Test 15", "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?", "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"Standard Test 16", "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure", "I have to pack 5 outfits. Packed 26 just to be sure"},
		{"Standard Test 17", "Don not be sad ,because sad backwards is das . And das not good", "Don not be sad, because sad backwards is das. And das not good"},
		{"Standard Test 18", "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
	}

	for _, cas := range tests {
		t.Run(cas.name, func(t *testing.T) {
			// Compile the expected regexp.
			want := regexp.MustCompile(cas.want)

			// Call the function.
			got := Format(cas.text)

			// Check if the output matches the expected regexp.
			if !want.MatchString(got) {
				t.Errorf("For input %q, expected match for %#q but got %q", cas.text, cas.want, got)
			}
		})
	}
}

// Format run all function in the main.go file
func Format(text string) string {
	text = FlagsWrongUsage(text)
	text = Flags(text)
	text = Apostrophe(text)
	text = BasicGrammar(text)
	text = Punctuation(text)
	text = CleanText(text)
	return text
}



// func TestMain(t *testing.T) {
// 	testcases := []struct {
// 		res         string
// 		expectedRes string
// 	}{
// 		{"tests/test1.txt", "tests/result1.txt"},
// 		{"tests/testStandard.txt", "tests/resultStandard.txt"},
// 	}
// 	for _, tc := range testcases {
// 		res_file, err := os.ReadFile(tc.res)
// 		if err != nil {
// 			fmt.Println("Error: ", err)
// 		}
// 		resStr := string(res_file)
// 		expected_res, err := os.ReadFile(tc.expectedRes)
// 		if err != nil {
// 			fmt.Println("Error: ", err)
// 		}
// 		expectedStr := string(expected_res)
// 		if resStr != expectedStr {
// 			t.Errorf("For result file %s:\nExpected\n%s\nbut got\n%s", tc.res, expectedStr, resStr)
// 		}
// 	}
// }

