package common

import (
    "log"
    "os"
    "github.com/spf13/viper"
)

func InitLogForService(servicename string) error {
    wd, err := os.Getwd()
    if err != nil {
        return err
    }

    logsDirectory := wd + "/logs"
    filename := logsDirectory + "/" + servicename + ".log"
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)                            
    if err != nil {                                                                                                  
        return err
    }                                                                                                                
    log.SetOutput(file)                                                                                              
    log.SetFlags(/*log.Lshortfile | */log.Ldate | log.Ltime)  

    return nil
}

func InitViperForService(service string) error {
    wd, err := os.Getwd()
    if err != nil {
        return err
    }
    configDirectory := wd + "/configs"

    viper.SetConfigName(service)
    viper.SetConfigType("yaml")
    viper.AddConfigPath(configDirectory)
    err = viper.ReadInConfig()
    if err != nil {
        panic(err)
    }

    return nil
}
