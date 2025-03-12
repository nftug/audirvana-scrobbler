package domain

type Config struct {
	APIKey              string `json:"apiKey"`
	APISecret           string `json:"apiSecret"`
	SessionKey          string `json:"sessionKey"`
	ScrobbleImmediately bool   `json:"scrobbleImmediately"`
	PositionThreshold   int    `json:"positionThreshold"`
}

var DefaultConfig = Config{
	ScrobbleImmediately: false,
	PositionThreshold:   80,
}
