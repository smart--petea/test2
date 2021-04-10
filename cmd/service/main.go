package main

import (
    "fmt"
    "os"

    service "github.com/smart--petea/test2/pkg/service"
)

func main() {
    if err := service.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }
}
