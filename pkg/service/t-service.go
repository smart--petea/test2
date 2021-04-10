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
    "flag"
    "fmt"
    "os"
    "os/signal"
    "net"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

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

    resp, err := http.Get("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=BTC&tsyms=USD,EUR")
    if err != nil {
        //todo
        //1. ask the db for the last info
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    log.Printf("response %+v", string(body))

    var rawDisplay struct{
        RAW string
        DISPLAY string
    }
    err = json.Unmarshal(body, &rawDisplay)
    if err != nil {
        //todo
        fmt.Printf("%+v", err)
        panic(err)
    }
    log.Printf("rawDisplay %+v", rawDisplay)
    

    rawCurrencies := make(map[string]*proto.TRawCurrency)
    rawCurrencies["USD"] = &proto.TRawCurrency{
        Price: 60600.96,
        Lastupdate:1618053792,
        Volume24Hour:33535.97002894,
        Volume24Hourto: 1995648478.1027222,
        Open24Hour: 58508.56,
        High24Hour: 61217.34,
        Low24Hour: 57862.88,
        Change24Hour: 2092.4000000000015, 
        Changepct24Hour: 3.5762288458304243,
        Supply: 18678725,
        Mktcap: 1131948666576,
    }
    raw := make(map[string]*proto.TRaw)
    raw["BTC"] = &proto.TRaw{Currencies: rawCurrencies}

    displayCurrencies := make(map[string]*proto.TDisplayCurrency)
    displayCurrencies["USD"] = &proto.TDisplayCurrency{
        Price: 60600.96,
        Volume24Hour:33535.97002894,
        Volume24Hourto: 1995648478.1027222,
        Open24Hour: 58508.56,
        High24Hour: 61217.34,
        Low24Hour: 57862.88,
        Change24Hour: 2092.4000000000015, 
        Changepct24Hour: 3.5762288458304243,
    }
    display := make(map[string]*proto.TDisplay)
    display["BTC"] = &proto.TDisplay{Currencies: displayCurrencies}

    return &proto.TServiceResponse{
        Raw: raw,
        Display: display,
    }, nil
}

type Config struct {
    GRPCPort string
    DatastoreDBHost string
    DatastoreDBUser string
    DatastoreDBPassword string
    DatastoreDBSchema string
}

func Run() error {
    ctx := context.Background()

    var cfg Config
    flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
    flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
    flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
    flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
    flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
    flag.Parse()

    if len(cfg.GRPCPort) == 0 {
        return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
    }

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

    return RunServer(ctx, protoAPI, cfg.GRPCPort)
}

func RunServer(ctx context.Context, protoAPI proto.TServiceServer, port string) error {
    listen, err := net.Listen("tcp", ":"+port)
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
