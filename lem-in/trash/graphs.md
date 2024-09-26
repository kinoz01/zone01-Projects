```go
// Create the graph from the room connections
	graph := Graph{
		"start": {"t", "h", "0"},
		"h":     {"A", "n"},
		"t":     {"E"},
		"0":     {"o"},
		"E":     {"a"},
		"A":     {"c"},
		"o":     {"n"},
		"n":     {"e", "m"},
		"e":     {"end"},
		"m":     {"end"},
		"c":     {"k"},
		"a":     {"m"},
		"k":     {"end"},
	}
```

```md
https://medium.com/@jamierobertdawson/lem-in-finding-all-the-paths-and-deciding-which-are-worth-it-2503dffb893
```
