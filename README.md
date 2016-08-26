# randomstrings
Random string generator in Golang.

#Installation
`go get -u github.com/anchenbaba/randomstrings`

# Sample code
```go
package main

import (
    "log"

    "github.com/anchenbaba/randomstrings"
)

func main() {
    r := randomstrings.RandomString(5, 1)
    log.Println(string(r))
    r2 := randomstrings.RandomStringAll(10)
    log.Println(r2)
}
```