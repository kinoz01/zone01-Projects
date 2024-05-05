```go
asciiTable[i] = append(asciiTable[i], lines...)
```

This line of code is working with slices in Go. Here's a detailed breakdown:

`asciiTable[i]`: This accesses the ith element of the asciiTable slice. In Go, you can access elements of a slice using square brackets `[]`.

`append()`: In Go, `append()` is a built-in function used to append elements to a slice. It takes a slice and one or more elements to append, and it returns a new slice with the appended elements.  
`lines...`: The `...` syntax is called a **variadic parameter** in Go. When you see `...` after a slice name, it unpacks the slice into individual elements. In this case, `lines...` takes each element of the lines slice and passes them as separate arguments to the `append()` function.  
So, putting it all together, `append(asciiTable[i], lines...)` takes each line from the lines slice and appends them to the end of the ith element of the asciiTable slice.