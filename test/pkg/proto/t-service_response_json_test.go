//testing:
//go test  ./test/pkg/proto/t-service_response_json_test.go --run=TestUnmarshalJSON

package prototest

import (
    "testing"
    "encoding/json"
    "github.com/smart--petea/test2/pkg/proto"
)

func TestMarshal(t *testing.T) {
    var encodedTServiceResponse []byte = []byte(`{"RAW":{"BTC":{"USD":{"TYPE":"5","MARKET":"CCCAGG","FROMSYMBOL":"BTC","TOSYMBOL":"USD","FLAGS":"2050","PRICE":59692.06,"LASTUPDATE":1618082612,"MEDIAN":59669.75,"LASTVOLUME":0.00080086,"LASTVOLUMETO":47.7786989464,"LASTTRADEID":"154219127","VOLUMEDAY":32185.116040684898,"VOLUMEDAYTO":1935264417.9284399,"VOLUME24HOUR":34972.740182819995,"VOLUME24HOURTO":2097598475.2247777,"OPENDAY":58106.74,"HIGHDAY":61193.5,"LOWDAY":57885.43,"OPEN24HOUR":58365.92,"HIGH24HOUR":61222.49,"LOW24HOUR":57866.04,"LASTMARKET":"Coinbase","VOLUMEHOUR":323.2267347100169,"VOLUMEHOURTO":19325003.241563436,"OPENHOUR":59986.87,"HIGHHOUR":59987.59,"LOWHOUR":59645.27,"TOPTIERVOLUME24HOUR":34972.740182819995,"TOPTIERVOLUME24HOURTO":2097598475.2247777,"CHANGE24HOUR":1326.1399999999994,"CHANGEPCT24HOUR":2.272113589574189,"CHANGEDAY":1585.3199999999997,"CHANGEPCTDAY":2.728289351631153,"CHANGEHOUR":-294.81000000000495,"CHANGEPCTHOUR":-0.49145754729327423,"CONVERSIONTYPE":"direct","CONVERSIONSYMBOL":"","SUPPLY":18679087,"MKTCAP":1114993181949.22,"MKTCAPPENALTY":0,"TOTALVOLUME24H":221033.3554319794,"TOTALVOLUME24HTO":13203939884.314514,"TOTALTOPTIERVOLUME24H":220856.3484081226,"TOTALTOPTIERVOLUME24HTO":13193373970.426033,"IMAGEURL":"/media/37746251/btc.png"},"EUR":{"TYPE":"5","MARKET":"CCCAGG","FROMSYMBOL":"BTC","TOSYMBOL":"EUR","FLAGS":"2052","PRICE":49897.2,"LASTUPDATE":1618082610,"MEDIAN":50042.693,"LASTVOLUME":0.00022,"LASTVOLUMETO":11.04972,"LASTTRADEID":"663459243","VOLUMEDAY":10066.615922835936,"VOLUMEDAYTO":505470335.8232153,"VOLUME24HOUR":10623.489622679997,"VOLUME24HOURTO":532781290.02157354,"OPENDAY":48922.72,"HIGHDAY":50860.52,"LOWDAY":48755.3,"OPEN24HOUR":49095.49,"HIGH24HOUR":50874.46,"LOW24HOUR":48733.06,"LASTMARKET":"Bitfinex","VOLUMEHOUR":99.80564314999916,"VOLUMEHOURTO":4986511.234716723,"OPENHOUR":50108.45,"HIGHHOUR":50108.68,"LOWHOUR":49866.44,"TOPTIERVOLUME24HOUR":10623.489622679997,"TOPTIERVOLUME24HOURTO":532781290.02157354,"CHANGE24HOUR":801.7099999999991,"CHANGEPCT24HOUR":1.6329605835485075,"CHANGEDAY":974.4799999999959,"CHANGEPCTDAY":1.9918761671468714,"CHANGEHOUR":-211.25,"CHANGEPCTHOUR":-0.42158558087508197,"CONVERSIONTYPE":"direct","CONVERSIONSYMBOL":"","SUPPLY":18679087,"MKTCAP":932034139856.3999,"MKTCAPPENALTY":0,"TOTALVOLUME24H":221033.3554319794,"TOTALVOLUME24HTO":11031644446.281347,"TOTALTOPTIERVOLUME24H":220856.3484081226,"TOTALTOPTIERVOLUME24HTO":11022812291.41056,"IMAGEURL":"/media/37746251/btc.png"}}},"DISPLAY":{"BTC":{"USD":{"FROMSYMBOL":"Ƀ","TOSYMBOL":"$","MARKET":"CryptoCompare Index","PRICE":"$ 59,692.1","LASTUPDATE":"Just now","LASTVOLUME":"Ƀ 0.0008009","LASTVOLUMETO":"$ 47.78","LASTTRADEID":"154219127","VOLUMEDAY":"Ƀ 32,185.1","VOLUMEDAYTO":"$ 1,935,264,417.9","VOLUME24HOUR":"Ƀ 34,972.7","VOLUME24HOURTO":"$ 2,097,598,475.2","OPENDAY":"$ 58,106.7","HIGHDAY":"$ 61,193.5","LOWDAY":"$ 57,885.4","OPEN24HOUR":"$ 58,365.9","HIGH24HOUR":"$ 61,222.5","LOW24HOUR":"$ 57,866.0","LASTMARKET":"Coinbase","VOLUMEHOUR":"Ƀ 323.23","VOLUMEHOURTO":"$ 19,325,003.2","OPENHOUR":"$ 59,986.9","HIGHHOUR":"$ 59,987.6","LOWHOUR":"$ 59,645.3","TOPTIERVOLUME24HOUR":"Ƀ 34,972.7","TOPTIERVOLUME24HOURTO":"$ 2,097,598,475.2","CHANGE24HOUR":"$ 1,326.14","CHANGEPCT24HOUR":"2.27","CHANGEDAY":"$ 1,585.32","CHANGEPCTDAY":"2.73","CHANGEHOUR":"$ -294.81","CHANGEPCTHOUR":"-0.49","CONVERSIONTYPE":"direct","CONVERSIONSYMBOL":"","SUPPLY":"Ƀ 18,679,087.0","MKTCAP":"$ 1,114.99 B","MKTCAPPENALTY":"0 %","TOTALVOLUME24H":"Ƀ 221.03 K","TOTALVOLUME24HTO":"$ 13.20 B","TOTALTOPTIERVOLUME24H":"Ƀ 220.86 K","TOTALTOPTIERVOLUME24HTO":"$ 13.19 B","IMAGEURL":"/media/37746251/btc.png"},"EUR":{"FROMSYMBOL":"Ƀ","TOSYMBOL":"€","MARKET":"CryptoCompare Index","PRICE":"€ 49,897.2","LASTUPDATE":"Just now","LASTVOLUME":"Ƀ 0.0002200","LASTVOLUMETO":"€ 11.05","LASTTRADEID":"663459243","VOLUMEDAY":"Ƀ 10,066.6","VOLUMEDAYTO":"€ 505,470,335.8","VOLUME24HOUR":"Ƀ 10,623.5","VOLUME24HOURTO":"€ 532,781,290.0","OPENDAY":"€ 48,922.7","HIGHDAY":"€ 50,860.5","LOWDAY":"€ 48,755.3","OPEN24HOUR":"€ 49,095.5","HIGH24HOUR":"€ 50,874.5","LOW24HOUR":"€ 48,733.1","LASTMARKET":"Bitfinex","VOLUMEHOUR":"Ƀ 99.81","VOLUMEHOURTO":"€ 4,986,511.2","OPENHOUR":"€ 50,108.5","HIGHHOUR":"€ 50,108.7","LOWHOUR":"€ 49,866.4","TOPTIERVOLUME24HOUR":"Ƀ 10,623.5","TOPTIERVOLUME24HOURTO":"€ 532,781,290.0","CHANGE24HOUR":"€ 801.71","CHANGEPCT24HOUR":"1.63","CHANGEDAY":"€ 974.48","CHANGEPCTDAY":"1.99","CHANGEHOUR":"€ -211.25","CHANGEPCTHOUR":"-0.42","CONVERSIONTYPE":"direct","CONVERSIONSYMBOL":"","SUPPLY":"Ƀ 18,679,087.0","MKTCAP":"€ 932.03 B","MKTCAPPENALTY":"0 %","TOTALVOLUME24H":"Ƀ 221.03 K","TOTALVOLUME24HTO":"€ 11.03 B","TOTALTOPTIERVOLUME24H":"Ƀ 220.86 K","TOTALTOPTIERVOLUME24HTO":"€ 11.02 B","IMAGEURL":"/media/37746251/btc.png"}}}}`)
    var r proto.TServiceResponse

    err := json.Unmarshal(encodedTServiceResponse, &r)
    if err != nil {
        t.Error(err)
    }

    data, err := json.Marshal(&r)
    if err != nil {
        t.Error(err)
    }
}

func TestUnmarshal(t *testing.T) {
    var encodedTServiceResponse []byte = []byte(`{"RAW":{"BTC":{"USD":{"TYPE":"5","MARKET":"CCCAGG","FROMSYMBOL":"BTC","TOSYMBOL":"USD","FLAGS":"2050","PRICE":59692.06,"LASTUPDATE":1618082612,"MEDIAN":59669.75,"LASTVOLUME":0.00080086,"LASTVOLUMETO":47.7786989464,"LASTTRADEID":"154219127","VOLUMEDAY":32185.116040684898,"VOLUMEDAYTO":1935264417.9284399,"VOLUME24HOUR":34972.740182819995,"VOLUME24HOURTO":2097598475.2247777,"OPENDAY":58106.74,"HIGHDAY":61193.5,"LOWDAY":57885.43,"OPEN24HOUR":58365.92,"HIGH24HOUR":61222.49,"LOW24HOUR":57866.04,"LASTMARKET":"Coinbase","VOLUMEHOUR":323.2267347100169,"VOLUMEHOURTO":19325003.241563436,"OPENHOUR":59986.87,"HIGHHOUR":59987.59,"LOWHOUR":59645.27,"TOPTIERVOLUME24HOUR":34972.740182819995,"TOPTIERVOLUME24HOURTO":2097598475.2247777,"CHANGE24HOUR":1326.1399999999994,"CHANGEPCT24HOUR":2.272113589574189,"CHANGEDAY":1585.3199999999997,"CHANGEPCTDAY":2.728289351631153,"CHANGEHOUR":-294.81000000000495,"CHANGEPCTHOUR":-0.49145754729327423,"CONVERSIONTYPE":"direct","CONVERSIONSYMBOL":"","SUPPLY":18679087,"MKTCAP":1114993181949.22,"MKTCAPPENALTY":0,"TOTALVOLUME24H":221033.3554319794,"TOTALVOLUME24HTO":13203939884.314514,"TOTALTOPTIERVOLUME24H":220856.3484081226,"TOTALTOPTIERVOLUME24HTO":13193373970.426033,"IMAGEURL":"/media/37746251/btc.png"},"EUR":{"TYPE":"5","MARKET":"CCCAGG","FROMSYMBOL":"BTC","TOSYMBOL":"EUR","FLAGS":"2052","PRICE":49897.2,"LASTUPDATE":1618082610,"MEDIAN":50042.693,"LASTVOLUME":0.00022,"LASTVOLUMETO":11.04972,"LASTTRADEID":"663459243","VOLUMEDAY":10066.615922835936,"VOLUMEDAYTO":505470335.8232153,"VOLUME24HOUR":10623.489622679997,"VOLUME24HOURTO":532781290.02157354,"OPENDAY":48922.72,"HIGHDAY":50860.52,"LOWDAY":48755.3,"OPEN24HOUR":49095.49,"HIGH24HOUR":50874.46,"LOW24HOUR":48733.06,"LASTMARKET":"Bitfinex","VOLUMEHOUR":99.80564314999916,"VOLUMEHOURTO":4986511.234716723,"OPENHOUR":50108.45,"HIGHHOUR":50108.68,"LOWHOUR":49866.44,"TOPTIERVOLUME24HOUR":10623.489622679997,"TOPTIERVOLUME24HOURTO":532781290.02157354,"CHANGE24HOUR":801.7099999999991,"CHANGEPCT24HOUR":1.6329605835485075,"CHANGEDAY":974.4799999999959,"CHANGEPCTDAY":1.9918761671468714,"CHANGEHOUR":-211.25,"CHANGEPCTHOUR":-0.42158558087508197,"CONVERSIONTYPE":"direct","CONVERSIONSYMBOL":"","SUPPLY":18679087,"MKTCAP":932034139856.3999,"MKTCAPPENALTY":0,"TOTALVOLUME24H":221033.3554319794,"TOTALVOLUME24HTO":11031644446.281347,"TOTALTOPTIERVOLUME24H":220856.3484081226,"TOTALTOPTIERVOLUME24HTO":11022812291.41056,"IMAGEURL":"/media/37746251/btc.png"}}},"DISPLAY":{"BTC":{"USD":{"FROMSYMBOL":"Ƀ","TOSYMBOL":"$","MARKET":"CryptoCompare Index","PRICE":"$ 59,692.1","LASTUPDATE":"Just now","LASTVOLUME":"Ƀ 0.0008009","LASTVOLUMETO":"$ 47.78","LASTTRADEID":"154219127","VOLUMEDAY":"Ƀ 32,185.1","VOLUMEDAYTO":"$ 1,935,264,417.9","VOLUME24HOUR":"Ƀ 34,972.7","VOLUME24HOURTO":"$ 2,097,598,475.2","OPENDAY":"$ 58,106.7","HIGHDAY":"$ 61,193.5","LOWDAY":"$ 57,885.4","OPEN24HOUR":"$ 58,365.9","HIGH24HOUR":"$ 61,222.5","LOW24HOUR":"$ 57,866.0","LASTMARKET":"Coinbase","VOLUMEHOUR":"Ƀ 323.23","VOLUMEHOURTO":"$ 19,325,003.2","OPENHOUR":"$ 59,986.9","HIGHHOUR":"$ 59,987.6","LOWHOUR":"$ 59,645.3","TOPTIERVOLUME24HOUR":"Ƀ 34,972.7","TOPTIERVOLUME24HOURTO":"$ 2,097,598,475.2","CHANGE24HOUR":"$ 1,326.14","CHANGEPCT24HOUR":"2.27","CHANGEDAY":"$ 1,585.32","CHANGEPCTDAY":"2.73","CHANGEHOUR":"$ -294.81","CHANGEPCTHOUR":"-0.49","CONVERSIONTYPE":"direct","CONVERSIONSYMBOL":"","SUPPLY":"Ƀ 18,679,087.0","MKTCAP":"$ 1,114.99 B","MKTCAPPENALTY":"0 %","TOTALVOLUME24H":"Ƀ 221.03 K","TOTALVOLUME24HTO":"$ 13.20 B","TOTALTOPTIERVOLUME24H":"Ƀ 220.86 K","TOTALTOPTIERVOLUME24HTO":"$ 13.19 B","IMAGEURL":"/media/37746251/btc.png"},"EUR":{"FROMSYMBOL":"Ƀ","TOSYMBOL":"€","MARKET":"CryptoCompare Index","PRICE":"€ 49,897.2","LASTUPDATE":"Just now","LASTVOLUME":"Ƀ 0.0002200","LASTVOLUMETO":"€ 11.05","LASTTRADEID":"663459243","VOLUMEDAY":"Ƀ 10,066.6","VOLUMEDAYTO":"€ 505,470,335.8","VOLUME24HOUR":"Ƀ 10,623.5","VOLUME24HOURTO":"€ 532,781,290.0","OPENDAY":"€ 48,922.7","HIGHDAY":"€ 50,860.5","LOWDAY":"€ 48,755.3","OPEN24HOUR":"€ 49,095.5","HIGH24HOUR":"€ 50,874.5","LOW24HOUR":"€ 48,733.1","LASTMARKET":"Bitfinex","VOLUMEHOUR":"Ƀ 99.81","VOLUMEHOURTO":"€ 4,986,511.2","OPENHOUR":"€ 50,108.5","HIGHHOUR":"€ 50,108.7","LOWHOUR":"€ 49,866.4","TOPTIERVOLUME24HOUR":"Ƀ 10,623.5","TOPTIERVOLUME24HOURTO":"€ 532,781,290.0","CHANGE24HOUR":"€ 801.71","CHANGEPCT24HOUR":"1.63","CHANGEDAY":"€ 974.48","CHANGEPCTDAY":"1.99","CHANGEHOUR":"€ -211.25","CHANGEPCTHOUR":"-0.42","CONVERSIONTYPE":"direct","CONVERSIONSYMBOL":"","SUPPLY":"Ƀ 18,679,087.0","MKTCAP":"€ 932.03 B","MKTCAPPENALTY":"0 %","TOTALVOLUME24H":"Ƀ 221.03 K","TOTALVOLUME24HTO":"€ 11.03 B","TOTALTOPTIERVOLUME24H":"Ƀ 220.86 K","TOTALTOPTIERVOLUME24HTO":"€ 11.02 B","IMAGEURL":"/media/37746251/btc.png"}}}}`)
    var r proto.TServiceResponse

    err := json.Unmarshal(encodedTServiceResponse, &r)
    if err != nil {
        t.Error(err)
    }

    if len(r.RAW) != 1 {
        t.Error("len(r.RAW) should be 1")
    }

    if _, ok := r.RAW["BTC"]; !ok {
        t.Error("wrong r.RAW['BTC']")
    }

    if len(r.DISPLAY) != 1 {
        t.Error("len(r.DISPLAY) should be 1")
    }

    if _, ok := r.DISPLAY["BTC"]; !ok {
        t.Error("wrong r.DISPLAY['BTC']")
    }
}

func TestTRawCurrencyFill(t *testing.T) {
	var PRICE float64 = 59692.06
	var LASTUPDATE int64 = 1618082612
	var VOLUME24HOUR float64 = 34972.740182819995
	var VOLUME24HOURTO float64 = 2097598475.2247777
	var OPEN24HOUR float64 = 58365.92
	var HIGH24HOUR float64 = 61222.49
	var LOW24HOUR float64 = 57866.04
	var CHANGE24HOUR float64 = 1326.1399999999994
	var CHANGEPCT24HOUR float64 = 2.272113589574189
	var SUPPLY int64 = 18679087
	var MKTCAP float64 = 1114993181949.22

    m := make(map[string]interface{})
	m["PRICE"] = PRICE
	m["LASTUPDATE"] = LASTUPDATE
	m["VOLUME24HOUR"] = VOLUME24HOUR
	m["VOLUME24HOURTO"] = VOLUME24HOURTO
	m["OPEN24HOUR"] = OPEN24HOUR
	m["HIGH24HOUR"] = HIGH24HOUR
	m["LOW24HOUR"] = LOW24HOUR
	m["CHANGE24HOUR"] = CHANGE24HOUR
	m["CHANGEPCT24HOUR"] = CHANGEPCT24HOUR
	m["SUPPLY"] = SUPPLY
	m["MKTCAP"] = MKTCAP

    var tRawCurrency proto.TRawCurrency
    err := tRawCurrency.Fill(m)
    if err != nil {
        t.Error(err)
    }

	if tRawCurrency.PRICE != PRICE {
        t.Error("PRICE not equal")
    }
	if tRawCurrency.LASTUPDATE != LASTUPDATE {
        t.Error("LASTUPDATE not equal")
    }

	if tRawCurrency.VOLUME24HOUR != VOLUME24HOUR {
        t.Error("VOLUME24HOUR not equal")
    }

	if tRawCurrency.VOLUME24HOURTO != VOLUME24HOURTO {
        t.Error("VOLUME24HOURTO not equal")
    }

	if tRawCurrency.OPEN24HOUR != OPEN24HOUR {
        t.Error("OPEN24HOUR not equal")
    }

	if tRawCurrency.HIGH24HOUR != HIGH24HOUR {
        t.Error("HIGH24HOUR not equal")
    }

	if tRawCurrency.LOW24HOUR != LOW24HOUR {
        t.Error("LOW24HOUR not equal")
    } 

	if tRawCurrency.CHANGE24HOUR != CHANGE24HOUR {
        t.Error("CHANGE24HOUR not equal")
    }

	if tRawCurrency.CHANGEPCT24HOUR != CHANGEPCT24HOUR {
        t.Error("CHANGEPCT24HOUR not equal")
    }

	if tRawCurrency.SUPPLY != SUPPLY {
        t.Error("SUPPLY not equal")
    }

	if tRawCurrency.MKTCAP != MKTCAP {
        t.Error("MKTCAP not equal")
    }
}

func TestTDisplayCurrencyFill(t *testing.T) {
	var PRICE string = "59692.06"
	var LASTUPDATE string = "1618082612"
	var VOLUME24HOUR string = "34972.740182819995"
	var VOLUME24HOURTO string = "2097598475.2247777"
	var OPEN24HOUR string = "58365.92"
	var HIGH24HOUR string = "61222.49"
	var LOW24HOUR string = "57866.04"
	var CHANGE24HOUR string = "1326.1399999999994"
	var CHANGEPCT24HOUR string = "2.272113589574189"
	var SUPPLY string = "18679087"
	var MKTCAP string = "1114993181949.22"

    m := make(map[string]interface{})
	m["PRICE"] = PRICE
	m["LASTUPDATE"] = LASTUPDATE
	m["VOLUME24HOUR"] = VOLUME24HOUR
	m["VOLUME24HOURTO"] = VOLUME24HOURTO
	m["OPEN24HOUR"] = OPEN24HOUR
	m["HIGH24HOUR"] = HIGH24HOUR
	m["LOW24HOUR"] = LOW24HOUR
	m["CHANGE24HOUR"] = CHANGE24HOUR
	m["CHANGEPCT24HOUR"] = CHANGEPCT24HOUR
	m["SUPPLY"] = SUPPLY
	m["MKTCAP"] = MKTCAP

    var tDisplayCurrency proto.TDisplayCurrency
    err := tDisplayCurrency.Fill(m)
    if err != nil {
        t.Error(err)
    }

	if tDisplayCurrency.PRICE != PRICE {
        t.Error("PRICE not equal")
    }
	if tDisplayCurrency.LASTUPDATE != LASTUPDATE {
        t.Error("LASTUPDATE not equal")
    }

	if tDisplayCurrency.VOLUME24HOUR != VOLUME24HOUR {
        t.Error("VOLUME24HOUR not equal")
    }

	if tDisplayCurrency.VOLUME24HOURTO != VOLUME24HOURTO {
        t.Error("VOLUME24HOURTO not equal")
    }

	if tDisplayCurrency.OPEN24HOUR != OPEN24HOUR {
        t.Error("OPEN24HOUR not equal")
    }

	if tDisplayCurrency.HIGH24HOUR != HIGH24HOUR {
        t.Error("HIGH24HOUR not equal")
    }

	if tDisplayCurrency.LOW24HOUR != LOW24HOUR {
        t.Error("LOW24HOUR not equal")
    } 

	if tDisplayCurrency.CHANGE24HOUR != CHANGE24HOUR {
        t.Error("CHANGE24HOUR not equal")
    }

	if tDisplayCurrency.CHANGEPCT24HOUR != CHANGEPCT24HOUR {
        t.Error("CHANGEPCT24HOUR not equal")
    }

	if tDisplayCurrency.SUPPLY != SUPPLY {
        t.Error("SUPPLY not equal")
    }

	if tDisplayCurrency.MKTCAP != MKTCAP {
        t.Error("MKTCAP not equal")
    }
}

/*
todo
func TestTServiceResponseFillDisplay(t *testing.T) {
    var tServiceResponse proto.TServiceResponse

    m := make(map[string]interface{})
    tsyms := getTsyms()
    m["BTC"] = tsyms 
    m["LINK"] = tsyms 

    if err := tServiceResponse.FillRaw(m); err != nil {
        t.Error(err)
    }
}
*/

func TestTServiceResponseFillDisplay(t *testing.T) {
    var tServiceResponse proto.TServiceResponse

    m := make(map[string]interface{})
    tsyms := getDisplayTsyms()
    m["BTC"] = tsyms 
    m["LINK"] = tsyms 

    if err := tServiceResponse.FillDisplay(m); err != nil {
        t.Error(err)
    }

    if len(tServiceResponse.DISPLAY) != 2 {
        t.Errorf("There should be 2 fsyms %+v", tServiceResponse.DISPLAY)
    }

    if _, ok := tServiceResponse.DISPLAY["BTC"]; !ok {
        t.Error("BTC should be filled in")
    }

    if _, ok := tServiceResponse.DISPLAY["LINK"]; !ok {
        t.Error("LINK should be filled in")
    }

    if len(tServiceResponse.DISPLAY["BTC"].Currencies) != 2 {
        t.Errorf("DISPLAY['BTC'].Currencies should have two currencies %+v", tServiceResponse.DISPLAY["BTC"])
    }
}

func TestTServiceResponseFillRaw(t *testing.T) {
    var tServiceResponse proto.TServiceResponse

    m := make(map[string]interface{})
    tsyms := getRawTsyms()
    m["BTC"] = tsyms 
    m["LINK"] = tsyms 

    if err := tServiceResponse.FillRaw(m); err != nil {
        t.Error(err)
    }

    if len(tServiceResponse.RAW) != 2 {
        t.Errorf("There should be 2 fsyms %+v", tServiceResponse.RAW)
    }

    if _, ok := tServiceResponse.RAW["BTC"]; !ok {
        t.Error("BTC should be filled in")
    }

    if _, ok := tServiceResponse.RAW["LINK"]; !ok {
        t.Error("LINK should be filled in")
    }

    if len(tServiceResponse.RAW["BTC"].Currencies) != 2 {
        t.Errorf("RAW['BTC'].Currencies should have two currencies %+v", tServiceResponse.RAW["BTC"])
    }
}

func TestTRawFillTsyms(t *testing.T) {
    var tRaw proto.TRaw

    m := make(map[string]interface{})
    currencyMap := getRawCurrencyMap()
    m["USD"] = currencyMap
    m["EUR"] = currencyMap

    err := tRaw.FillTsyms(m)
    if err != nil {
        t.Error(err)
    }

    if len(tRaw.Currencies) != 2 {
        t.Error("2 Currencies should be filled")
    }

    if _, ok := tRaw.Currencies["USD"]; !ok {
        t.Error("'USD' should be filled in Currencies")
    }

    if _, ok := tRaw.Currencies["EUR"]; !ok {
        t.Error("'EUR' should be filled in Currencies")
    }

    if tRaw.Currencies["USD"].PRICE !=  currencyMap["PRICE"] {
        t.Error("['USD']['PRICE'] should be filled in Currencies")
    }
}

func TestTDisplayFillTsyms(t *testing.T) {
    var tDisplay proto.TDisplay

    m := make(map[string]interface{})
    currencyMap := getDisplayCurrencyMap()
    m["USD"] = currencyMap
    m["EUR"] = currencyMap

    err := tDisplay.FillTsyms(m)
    if err != nil {
        t.Error(err)
    }

    if len(tDisplay.Currencies) != 2 {
        t.Error("2 Currencies should be filled")
    }

    if _, ok := tDisplay.Currencies["USD"]; !ok {
        t.Error("'USD' should be filled in Currencies")
    }

    if _, ok := tDisplay.Currencies["EUR"]; !ok {
        t.Error("'EUR' should be filled in Currencies")
    }

    if tDisplay.Currencies["USD"].PRICE !=  currencyMap["PRICE"] {
        t.Error("['USD']['PRICE'] should be filled in Currencies")
    }
}

func getRawTsyms() map[string]interface{} {
    m := make(map[string]interface{})

    currencyMap := getRawCurrencyMap()
    m["USD"] = currencyMap
    m["EUR"] = currencyMap

    return m
}

func getDisplayTsyms() map[string]interface{} {
    m := make(map[string]interface{})

    currencyMap := getDisplayCurrencyMap()
    m["USD"] = currencyMap
    m["EUR"] = currencyMap

    return m
}

func getRawCurrencyMap() map[string]interface{} {
    m := make(map[string]interface{})
	m["PRICE"] = float64(59692.06)
	m["LASTUPDATE"] = int64(1618082612)
	m["VOLUME24HOUR"] = float64(34972.740182819995)
	m["VOLUME24HOURTO"] = float64(2097598475.2247777)
	m["OPEN24HOUR"] = float64(58365.92)
	m["HIGH24HOUR"] = float64(61222.49)
	m["LOW24HOUR"] = float64(57866.04)
	m["CHANGE24HOUR"] = float64(1326.1399999999994)
	m["CHANGEPCT24HOUR"] = float64(2.272113589574189)
	m["SUPPLY"] = int64(18679087)
	m["MKTCAP"] = float64(1114993181949.22)

    return m
}

func getDisplayCurrencyMap() map[string]interface{} {
    m := make(map[string]interface{})
	m["PRICE"] = "59692.06"
	m["LASTUPDATE"] = "1618082612"
	m["VOLUME24HOUR"] = "34972.740182819995"
	m["VOLUME24HOURTO"] = "2097598475.2247777"
	m["OPEN24HOUR"] = "58365.92"
	m["HIGH24HOUR"] = "61222.49"
	m["LOW24HOUR"] = "57866.04"
	m["CHANGE24HOUR"] = "1326.1399999999994"
	m["CHANGEPCT24HOUR"] = "2.272113589574189"
	m["SUPPLY"] = "18679087"
	m["MKTCAP"] = "1114993181949.22"

    return m
}
