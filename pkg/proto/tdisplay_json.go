package proto

import (
    "fmt"
)

func (t *TDisplay) FillTsyms(tsymsRaw interface{}) error {
    tsyms, ok := tsymsRaw.(map[string]interface{})
    if !ok {
        return fmt.Errorf("tsymsRaw %+v can't be converted to map[string]interface{}")
    }

    t.Currencies = make(map[string]*TDisplayCurrency)

    for tsym, currencyRaw := range tsyms {
        var currency TDisplayCurrency 
        err := currency.Fill(currencyRaw)
        if err != nil {
            return err
        }

        t.Currencies[tsym] =  &currency
    }

    return nil
}
