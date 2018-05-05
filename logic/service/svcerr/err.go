package svcerr

const (
	SUCCESS = iota

	INPUT_ERROR = iota + 3000
	FUTURE_TRADE_ERROR
	FUTURE_TRADE_CANCEL_ERROR
)

var (
	ErrMap = map[int]string{
		INPUT_ERROR:               "query input error",
		FUTURE_TRADE_ERROR:        "future trade error",
		FUTURE_TRADE_CANCEL_ERROR: "future trade cancel error",
	}
)
