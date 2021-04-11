package proto

import (
    "encoding/json"
    "fmt"
)

func (t TServiceResponse) MarshalJSON() ([]byte, error) {
    rawFsymsTsyms := make(map[string]interface{})
    for fsym, raw := range t.RAW {
        rawFsymsTsyms[fsym] = raw.Currencies
    }

    displayFsymsTsyms := make(map[string]interface{})
    for fsym, display := range t.DISPLAY {
        displayFsymsTsyms[fsym] = display.Currencies
    }

    all := make(map[string]interface{})
    all["DISPLAY"] = displayFsymsTsyms
    all["RAW"] = rawFsymsTsyms

    return json.Marshal(all)
}

func (t *TServiceResponse) UnmarshalJSON(data []byte) error {
    rawDisplay := make(map[string]interface{})

    err := json.Unmarshal(data, &rawDisplay)
    if err != nil {
        return err
    }

    err = t.FillRaw(rawDisplay["RAW"])
    if err != nil {
        return err
    }

    err = t.FillDisplay(rawDisplay["DISPLAY"])
    if err != nil {
        return err
    }

    return nil
}

func (t *TServiceResponse) FillRaw(rawRaw interface{}) error {
    raw, ok := rawRaw.(map[string]interface{})
    if !ok {
        return fmt.Errorf("RAW part can't be parsed by json")
    }

    t.RAW = make(map[string]*TRaw)

    for fsym, tsymsRaw := range raw {
        var tRaw TRaw
        err := tRaw.FillTsyms(tsymsRaw)
        if err != nil {
            return err
        }

        t.RAW[fsym] = &tRaw
    }

    return nil
}

func (t *TServiceResponse) FillDisplay(displayRaw interface{}) error {
    display, ok := displayRaw.(map[string]interface{})
    if !ok {
        return fmt.Errorf("DISPLAY part can't be parsed by json")
    }

    t.DISPLAY = make(map[string]*TDisplay)

    for fsym, tsymsRaw := range display {
        var tDisplay TDisplay
        err := tDisplay.FillTsyms(tsymsRaw)
        if err != nil {
            return err
        }

        t.DISPLAY[fsym] = &tDisplay
    }

    return nil
}
