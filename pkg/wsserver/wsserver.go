package wsserver

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "encoding/json"

    "github.com/smart--petea/test2/pkg/proto"
    "github.com/gorilla/websocket"
    "github.com/smart--petea/test2/pkg/common"
    "github.com/spf13/viper"
)

func init() {
    service :="wsserver"
    if err := common.InitLogForService(service); err != nil {
        panic(err)
    }

    if err := common.InitViperForService(service); err != nil {
        panic(err)
    }
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }

    log.Println("Client Connected")
    fsyms := "BTC"
    tsyms := "USD"
    Fsyms := strings.Split(fsyms, ",")
    Tsyms := strings.Split(tsyms, ",")

    serviceRes, err := proto.ServiceCall(Fsyms, Tsyms)
    var data []byte
    if err != nil {
        data = []byte(err.Error())
    } else {
        data, err = json.Marshal(serviceRes)
        if err != nil {
            data = []byte(err.Error())
        }
    }

    err = ws.WriteMessage(1, data)
    if err != nil {
        log.Println(err)
    }

    reader(ws)
}

func setupRoutes() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", wsEndpoint)
}

func Run() error {
    fmt.Println("Hello World")
    setupRoutes()

    httpHost := viper.GetString("http.host")
    httpPort := viper.GetString("http.port")
    httpAddress := httpHost + ":" + httpPort
    return http.ListenAndServe(httpAddress, nil)
}

func reader(conn *websocket.Conn) {
    for {
        // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}
