package restserver

import (
    "context"
    "google.golang.org/grpc"
    "github.com/smart--petea/test2/pkg/proto"
    "github.com/smart--petea/test2/pkg/common"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "github.com/gin-gonic/gin"
    "time"
    "strings"
    "github.com/spf13/viper"
    "log"
)

func init() {
    service :="restserver"
    if err := common.InitLogForService(service); err != nil {
        panic(err)
    }

    if err := common.InitViperForService(service); err != nil {
        panic(err)
    }
}

func Run() error {
    ginServer := gin.Default()
    ginServer.GET("/service/price", func(c *gin.Context) {
        fsyms := c.Query("fsyms")
        tsyms := c.Query("tsyms")

        if fsyms == "" {
            c.JSON(400, gin.H{"msg": "fsyms not set"})
            return
        }

        if tsyms == "" {
            c.JSON(400, gin.H{"msg": "tsyms not set"})
            return
        }

        Fsyms := strings.Split(fsyms, ",")
        Tsyms := strings.Split(tsyms, ",")

        serviceRes, err := serviceCall(Fsyms, Tsyms)
        if err != nil {
            st, ok := status.FromError(err)
            if !ok {
                c.JSON(500, gin.H{"msg": err.Error()})
                return
            }

            switch st.Code() {
            case codes.InvalidArgument:
                c.JSON(400, gin.H{"msg": st.Message()})
                return
            default:
                log.Printf("Unknows status error code='%+v' message='%s'", st.Code(), st.Message())
                c.JSON(500, gin.H{"msg": st.Message()})
                return
            }
        }

        c.JSON(200, serviceRes)
    })

    httpHost := viper.GetString("http.host")
    httpPort := viper.GetString("http.port")
    httpAddress := httpHost + ":" + httpPort 
    ginServer.Run(httpAddress)

    return nil
}

func serviceCall(fsyms, tsyms []string) (*proto.TServiceResponse, error) {
    grpcHost := viper.GetString("grpc.host")
    grpcPort := viper.GetString("grpc.port")
    grpcAddress := grpcHost + ":" +  grpcPort

    serviceConn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
    if err != nil {
        return nil, err
    }
    defer serviceConn.Close()


    service := proto.NewTServiceClient(serviceConn)
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    return service.Get(ctx, &proto.TServiceRequest{
        Fsyms: fsyms,
        Tsyms: tsyms,
    })
}
