package service

import (
    "log"
    "context"
    "database/sql"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/grpc"
    _ "github.com/lib/pq"

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

type tServiceServer struct {
    db *sql.DB
}

func NewTServiceServer(db *sql.DB) proto.TServiceServer {
    return &tServiceServer{db:db}
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

    if err := ValidateFsyms(req.Fsyms); err != nil {
        return nil, status.Error(codes.InvalidArgument, err.Error())
    }

    if err := ValidateTsyms(req.Tsyms); err != nil {
        return nil, status.Error(codes.InvalidArgument, err.Error())
    }

    fsyms := strings.Join(req.Fsyms, ",")
    tsyms := strings.Join(req.Tsyms, ",")

    tServiceResponse, err := GetFromHttp(ctx, fsyms, tsyms)
    if err != nil {
        //return GetFromDB(fsyms, tsyms)
        panic(err)
    }

    return tServiceResponse, nil
}

func ValidateFsyms(fsyms []string) error {
    cFsyms := viper.GetStringSlice("currencies.fsyms")

    FSYMS_OUTER_LOOP:
    for _, fsym := range fsyms {
        for _, cfsym := range cFsyms {
            if fsym == cfsym {
                continue FSYMS_OUTER_LOOP
            }
        }

        return fmt.Errorf("fsym '%s' is not allowed", fsym)
    }

    return nil
}

func ValidateTsyms(tsyms []string) error {
    cTsyms := viper.GetStringSlice("currencies.tsyms")

    TSYMS_OUTER_LOOP:
    for _, tsym := range tsyms {
        for _, ctsym := range cTsyms {
            if tsym == ctsym {
                continue TSYMS_OUTER_LOOP
            }
        }

        return fmt.Errorf("tsym '%s' is not allowed", tsym)
    }

    return nil
}

func GetFromHttp(ctx context.Context, fsyms, tsyms string) (*proto.TServiceResponse, error) {
    url := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%s&tsyms=%s", fsyms, tsyms)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req = req.WithContext(ctx)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var tServiceResponse proto.TServiceResponse
    err = json.Unmarshal(body, &tServiceResponse)
    if err != nil {
        return nil, err
    }

    return &tServiceResponse, nil
}

func Run() error {
    if err := ValidateConfiguration(); err != nil {
        return err
    }

    db, err := GetDB()
    if err != nil {
        return err
    }
    defer db.Close()

    protoAPI := NewTServiceServer(db)
    ctx := context.Background()

    //RunScheduler(ctx)
    return RunServer(ctx, protoAPI)
}

func ValidateConfiguration() error {
    fsyms := viper.GetStringSlice("currencies.fsyms")
    if len(fsyms) == 0 {
        return fmt.Errorf("currencies.fsyms is empty. Please config it")
    }

    tsyms := viper.GetStringSlice("currencies.tsyms")
    if len(tsyms) == 0 {
        return fmt.Errorf("currencies.tsyms is empty. Please config it")
    }

    return nil
}

func GetDB() (*sql.DB, error) {
    psqlHost := viper.GetString("db.host")
    psqlPort := viper.GetInt("db.port")
    psqlUser := viper.GetString("db.user")
    psqlPassword := viper.GetString("db.password")
    psqlDbname := viper.GetString("db.name")
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", psqlHost, psqlPort, psqlUser, psqlPassword, psqlDbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}

/*
func RunScheduler(ctx context.Context, db *sql.DB) error {
    interval := viper.GetInt("scheduler.interval", 0)
    if interval < 1 {
        return fmt.Errorf("Wrong scheduler interval %d", interval)
    }

    ticker := time.NewTicker(interval * timeSecond)
    log.Printf("starting Scheduler at interval %d seconds", interval)
    go func() {
        for {
            select {
            case <- ticker.C:

            case <- ctx.Done():
                ticker.Stop()
                log.Printf("shutting down Scheduler")
                return
            }
        }
    }

    return nil
}
*/

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
