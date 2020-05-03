# Prerequisites 
- Go 1.14

# run code
```bash
go build
./finplay

```

# test code
```bash
go test

# run single test function
go test -timeout 30s github.com/valentinsavenko/finplay -run ^(TestParseCSV)$

```