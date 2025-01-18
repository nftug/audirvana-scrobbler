package domain

type Config struct {
	APIKey    string `env:"LASTFM_API_KEY"`
	APISecret string `env:"LASTFM_API_SECRET"`
	UserName  string `env:"LASTFM_USERNAME"`
	Password  string `env:"LASTFM_PASSWORD"`
}
