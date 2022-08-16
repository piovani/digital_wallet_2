package domain

var (
	COINS = [...]string{
		"btc",
		"eth",
		"ada",
		"xrp",
		"doge",
	}
)

func ValidCoin(c string) bool {
	for _, coin := range COINS {
		if coin == c {
			return true
		}
	}

	return false
}
