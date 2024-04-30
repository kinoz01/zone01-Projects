Regular expressions (regex) in Go (Golang) are handled through the regexp package. This package provides powerful functions to search, match, and manipulate text based on patterns. Here's a step-by-step breakdown of how to use regex in Go:

## 1. Import the `regexp` Package

To use regex, you first need to import the regexp package:

```go
import "regexp"
```

## 2. Compile a Regex Pattern

Before you can use a regex, you need to compile it into a `Regexp` object using `regexp.Compile` or `regexp.MustCompile`.  
`Compile` returns a regex object and an error if the pattern is not valid, while `MustCompile` panics if the pattern is not valid (which is useful for static patterns).

```go
re, err := regexp.Compile("pattern")
if err != nil {
    // handle error
}
```

Or for a static pattern:

```go
re := regexp.MustCompile("pattern")
```

## 3. Using Regex Methods

Once you have a `Regexp` object, you can use its **methods** to perform various operations:

### Matching

`MatchString`: Checks if a string matches the pattern (return true or false).

```go
matched := re.MatchString("string to check")
```

### Searching

`FindString`: Returns the first match of the pattern in the string.
```go
match := re.FindString("search in this string")
```

`FindStringIndex`: Returns the start and end index of the first match.
```go
indexes := re.FindStringIndex("search in this string")
```

`FindAllString`: Returns all non-overlapping matches.
```go
allMatches := re.FindAllString("search in this string", -1)
```
Signature:
```go
func (re *Regexp) FindAllString(s string, n int) []string
```

When you pass `-1` as the second argument, it tells the method to find all possible matches in the string. This is commonly used when you want to retrieve every match without any restriction on the count.

If you specify a non-negative integer `n`, then the method returns at most `n` matches. This can be useful if you're only interested in the first few matches and want to limit the output to improve performance or reduce output size.


### Extracting Submatches

`FindStringSubmatch`: Returns slices containing matched parts of the string including capturing groups (**return a slice of string** `[]string`).

```go
submatches := re.FindStringSubmatch("search in this string")
```

✋ `FindAllStringSubmatch`: The `FindAllStringSubmatch` method is designed to find all matches of the pattern in the provided text and also to capture the contents of any groups defined within the pattern (return  slice of slices of strings, specifically `[][]string`). Here's a breakdown of the returned data structure:

- **Outer Slice**: Each element of this outer slice represents one match of the entire pattern within the text. If your pattern is found three times in the text, the outer slice will have three elements.
- **Inner Slice**: Each element of this inner slice corresponds to one group of captured substrings for a single match. - The first element (index 0) is always the entire match, and subsequent elements (index 1, 2, ...) contain the substrings captured by any capturing groups in your regex pattern.

Signature:
```go
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
```

### Replacing

`ReplaceAllString`: Replaces all matches with a replacement string.
```go
result := re.ReplaceAllString("string to modify", "replacement")
```

`ReplaceAllStringFunc`: takes a source string (text) and a function. The function is called for each substring that matches the regular expression re. The function must return a string, which will replace the matched substring in the original text.
```go
return re.ReplaceAllStringFunc(text, func(match string) string {
    // function body
})
```

## 4. Handling Special Characters and Flags

In Go, a pattern in the context of regular expressions is indeed a string that describes the kind of text the regular expression is meant to match. Defining a pattern involves using a combination of regular characters and special characters that have specific meanings within the regex syntax. Here’s a rundown of how to define patterns and the special syntax used:

### Literal Characters

Most characters, like a, 1, or B, are "literals" and match the same character in the text.

### Special Characters and Metacharacters

Regular expressions use several special characters (also known as "metacharacters") that have specific functions:

- `.`: Matches any single character except newline (`\n`).
- `^`: Anchors the match at the start of the string.  
  - Regex: `^cat `  
  - Matches: "cat" in "catapult" but not in "concatenate"
- `$`: Anchors the match at the end of the string.    
  - Regex: `end$`  
  - Matches: "end" in "friend" but not in "endless"
- `*`: Matches zero or more of the preceding element (this can be a single character, a class of characters, or a grouped subpattern)..  
  - Regex: bo*  
  - Matches: "b", "bo", "boo", "booo", etc., in "A ghost booooed"  
  - Note: "b" followed by zero or more "o"s.  
- `+`: Matches one or more of the preceding element. (The `+` character is similar to the `*` but it requires at least one or more occurrences of the preceding element to be present for a match.)  
  - Regex: bo+  
  - Matches: "bo", "boo", "booo", etc., in "A ghost booooed"  
  - Note: "b" followed by at least one "o".  
- `?`: Makes the preceding element optional (It matches zero or one occurrence of the preceding element).
  - Regex: colou?r
  - Matches: Both "color" and "colour"
- `\`: Escapes a special character, treating it as a literal.


### Character Classes

- `[abc]`: Matches any single character in the brackets (a, b, or c).
- `[^abc]`: Matches any single character not in the brackets.
- `[a-z]`: Matches any single character in the range from a to z.
- `[A-Z]`: Matches any single character in the range from A to Z.
- `[0-9]`: Matches any single digit.

### Predefined Character Classes

- `\d`: Matches any digit, equivalent to [0-9].
- `\D`: Matches any non-digit.
- `\s`: Matches any whitespace character (space, tab, newline).
- `\S`: Matches any non-whitespace character.
- `\w`: Matches any word character (letters, digits, underscore), equivalent to [a-zA-Z0-9_].
- `\W`: Matches any non-word character.

### Quantifiers

Quantifiers specify how many instances of a character, group, or character class must be present in the target string for a match to be found:

- `{n}`: Matches exactly `n` times.
- `{n,}`: Matches at least `n` times.
- `{n,m}`: Matches between `n` and `m` times, inclusively.

### Grouping and Capturing

- `(abc)`: Matches the characters `abc` and remembers the match.
- `(?:abc)`: Matches the characters `abc` but does not remember the match (non-capturing group).

### Alternation

- `a|b`: Matches either `a` or `b`.

### Assertions

- `\b`: Matches a word boundary (the position between a word character and a non-word character).
- `\B`: Matches only when not at a word boundary.

### Flags

Flags can be included in the pattern to modify its behavior:

- `i`: Case-insensitive matching.
- `m`: Multiline mode (changes the behavior of `^` and `$` to match the start and end of each line).
- `s`: Dotall mode (makes . match newlines).

When using these in Go, you can incorporate them directly into your pattern string or use syntax like `(?ims)` at the beginning of your regex string to set multiple flags.

## Note 

In Go, both backquotes (\` \`) and double quotes (" ") can be used to delimit strings, but they serve different purposes and behave in subtly different ways:

1. **Backquotes**: These are used to create raw string literals. The contents between the backquotes are taken exactly as they are, including any newlines, tabs, and other special characters, without the need for escaping them. This makes raw string literals particularly useful for regular expressions, as they often contain backslashes (\` \`) that would otherwise need to be doubled up (escaped) when using double quotes.
2. **Double Quotes** (" "): These create interpreted string literals, where certain escape sequences (like \n for a newline, `\t` for a tab, and `\\` for a literal backslash) are processed and converted to their actual character values.

For regular expressions, backquotes are generally preferred because:

You avoid the need to escape backslashes. In regex patterns, backslashes are used very frequently (e.g., `\b`, `\w`, `\s`). If you were to use double quotes, **every backslash in the pattern would need to be escaped** (`\\b`, `\\w`, `\\s`), making the regex less readable and more prone to errors.

## Using `()` While Creating a `Regexp`
In Go, once you use parentheses to create capturing groups in a regular expression, you can access the captured content by using methods provided by the regexp package, such as `FindStringSubmatch`. This method returns a slice of strings where each element corresponds to a part of the string matched by the entire regular expression and each subsequent capturing group.

Here’s a simple example to demonstrate how you can define a regex with capturing groups and then access the captured content:

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // Define a regex pattern with capturing groups for year, month, and day
    dateRegex := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

    // Sample date string
    dateStr := "2023-04-15"

    // Use FindStringSubmatch to extract parts
    matches := dateRegex.FindStringSubmatch(dateStr)
    fmt.Println(matches)
}
```

**Note:**
When using the `MatchString` method from the regexp package in Go with regex patterns that include capturing groups (`()`), the purpose of the capturing groups within `MatchString` does not change the behavior of the match check. `MatchString` simply checks if the string matches the regex pattern and returns a boolean result (`true` or `false`), but it does not return any of the contents of the capturing groups.

# Examples

## Example 1:

This example defines a pattern for matching strings that look like U.S. Social Security numbers, using quantifiers and the case-insensitive flag.

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    pattern := `(?i)^\d{3}-\d{2}-\d{4}$`  // Case-insensitive, matches a specific format like "123-45-6789"
    re := regexp.MustCompile(pattern)
    fmt.Println(re.MatchString("123-45-6789"))  // Outputs: true
}
```

## Example 2:

```go
re := regexp.MustCompile(`(\b\w+\b[.,:;']*)\s+\((low|up|case|hex|bin|cap)\)\s`)
```

- \` to define raw string literals.
- `\b` for boundaries (pos between word and non-word).
- `[...]` Matches any single character in the brackets.
- `*` Matches zero or more of the preceding element.
- `+` Matches one or more of the preceding element.
- `\s` Matches any whitespace character.
- `\w` Matches any word character.
- Note: I need to use `()` so it won't go and include all the string before `low` as an or (`|`) expression, the other pranthese `()` before (I mean here: (\b\w+\b[.,:;']*)) are optional and we can safely remove them.

[Great Tool](https://regex101.com/)