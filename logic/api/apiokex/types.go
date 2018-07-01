package apiokex

var (
	BASE_URI           = "https://www.okex.com/api/v1"
	FTRADE_URI         = "/future_trade.do"
	FTRADE_CANCEL_URI  = "/future_cancel.do"
	FTRADE_DEVOLVE_URI = "/future_devolve.do"
)

type FutureResult struct {
	ErrorCode int64 `json:"error_code"`
	OrderId   int64 `json:"order_id"`
	Result    bool  `json:"result"`
}
