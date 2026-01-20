# docc
Simple ".docx" converter implemented by Go. Convert ".docx" to plain text.

## Note
This repository is an alpha version. Some disruptive changes could be applied.

## License
MIT

## Features
- Less dependency.
- No need for Microsoft Office.
- Only on limited environment, also ".doc" could be converted.
  - Windows in which MS Office has been installed.

## Usage

### As a package
This is a simple example to read all paragraphs.

```go
package main

import (
	"github.com/bax/go-docc"
	"filepath"
)

func main(){
    fp := filepath.Clean("./target.docx")
    r, err := docc.NewReader(fp)
    if err != nil {
        panic(err)
    }
    defer r.Close()
    ps, _ := r.ReadAll()
    // do something with ps:[]string
}
```
See docc_test.go for other examples.

Before compiling, you shall execute `go mod tidy` or `go get github.com/bax/go-docc`to get this package.

### As a binary
`go install` is available.

```shell
go install github.com/bax/go-docc/cmd/docc@latest
```

Then, `docc` command could be used. This is a simple example.

```shell
docc target.docx > plain.txt
```

## Contribution
Your contribution is really welcomed!

## Author
bax
