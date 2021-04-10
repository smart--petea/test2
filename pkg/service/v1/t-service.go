package v1

import (
    //"log"
    "context"
    "database/sql"
    //"github.com/golang/protobuf/ptypes/empty"
    //"google.golang.org/grpc/codes"
    //"google.golang.org/grpc/status"

    "github.com/smart--petea/test2/pkg/api/v1"
)

// userServiceServer is implementation of v1.UserServiceServer proto interface
type tServiceServer struct {
    db *sql.DB
}

/*
func NewUserServiceServer(db *sql.DB) v1.UserServiceServer {
    return &userServiceServer{db: db}
}
*/

func NewTServiceServer() v1.TServiceServer {
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
func (t *tServiceServer) Get(ctx context.Context, req *v1.TServiceRequest) (*v1.TServiceResponse, error) {
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

    return &v1.CreateUserResponse{
        Id:  id,
    }, nil
    */
    rawCurrencies := make(map[string]*v1.TRawCurrency)
    rawCurrencies["USD"] = &v1.TRawCurrency{
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
    raw := make(map[string]*v1.TRaw)
    raw["BTC"] = &v1.TRaw{Currencies: rawCurrencies}

    displayCurrencies := make(map[string]*v1.TDisplayCurrency)
    displayCurrencies["USD"] = &v1.TDisplayCurrency{
        Price: 60600.96,
        Volume24Hour:33535.97002894,
        Volume24Hourto: 1995648478.1027222,
        Open24Hour: 58508.56,
        High24Hour: 61217.34,
        Low24Hour: 57862.88,
        Change24Hour: 2092.4000000000015, 
        Changepct24Hour: 3.5762288458304243,
    }
    display := make(map[string]*v1.TDisplay)
    display["BTC"] = &v1.TDisplay{Currencies: displayCurrencies}

    return &v1.TServiceResponse{
        Raw: raw,
        Display: display,
    }, nil
}
