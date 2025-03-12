package app

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/usecase/settings"
	"audirvana-scrobbler/app/usecase/trackinfo"
	"context"

	"github.com/samber/do"
)

type TrackInfoService struct {
	getTrackInfoList trackinfo.GetTrackInfoList
	saveTrackInfo    trackinfo.SaveTrackInfo
	deleteTrackInfo  trackinfo.DeleteTrackInfo
	scrobbleAll      trackinfo.ScrobbleAll
	getSessionKey    settings.GetLastFMSessionKey
	login            settings.LoginLastFM
	logout           settings.LogoutLastFM
}

func NewApp(i *do.Injector) (*TrackInfoService, error) {
	return &TrackInfoService{
		getTrackInfoList: do.MustInvoke[trackinfo.GetTrackInfoList](i),
		saveTrackInfo:    do.MustInvoke[trackinfo.SaveTrackInfo](i),
		deleteTrackInfo:  do.MustInvoke[trackinfo.DeleteTrackInfo](i),
		scrobbleAll:      do.MustInvoke[trackinfo.ScrobbleAll](i),
		getSessionKey:    do.MustInvoke[settings.GetLastFMSessionKey](i),
		login:            do.MustInvoke[settings.LoginLastFM](i),
		logout:           do.MustInvoke[settings.LogoutLastFM](i),
	}, nil
}

func (t *TrackInfoService) GetTrackInfoList(ctx context.Context) ([]bindings.TrackInfoResponse, *bindings.ErrorResponse) {
	return t.getTrackInfoList.Execute(ctx)
}

func (t *TrackInfoService) SaveTrackInfo(ctx context.Context, id int, form bindings.TrackInfoForm) *bindings.ErrorResponse {
	return t.saveTrackInfo.Execute(ctx, id, form)
}

func (t *TrackInfoService) DeleteTrackInfo(ctx context.Context, id int) *bindings.ErrorResponse {
	return t.deleteTrackInfo.Execute(ctx, id)
}

func (t *TrackInfoService) ScrobbleAll(ctx context.Context) *bindings.ErrorResponse {
	return t.scrobbleAll.Execute(ctx)
}

func (t *TrackInfoService) GetSessionKey(ctx context.Context, username, password string) *bindings.ErrorResponse {
	return t.getSessionKey.Execute(ctx, username, password)
}

func (t *TrackInfoService) Login(ctx context.Context) bool {
	return t.login.Execute(ctx)
}

func (t *TrackInfoService) Logout(ctx context.Context) *bindings.ErrorResponse {
	return t.logout.Execute(ctx)
}
