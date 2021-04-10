package cmd

import (
    "context"
    "flag"
    "fmt"

//    "database/sql"
//    _ "github.com/lib/pq"

    "github.com/smart--petea/test2/pkg/protocol/grpc"
    "github.com/smart--petea/test2/pkg/service/v1"
)

type Config struct {
    GRPCPort string
    DatastoreDBHost string
    DatastoreDBUser string
    DatastoreDBPassword string
    DatastoreDBSchema string
}

func RunService() error {
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
    v1API := v1.NewUserServiceServer(db)
    */

    v1API := v1.NewTServiceServer()

    return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
