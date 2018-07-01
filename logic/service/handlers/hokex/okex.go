package hokex

import (
	"net/http"
	"sidekick/strader/logic/api/apiokex"
	"sidekick/strader/logic/service/requests/reqokex"
	"sidekick/strader/logic/service/svcerr"
	"xframe/log"
)

func FutureTrade(r *http.Request) (interface{}, int) {
	//parse input
	req, err := reqokex.ParseFutureTradeReq(r)
	if err != nil {
		log.ERRORF("[okex_future_trade]parse future trade request error: %v", err)
		return nil, svcerr.INPUT_ERROR
	}
	//init OkexApi
	okexCli := apiokex.NewOkexApi()
	//send to upstream
	res, ok := okexCli.FutureTrade(req.Symbol, req.ContractType, req.Price, req.Amount, req.Type, req.MatchPrice, req.LeverRate)
	if !ok {
		log.ERROR("[okex_future_trade]future trade error")
		return nil, svcerr.FUTURE_TRADE_ERROR
	}
	//return
	return res, svcerr.SUCCESS
}

func FutureCancel(r *http.Request) (interface{}, int) {
	//parse input
	req, err := reqokex.ParseFutureTradeCancelReq(r)
	if err != nil {
		log.ERRORF("[okex_future_trade_cancel]parse future trade cancel request error: %v", err)
		return nil, svcerr.INPUT_ERROR
	}
	//init OkexApi
	okexCli := apiokex.NewOkexApi()
	//send to upstream
	res, ok := okexCli.FutureTradeCancel(req.Symbol, req.ContractType, req.OrderId)
	if !ok {
		log.ERROR("[okex_future_trade_cancel]future trade error")
		return nil, svcerr.FUTURE_TRADE_ERROR
	}
	//return
	return res, svcerr.SUCCESS
}

func FutureDevolve(r *http.Request) (interface{}, int) {
	//parse input
	req, err := reqokex.ParseFutureDevolveReq(r)
	if err != nil {
		log.ERRORF("[okex_future_trade_devolve]parse future trade devolve request error: %v", err)
		return nil, svcerr.INPUT_ERROR
	}
	//init OkexApi
	okexCli := apiokex.NewOkexApi()
	//send to upstream
	res, ok := okexCli.FutureTradeDevolve(req.Symbol, req.Type, req.Amount)
	if !ok {
		log.ERROR("[okex_future_trade_devolve]future trade devolve error")
		return nil, svcerr.FUTURE_TRADE_ERROR
	}
	//return
	return res, svcerr.SUCCESS
}
