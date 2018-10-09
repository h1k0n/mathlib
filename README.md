# mathlib

## how to use
main.go
```
package main // import "github.com/h1k0n/usemathlib"

import (
        "fmt"

        "github.com/h1k0n/mathlib"
)

func main() {
    fmt.Println(mathlib.Add(1,2))
}
```

go.mod
```
module github.com/h1k0n/usemathlib

require github.com/h1k0n/mathlib v0.0.0-20181009113559-20f1358a6b70
```