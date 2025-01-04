package response

type TrackInfo struct {
	Id       string `json:"id"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Track    string `json:"track"`
	PlayedAt string `json:"playedAt"`
}
