package bindings

type TrackInfo struct {
	ID       string `json:"id"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Track    string `json:"track"`
	PlayedAt string `json:"playedAt"`
}
