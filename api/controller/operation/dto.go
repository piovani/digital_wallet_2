package operation

type Input struct {
	UserName string  `json:"user_name"`
	Coin     string  `json:"coin"`
	Value    float64 `json:"value"`
}
