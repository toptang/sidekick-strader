package resokex

type FutureTradeSingleRes struct {
	OrderId int64 `json:"order_id"`
	Result  bool  `json:"result"`
}

type FutureTradeRes struct {
	FutureTradeSingleRes
}

type FutureTradeCancelRes struct {
	FutureTradeSingleRes
}
