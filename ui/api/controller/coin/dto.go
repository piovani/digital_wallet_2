package coin

type Output struct {
	BtcUds float64 `json:"btc_uds"`
	BtcEur float64 `json:"btc_eur"`

	EthUds float64 `json:"eth_uds"`
	EthEur float64 `json:"eth_eur"`

	AdahUds float64 `json:"ada_uds"`
	AdaEur  float64 `json:"ada_eur"`

	XrphUds float64 `json:"xrp_uds"`
	XrpEur  float64 `json:"xrp_eur"`

	DogehUds float64 `json:"doge_uds"`
	DogeEur  float64 `json:"doge_eur"`
}

func NewOutput(p map[string]float64) *Output {
	return &Output{
		BtcUds: p["BTC_USD"],
		BtcEur: p["BTC_EUR"],

		EthUds: p["ETH_USD"],
		EthEur: p["ETH_EUR"],

		AdahUds: p["ADA_USD"],
		AdaEur:  p["ADA_EUR"],

		XrphUds: p["XRP_USD"],
		XrpEur:  p["XRP_EUR"],

		DogehUds: p["DOG_USD"],
		DogeEur:  p["DOG_EUR"],
	}
}
