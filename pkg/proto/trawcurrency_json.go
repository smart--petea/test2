package proto

import (
    "fmt"
)

func (t *TRawCurrency) Fill(dataRaw interface{}) error {
    data, ok := dataRaw.(map[string]interface{}) 
    if !ok {
        return fmt.Errorf("TRawCurrency can't convert %+v to map[string]interface{}", dataRaw)
    }


    if PRICE, ok := data["PRICE"].(float64); ok {
        t.PRICE = PRICE
    } else {
        return fmt.Errorf("PRICE %+v can't be converted to float64", data["PRICE"])
    }

    if LASTUPDATE, ok := data["LASTUPDATE"].(float64); ok {
        t.LASTUPDATE = int64(LASTUPDATE)
    } else {
        return fmt.Errorf("LASTUPDATE %+v can't be converted to int64", data["LASTUPDATE"])
    }

    if VOLUME24HOUR, ok := data["VOLUME24HOUR"].(float64); ok {
        t.VOLUME24HOUR = VOLUME24HOUR
    } else {
        return fmt.Errorf("VOLUME24HOUR %+v can't be converted to float64", data["VOLUME24HOUR"])
    }

    if VOLUME24HOURTO, ok := data["VOLUME24HOURTO"].(float64); ok {
        t.VOLUME24HOURTO = VOLUME24HOURTO
    } else {
        return fmt.Errorf("VOLUME24HOURTO %+v can't be converted to float64", data["VOLUME24HOURTO"])
    }

    if OPEN24HOUR, ok := data["OPEN24HOUR"].(float64); ok {
            t.OPEN24HOUR = OPEN24HOUR
    } else {
        return fmt.Errorf("OPEN24HOUR %+v can't be converted to float64", data["OPEN24HOUR"])
    }

    if HIGH24HOUR, ok :=  data["HIGH24HOUR"].(float64); ok {
        t.HIGH24HOUR = HIGH24HOUR
    } else {
        return fmt.Errorf("HIGH24HOUR %+v can't be converted to float64", data["HIGH24HOUR"])
    }

    if LOW24HOUR, ok := data["LOW24HOUR"].(float64); ok {
        t.LOW24HOUR = LOW24HOUR
    } else {
        return fmt.Errorf("LOW24HOUR %+v can't be converted to float64", data["LOW24HOUR"])
    }

    if CHANGE24HOUR, ok := data["CHANGE24HOUR"].(float64); ok {
        t.CHANGE24HOUR = CHANGE24HOUR
    } else {
        return fmt.Errorf("CHANGE24HOUR %+v can't be converted to float64", data["CHANGE24HOUR"])
    }

    if CHANGEPCT24HOUR, ok := data["CHANGEPCT24HOUR"].(float64); ok {
        t.CHANGEPCT24HOUR = CHANGEPCT24HOUR
    } else {
        return fmt.Errorf("CHANGEPCT24HOUR %+v can't be converted to float64", data["CHANGEPCT24HOUR"])
    }

    if SUPPLY, ok := data["SUPPLY"].(float64); ok {
        t.SUPPLY = int64(SUPPLY)
    } else {
        return fmt.Errorf("SUPPLY %+v can't be converted to float64", data["SUPPLY"])
    }

    if MKTCAP, ok := data["MKTCAP"].(float64); ok {
        t.MKTCAP = MKTCAP
    } else {
        return fmt.Errorf("MKTCAP %+v can't be converted to float64", data["MKTCAP"])
    }

    return nil
}
