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

`MatchString`: Checks if a string matches the pattern.

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

When you pass `-1` as the second argument, it tells the method to find all possible matches in the string. This is commonly used when you want to retrieve every match without any restriction on the count.

If you specify a non-negative integer `n`, then the method returns at most `n` matches. This can be useful if you're only interested in the first few matches and want to limit the output to improve performance or reduce output size.

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

### Extracting Submatches

`FindStringSubmatch`: Returns slices containing matched parts of the string including capturing groups.

```go
submatches := re.FindStringSubmatch("search in this string")
```

##