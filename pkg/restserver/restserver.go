package restserver

import (
    "log"
    "context"
    //"github.com/golang/protobuf/ptypes/empty"
    //"google.golang.org/grpc/codes"
    //"google.golang.org/grpc/status"
    "google.golang.org/grpc"

    "github.com/smart--petea/test2/pkg/proto"
    "github.com/smart--petea/test2/pkg/common"
    "flag"
    "github.com/gin-gonic/gin"
    "time"
)

func init() {
    common.ConfigLog("restserver") 
}

func Run() error {
    address := flag.String("server", "", "gRPC server in format host:port")
    flag.Parse()

    serviceConn, err := grpc.Dial(*address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer serviceConn.Close()
    log.Println(" connection state ====> ", serviceConn.GetState())

    ginServer := gin.Default()
    ginServer.GET("/service/price", func(c *gin.Context) {
        service := proto.NewTServiceClient(serviceConn)

        ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancel()

        seviceReq := proto.TServiceRequest {
            Fsyms: []string{"BTC"},
            Tsyms: []string{"USD,EUR"},
        }

        serviceRes, err := service.Get(ctx, &seviceReq)
        if err != nil {
            //todo
            log.Fatalf("Create user failed: %v", err)
        }

        c.JSON(200, serviceRes)
    })

    ginServer.Run()

    return nil
}
