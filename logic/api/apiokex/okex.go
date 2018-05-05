package apiokex

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sidekick/strader/logic/api"
	"sidekick/strader/utils"
	"sort"
	"strings"
	"xframe/log"
)

type OkexApi struct {
}

func NewOkexApi() *OkexApi {
	return new(OkexApi)
}

func (this *OkexApi) FutureTrade(symbol, contractType, price, amount, ttype, matchPrice, leverRate string) ([]byte, bool) {
	var (
		params = make(map[string]string)
		header = make(http.Header)
	)
	//params and sign
	params["symbol"] = symbol
	params["contract_type"] = contractType
	params["price"] = price
	params["amount"] = amount
	params["type"] = ttype
	if matchPrice != "" {
		params["match_price"] = matchPrice
	}
	if leverRate != "" {
		params["lever_rate"] = leverRate
	}
	params["api_key"] = utils.GetOkexKey()
	signature := this.Sign(params, utils.GetOkexSecret())
	params["sign"] = signature
	//url
	addr := fmt.Sprintf("%s%s", BASE_URI, FTRADE_URI)
	//header
	header.Set("content-type", "application/x-www-form-urlencoded")
	//body
	val := url.Values{}
	for k, v := range params {
		val.Set(k, v)
	}
	log.DEBUG(val)
	body := bytes.NewBuffer([]byte(val.Encode()))
	res, err, statusCode := api.SendHttpPostRequest(addr, header, body, 10)
	//res, err := http.PostForm(addr, val)
	if err != nil {
		log.ERRORF("[api_okex]send ttrade to upstream error: %v", err)
		return nil, false
	}
	if statusCode != 200 {
		log.ERRORF("[api_okex]ttrade status code is not 200")
		return nil, false
	}
	/*defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.ERRORF("[api_okex]ttrade read response body error:%v", err)
		return nil, false
	}
	//TODO add polling and callback
	return result, true*/
	return res, true
}

func (this *OkexApi) FutureTradeCancel(symbol, contractType, orderId string) ([]byte, bool) {
	var (
		params = make(map[string]string)
		header = make(http.Header)
	)
	//params and sign
	params["symbol"] = symbol
	params["contract_type"] = contractType
	params["order_id"] = orderId
	params["api_key"] = utils.GetOkexKey()
	signature := this.Sign(params, utils.GetOkexSecret())
	params["sign"] = signature
	//url
	addr := fmt.Sprintf("%s%s", BASE_URI, FTRADE_CANCEL_URI)
	//header
	header.Set("content-type", "application/x-www-form-urlencoded")
	//body
	val := url.Values{}
	for k, v := range params {
		val.Set(k, v)
	}
	log.DEBUG(val)
	body := bytes.NewBuffer([]byte(val.Encode()))
	res, err, statusCode := api.SendHttpPostRequest(addr, header, body, 10)
	//res, err := http.PostForm(addr, val)
	if err != nil {
		log.ERRORF("[api_okex]send ttrade cancel to upstream error: %v", err)
		return nil, false
	}
	if statusCode != 200 {
		log.ERRORF("[api_okex]ttrade cancel status code is not 200")
		return nil, false
	}
	/*defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.ERRORF("[api_okex]ttrade cancel read response body error:%v", err)
		return nil, false
	}
	return result, true*/
	return res, true
}

func (this *OkexApi) Sign(params map[string]string, apiSecret string) string {
	var (
		keyLst     = make([]string, 0)
		sortParams string
	)
	for key, _ := range params {
		keyLst = append(keyLst, key)
	}
	sort.Strings(keyLst)
	for _, key := range keyLst {
		sortParams += key + "=" + params[key] + "&"
	}
	sortParams += "secret_key=" + apiSecret
	h := md5.New()
	io.WriteString(h, sortParams)
	sign := strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
	return sign
}
