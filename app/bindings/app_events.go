package bindings

type AppEvent string

const (
	NotifyNowPlaying = AppEvent("NotifyNowPlaying")
)

var AppEvents = []AppEvent{NotifyNowPlaying}

func (a AppEvent) TSName() string { return string(a) }
