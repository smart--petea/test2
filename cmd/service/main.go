package main

import (
    "fmt"
    "os"

    cmd "github.com/smart--petea/test2/pkg/cmd/service"
)

func main() {
    if err := cmd.RunService(); err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }
}
