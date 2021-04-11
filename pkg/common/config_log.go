package common

import (
    "log"
    "os"
)

func ConfigLog(servicename string) error {
    file, err := GetLogFileForService(servicename)
    if err != nil {                                                                                                  
        return err
    }                                                                                                                
    log.SetOutput(file)                                                                                              
    log.SetFlags(/*log.Lshortfile | */log.Ldate | log.Ltime)  

    return nil
}

func GetLogFileForService(servicename string) (*os.File, error) {
    wd, err := os.Getwd()
    if err != nil {
        return nil, err
    }

    return os.OpenFile(wd + "/logs/" + servicename + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)                            
}
