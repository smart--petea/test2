package main

import (
//    "github.com/golang/protobuf/ptypes/empty"
    "context"
    "flag"
    "log"
    "time"
    "google.golang.org/grpc"

    "github.com/smart--petea/test2/pkg/api/v1"
)

const (
    apiVersion = "v1"
)

func main() {
    address := flag.String("server", "", "gRPC server in format host:port")
    flag.Parse()

    conn, err := grpc.Dial(*address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    log.Println(" connection state ====> ", conn.GetState())

    c := v1.NewTServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    req := v1.TServiceRequest {
        Fsyms: []string{"BTC"},
        Tsyms: []string{"USD,EUR"},
    }

    res, err := c.Get(ctx, &req)
    if err != nil {
        log.Fatalf("Create user failed: %v", err)
    }
    log.Printf("Get info: <%+v>\n\n", res)
}
