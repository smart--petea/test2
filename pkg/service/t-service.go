package service

import (
    "log"
    "context"
    "database/sql"
    //"github.com/golang/protobuf/ptypes/empty"
    //"google.golang.org/grpc/codes"
    //"google.golang.org/grpc/status"
    "google.golang.org/grpc"

    "github.com/smart--petea/test2/pkg/proto"
    "github.com/smart--petea/test2/pkg/common"
    "fmt"
    "os"
    "os/signal"
    "net"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
     "github.com/spf13/viper"
)

func init() {
    service :="service"
    if err := common.InitLogForService(service); err != nil {
        panic(err)
    }

    if err := common.InitViperForService(service); err != nil {
        panic(err)
    }
}

// tServiceServer is implementation of proto.TServiceServer proto interface
type tServiceServer struct {
    db *sql.DB
}

/*
func NewUserServiceServer(db *sql.DB) proto.UserServiceServer {
    return &userServiceServer{db: db}
}
*/

func NewTServiceServer() proto.TServiceServer {
    return &tServiceServer{}
}

// connect returns SQL database connection from the pool
/*
func (s *userServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
    c, err := s.db.Conn(ctx)
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
    }
    return c, nil
}
*/

// Create new user
func (t *tServiceServer) Get(ctx context.Context, req *proto.TServiceRequest) (*proto.TServiceResponse, error) {
    /*
    conn, err := s.connect(ctx)
    if err != nil {
        log.Println("Error in connecting to database", err)
    }
    defer conn.Close()

    result, err := conn.ExecContext(context.Background(), "Insert into user(name, email) values(?,?)", req.GetUser().GetName(), req.GetUser().GetEmail())
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to insert into user-> "+err.Error())
    }

    id, err := result.LastInsertId()
    if err != nil {
        return nil, status.Error(codes.Unknown, "failed to retrieve id for created ToDo-> "+err.Error())
    }

    return &proto.CreateUserResponse{
        Id:  id,
    }, nil
    */

    fsyms := strings.Join(req.Fsyms, ",")
    tsyms := strings.Join(req.Tsyms, ",")
    url := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%s&tsyms=%s", fsyms, tsyms)
    resp, err := http.Get(url)
    if err != nil {
        //todo
        //1. ask the db for the last info
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        //todo
        panic(err)
    }

    var tServiceResponse proto.TServiceResponse
    err = json.Unmarshal(body, &tServiceResponse)
    if err != nil {
        //todo
        fmt.Printf("%+v", err)
        panic(err)
    }

    return &tServiceResponse, nil
}

/*
type Config struct {
    GRPCPort string
    DatastoreDBHost string
    DatastoreDBUser string
    DatastoreDBPassword string
    DatastoreDBSchema string
}
*/

func Run() error {
    ctx := context.Background()

    /*
    var cfg Config
    flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
    flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
    flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
    flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
    flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
    flag.Parse()
    */

    /*
    param := "parseTime=true"

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
        cfg.DatastoreDBUser,
        cfg.DatastoreDBPassword,
        cfg.DatastoreDBHost,
        cfg.DatastoreDBSchema,
        param)
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return fmt.Errorf("failed to open database: %v", err)
    }

    if err = db.Ping(); err != nil {
        panic(err.Error())
    }
    defer db.Close()
    protoAPI := proto.NewUserServiceServer(db)
    */

    protoAPI := NewTServiceServer()

    return RunServer(ctx, protoAPI)
}

func RunServer(ctx context.Context, protoAPI proto.TServiceServer) error {
    grpcHost := viper.GetString("grpc.host")
    grpcPort := viper.GetString("grpc.port")
    grpcAddress := grpcHost + ":" +  grpcPort

    listen, err := net.Listen("tcp", grpcAddress)
    if err != nil {
        return err
    }

    server := grpc.NewServer()
    proto.RegisterTServiceServer(server, protoAPI)

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        for range c {
            log.Println("shutting down gRPC server...")

            server.GracefulStop()

            <-ctx.Done()
        }
    }()

    log.Println("starting gRPC server...")
    return server.Serve(listen)
}
