package main

import (
    "github.com/smart--petea/test2/pkg/wsserver"
    "os"
    "fmt"
)

func main() {
    if err := wsserver.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }
}
