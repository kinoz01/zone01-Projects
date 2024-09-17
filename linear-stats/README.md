## How to Test the Program

To test the program, move your test executable file (`linear-stats`) ([stats-tests](https://assets.01-edu.org/stats-projects/stat-bin-dockerized.zip)) to the project folder, and then simply run:

```bash
go test
```

### Test using docker:

1. copy the `bin` folder to the project folder.
2. open the `calc_test.go` file and replace the `"./linear-stats"` in code `line 36` by:

```go
`"./stat-bin-dockerized/stat-bin/run.sh", "linear-stats"`
```

3. run:

```bash
go test
```

> Note: when using docker make sure you already run the program using docker before using `go test`.