package proto

import (
    "fmt"
)

func (t *TDisplayCurrency) Fill(dataRaw interface{}) error {
    data, ok := dataRaw.(map[string]interface{}) 
    if !ok {
        return fmt.Errorf("TDisplayCurrency can't convert %+v to map[string]interface{}", dataRaw)
    }


    if PRICE, ok := data["PRICE"].(string); ok {
        t.PRICE = PRICE
    } else {
        return fmt.Errorf("PRICE %+v can't be converted to string", data["PRICE"])
    }

    if LASTUPDATE, ok := data["LASTUPDATE"].(string); ok {
        t.LASTUPDATE = LASTUPDATE
    } else {
        return fmt.Errorf("LASTUPDATE %+v can't be converted to string", data["LASTUPDATE"])
    }

    if VOLUME24HOUR, ok := data["VOLUME24HOUR"].(string); ok {
        t.VOLUME24HOUR = VOLUME24HOUR
    } else {
        return fmt.Errorf("VOLUME24HOUR %+v can't be converted to string", data["VOLUME24HOUR"])
    }

    if VOLUME24HOURTO, ok := data["VOLUME24HOURTO"].(string); ok {
        t.VOLUME24HOURTO = VOLUME24HOURTO
    } else {
        return fmt.Errorf("VOLUME24HOURTO %+v can't be converted to string", data["VOLUME24HOURTO"])
    }

    if OPEN24HOUR, ok := data["OPEN24HOUR"].(string); ok {
            t.OPEN24HOUR = OPEN24HOUR
    } else {
        return fmt.Errorf("OPEN24HOUR %+v can't be converted to string", data["OPEN24HOUR"])
    }

    if HIGH24HOUR, ok :=  data["HIGH24HOUR"].(string); ok {
        t.HIGH24HOUR = HIGH24HOUR
    } else {
        return fmt.Errorf("HIGH24HOUR %+v can't be converted to string", data["HIGH24HOUR"])
    }

    if LOW24HOUR, ok := data["LOW24HOUR"].(string); ok {
        t.LOW24HOUR = LOW24HOUR
    } else {
        return fmt.Errorf("LOW24HOUR %+v can't be converted to string", data["LOW24HOUR"])
    }

    if CHANGE24HOUR, ok := data["CHANGE24HOUR"].(string); ok {
        t.CHANGE24HOUR = CHANGE24HOUR
    } else {
        return fmt.Errorf("CHANGE24HOUR %+v can't be converted to string", data["CHANGE24HOUR"])
    }

    if CHANGEPCT24HOUR, ok := data["CHANGEPCT24HOUR"].(string); ok {
        t.CHANGEPCT24HOUR = CHANGEPCT24HOUR
    } else {
        return fmt.Errorf("CHANGEPCT24HOUR %+v can't be converted to string", data["CHANGEPCT24HOUR"])
    }

    if SUPPLY, ok := data["SUPPLY"].(string); ok {
        t.SUPPLY = SUPPLY
    } else {
        return fmt.Errorf("SUPPLY %+v can't be converted to string", data["SUPPLY"])
    }

    if MKTCAP, ok := data["MKTCAP"].(string); ok {
        t.MKTCAP = MKTCAP
    } else {
        return fmt.Errorf("MKTCAP %+v can't be converted to string", data["MKTCAP"])
    }

    return nil
}
