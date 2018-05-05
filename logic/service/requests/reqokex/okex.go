package reqokex

import (
	"errors"
	"net/http"
	"xframe/utils"
)

type FutureTradeReq struct {
	Symbol       string `json:"symbol"`        //btc_usd ltc_usd eth_usd etc_usd bch_usd
	ContractType string `json:"contract_type"` //e.g this_week
	Price        string `json:"price"`
	Amount       string `json:"amount"`
	Type         string `json:"type"`        //1:开多 2:开空 3:平多 4:平空
	MatchPrice   string `json:"match_price"` //是否为对手价 0:不是 1:是 ,当取值为1时,price无效
	LeverRate    string `json:"lever_rate"`
}

func checkFutureTradeReq(req FutureTradeReq) bool {
	return !(req.Symbol == "" || req.ContractType == "" ||
		req.Price == "" || req.Amount == "" ||
		req.Type == "")
}

func ParseFutureTradeReq(r *http.Request) (req FutureTradeReq, err error) {
	err = utils.ParsePostRequest(r, &req)
	if err != nil {
		return
	}
	if !checkFutureTradeReq(req) {
		err = errors.New("check future_trade query empty")
	}
	return
}

//---------------------------
type FutureTradeCancelReq struct {
	Symbol       string `json:"symbol"`
	OrderId      string `json:"order_id"`
	ContractType string `json:"contract_type"`
}

func checkFutureTradeCancelReq(req FutureTradeCancelReq) bool {
	return !(req.Symbol == "" || req.OrderId == "" ||
		req.ContractType == "")
}

func ParseFutureTradeCancelReq(r *http.Request) (req FutureTradeCancelReq, err error) {
	err = utils.ParsePostRequest(r, &req)
	if err != nil {
		return
	}
	if !checkFutureTradeCancelReq(req) {
		err = errors.New("check future_trade_cancel query empty")
	}
	return
}
