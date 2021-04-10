package proto

import (
    "encoding/json"
    "fmt"
)

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

    /*
    err = t.FillDisplay(rawDisplay["DISPLAY"])
    if err != nil {
        return err
    }
    */

    panic("ok")

    return nil
}

func (t *TServiceResponse) FillRaw(rawRaw interface{}) error {
    raw, ok := rawRaw.(map[string]interface{})
    if !ok {
        return fmt.Errorf("RAW part can't be parsed by json")
    }

    t.RAW = make(map[string]*TRaw)

    for fsym, tsymsRaw := range raw {
        err := t.FillRawTsyms(fsym, tsymsRaw)
        if err != nil {
            return err
        }
    }

    return nil
}

func (t *TServiceResponse) FillRawTsyms(fsym string, tsymsRaw interface{}) error {
    tsyms, ok := tsymsRaw.(map[string]interface{})
    if !ok {
        return fmt.Errorf("[RAW] fsym=%s tsyms=%+v can't be converted in map[string]interface{}", fsym, tsymsRaw)
    }

    //t.RAW[fsym] = make(TRaw)
    err := t.RAW[fsym].FillTsyms(tsyms)
    return err
}

/*
func (t *TServiceResponse) FillDisplay(display interface{}) error {
}
*/
