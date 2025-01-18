package bindings

type TrackInfo struct {
	ID       string `json:"id"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Track    string `json:"track"`
	PlayedAt string `json:"playedAt"`
}

type TrackInfoForm struct {
	Artist string `json:"artist" validate:"required"`
	Album  string `json:"album" validate:"required"`
	Track  string `json:"track" validate:"required"`
}
