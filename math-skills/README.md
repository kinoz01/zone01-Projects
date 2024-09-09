## How to Test the Program

To test the program, move your test executable file (`math-skills`) to the project folder, and then simply run:

```bash
go test
```

### Test using docker:

1. copy the `bin` folder to the project folder.
2. open the `calc_test.go` file and replace the `"./math-skills"` in code `line 36` by:

```go
`"./stat-bin-dockerized/stat-bin/run.sh", "math-skills"`
```

3. run:

```bash
go test
```
