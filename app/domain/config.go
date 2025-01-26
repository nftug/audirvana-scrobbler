package domain

type Config struct {
	APIKey              string `json:"apiKey"`
	APISecret           string `json:"apiSecret"`
	UserName            string `json:"userName"`
	Password            string `json:"password"`
	ScrobbleImmediately bool   `json:"scrobbleImmediately"`
	PositionThreshold   int    `json:"positionThreshold"`
}

var DefaultConfig = Config{
	ScrobbleImmediately: false,
	PositionThreshold:   80,
}
