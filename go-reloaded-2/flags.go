package main

import (
	"regexp"
	"strconv"
	"strings"
)

var globalRe *regexp.Regexp

func initPattern() {
	globalRe = regexp.MustCompile(`\((up|low|cap|hex|bin),(\d+)\)`)
}

func Flags(text string) string {
	initPattern()

	// flags, globalisation:
	re1 := regexp.MustCompile(`\((up|low|cap|hex|bin)\)`)
	text = re1.ReplaceAllString(text, "($1,1)")
	re2 := regexp.MustCompile(`\((up|low|cap), (\d+)\)`)
	text = re2.ReplaceAllString(text, "($1,$2)")
	reSpace := regexp.MustCompile(`(\S)\((up|low|cap|hex|bin),(\d+)\)`)
	for reSpace.MatchString(text){
		text = reSpace.ReplaceAllString(text, "$1 ($2,$3)") // resolving the overlapping problem by lopping and replacing the first match each time
	}

	lines := strings.Split(text, "\n")

	var newLines []string

	for _, line := range lines {
		Words := strings.Fields(line)

		for i, word := range Words {
			if i != 0 && globalRe.MatchString(word) {
				submatches := globalRe.FindStringSubmatch(word)
				nb, _ := strconv.Atoi(submatches[2])

				j := i - 1

				for nb > 0 && j >= 0 {
					switch submatches[1] {
					case "hex":
						if ValidHex(Words[j]) {
							Words[j] = convertHex(Words[j])
							nb--
						}
					case "bin":
						if ValidBin(Words[j]) {
							Words[j] = convertBin(Words[j])
							nb--
						}
					case "up":
						if ValidWord(Words[j]) {
							Words[j] = strings.ToUpper(Words[j])
							nb--
						}
					case "cap":
						if ValidWord(Words[j]) {
							Words[j] = Title(Words[j])
							nb--
						}
					case "low":
						if ValidWord(Words[j]) {
							Words[j] = strings.ToLower(Words[j])
							nb--
						}
					}
					j--
				}
			}
		}

		line = strings.Join(Words, " ")
		line = strings.TrimSpace(line)
		newLines = append(newLines, line)
	}


	text = strings.Join(newLines, "\n")

	text = globalRe.ReplaceAllString(text, "") // remove flags

	reClean := regexp.MustCompile(` +`)
	text = reClean.ReplaceAllString(text, " ") // remove trailing spaces
	reClean = regexp.MustCompile(` +\n`)
	text = reClean.ReplaceAllString(text, "\n") // remove trailing spaces
	reClean = regexp.MustCompile(` +$`)
	text = reClean.ReplaceAllString(text, "") // remove trailing spaces

	return text
}
