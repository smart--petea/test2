package main

import (
    "github.com/smart--petea/test2/pkg/restserver"
    "os"
    "fmt"
)

func main() {
    if err := restserver.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }
}
