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
     "time"
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

func (t *tServiceServer) Get(ctx context.Context, req *proto.TServiceRequest) (*proto.TServiceResponse, error) {
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
        log.Printf("http request encountered error %+v", err)

        tServiceResponse, err = DbGetTServiceResponse(ctx, t.db, req.Fsyms, req.Tsyms)
        if err != nil {
            return nil, status.Error(codes.Internal, err.Error())
        }
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

    if err := RunScheduler(ctx, db); err != nil {
        return err
    }

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

func RunScheduler(ctx context.Context, db *sql.DB) error {
    interval := viper.GetInt("scheduler.interval")
    if interval < 1 {
        return fmt.Errorf("Wrong scheduler interval %d", interval)
    }

    ticker := time.NewTicker(time.Duration(interval) * time.Second)
    log.Printf("starting Scheduler at interval %d seconds", interval)

    fsyms := strings.Join(viper.GetStringSlice("currencies.fsyms"), ",")
    tsyms := strings.Join(viper.GetStringSlice("currencies.tsyms"), ",")

    go func() {
        for {
            select {
            case <- ticker.C:
                tServiceResponse, err := GetFromHttp(ctx, fsyms, tsyms)
                if err != nil {
                    log.Printf("[Scheduler] something wrong msg='%s'", err.Error())
                    continue
                } 

                log.Printf("[Scheduler] saving...")
                DbSaveTServiceResponse(ctx, db, tServiceResponse)
            case <- ctx.Done():
                ticker.Stop()
                log.Printf("[Scheduler] shutting down ...")
                return
            }
        }
    }()

    return nil
}

func DbGetNextId(ctx context.Context, db *sql.DB) (int, error) {
    query := fmt.Sprintf("SELECT COALESCE(MAX(id), 0) + 1 as id FROM pairs")
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return 0, err
    }
    defer rows.Close()

    rows.Next()

    var id int
    if err := rows.Scan(&id); err != nil {
        return 0, err
    }

    return id, nil
}

func DbGetTServiceResponse(ctx context.Context, db *sql.DB, Fsyms []string, Tsyms []string) (*proto.TServiceResponse, error) {
    fsyms := "'" + strings.Join(Fsyms, "','") + "'"
    tsyms := "'" + strings.Join(Tsyms, "','") + "'"
    query := fmt.Sprintf("SELECT fsym, tsym, raw, display FROM pairs WHERE (NOW() - interval '2 minutes') < updated_at AND fsym in (%s) AND tsym in (%s) ", fsyms, tsyms)
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tServiceResponse proto.TServiceResponse
    var raw, display, fsym, tsym string
    for rows.Next() {
        if err := rows.Scan(&fsym, &tsym, &raw, &display); err != nil {
            return nil, err
        }

        var tRawCurrency proto.TRawCurrency
        if err := json.Unmarshal([]byte(raw), &tRawCurrency); err != nil {
            return nil, err
        }

        var tDisplayCurrency proto.TDisplayCurrency
        if err := json.Unmarshal([]byte(display), &tDisplayCurrency); err != nil {
            return nil, err
        }

        if tServiceResponse.RAW == nil {
            tServiceResponse.RAW = make(map[string]*proto.TRaw)
        }

        if tServiceResponse.DISPLAY == nil {
            tServiceResponse.DISPLAY = make(map[string]*proto.TDisplay)
        }

        if _, ok := tServiceResponse.RAW[fsym]; !ok  {
            var tRaw proto.TRaw
            tRaw.Currencies = make(map[string]*proto.TRawCurrency)
            tServiceResponse.RAW[fsym] = &tRaw
        }
        tServiceResponse.RAW[fsym].Currencies[tsym] = &tRawCurrency

        if _, ok := tServiceResponse.DISPLAY[fsym]; !ok  {
            var tDisplay proto.TDisplay
            tDisplay.Currencies = make(map[string]*proto.TDisplayCurrency)
            tServiceResponse.DISPLAY[fsym] = &tDisplay
        }
        tServiceResponse.DISPLAY[fsym].Currencies[tsym] = &tDisplayCurrency
    }

    return &tServiceResponse, nil
}

func DbGetIdFsymTsym(ctx context.Context, db *sql.DB, fsym, tsym string) (int, error) {
    query := fmt.Sprintf("SELECT id FROM pairs WHERE fsym='%s' AND tsym='%s' LIMIT 1", fsym, tsym)
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return 0, err
    }
    defer rows.Close()

    if rows.Next() == false {
        return 0, nil
    }

    var id int
    if err := rows.Scan(&id); err != nil {
        return 0, err
    }

    return id, nil
}

func DbGetInsertFsymTsym(ctx context.Context, db *sql.DB, id int, fsym, tsym, raw, display string) error {
    query := `INSERT INTO public.pairs (id, fsym, tsym, raw, display) VALUES(%d, '%s', '%s', '%s', '%s')`
    query = fmt.Sprintf(query, id, fsym, tsym, raw, display)
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return err
    }
    defer rows.Close()

    return nil
}

func DbGetUpdateFsymTsym(ctx context.Context, db *sql.DB, id int, fsym, tsym, raw, display string) error {
    query := `UPDATE public.pairs SET fsym='%s', tsym='%s', updated_at=NOW(), raw='%s', display='%s' WHERE id=%d`
    query = fmt.Sprintf(query, fsym, tsym, raw, display, id)
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return err
    }
    defer rows.Close()

    return nil
}

func DbSaveTServiceResponse(ctx context.Context, db *sql.DB, tServiceResponse *proto.TServiceResponse) {
    for fsym, fsymTRaw := range tServiceResponse.RAW {
        for tsym, tsymRawCurrency := range fsymTRaw.Currencies {
            id, err := DbGetIdFsymTsym(ctx, db, fsym, tsym)
            if err != nil {
                log.Printf("error %+v\n", err)
            }

            tsymDisplayCurrency := tServiceResponse.DISPLAY[fsym].Currencies[tsym]
            display, err := json.Marshal(tsymDisplayCurrency)
            if err != nil {
                log.Printf("error %+v\n", err)
                continue
            }

            raw, err := json.Marshal(tsymRawCurrency)
            if err != nil {
                log.Printf("error %+v\n", err)
                continue
            }

            if id == 0 {
                id, err := DbGetNextId(ctx, db)
                if err != nil {
                    log.Printf("error %+v\n", err)
                    continue
                }

                err = DbGetInsertFsymTsym(ctx, db, id, fsym, tsym, string(raw), string(display))
                if err != nil {
                    log.Printf("error %+v\n", err)
                    continue
                }
            } else {
                err := DbGetUpdateFsymTsym(ctx, db, id, fsym, tsym, string(raw), string(display))
                if err != nil {
                    log.Printf("error %+v\n", err)
                    continue
                }
            }
        }
    }
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
