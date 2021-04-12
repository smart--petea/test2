package proto

import (
    "context"
    "google.golang.org/grpc"
    "time"
    "github.com/spf13/viper"
)

func ServiceCall(fsyms, tsyms []string) (*TServiceResponse, error) {
    grpcHost := viper.GetString("grpc.host")
    grpcPort := viper.GetString("grpc.port")
    grpcAddress := grpcHost + ":" +  grpcPort

    serviceConn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
    if err != nil {
        return nil, err
    }
    defer serviceConn.Close()


    service := NewTServiceClient(serviceConn)
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    return service.Get(ctx, &TServiceRequest{
        Fsyms: fsyms,
        Tsyms: tsyms,
    })
}
