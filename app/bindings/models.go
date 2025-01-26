package bindings

type TrackInfoResponse struct {
	ID          int     `json:"id"`
	Artist      string  `json:"artist"`
	Album       string  `json:"album"`
	Track       string  `json:"track"`
	PlayedAt    string  `json:"playedAt"`
	ScrobbledAt *string `json:"scrobbledAt"`
}

type TrackInfoForm struct {
	Artist string `json:"artist" validate:"required"`
	Album  string `json:"album" validate:"required"`
	Track  string `json:"track" validate:"required"`
}

type NowPlayingResponse struct {
	AppName  string  `json:"appName"`
	Track    string  `json:"track"`
	Artist   string  `json:"artist"`
	Album    string  `json:"album"`
	Duration float64 `json:"duration"`
	Position float64 `json:"position"`
}
