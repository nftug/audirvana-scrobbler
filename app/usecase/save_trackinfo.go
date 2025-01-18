package usecase

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/samber/do"
)

type SaveTrackInfo interface {
	Execute(ctx context.Context, id string, form bindings.TrackInfoForm) *bindings.ErrorResponse
}

type saveTrackInfoImpl struct {
	repo domain.TrackInfoRepository
}

func NewSaveTrackInfo(i *do.Injector) (SaveTrackInfo, error) {
	return &saveTrackInfoImpl{
		repo: do.MustInvoke[domain.TrackInfoRepository](i),
	}, nil
}

func (s *saveTrackInfoImpl) Execute(ctx context.Context, id string, form bindings.TrackInfoForm) *bindings.ErrorResponse {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var errors []bindings.ErrorData
	if err := validator.New().Struct(form); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, bindings.ErrorData{
				Field:   err.Field(),
				Message: fmt.Sprintf("Field %s is %s", err.Field(), err.Tag()),
			})
		}
	}
	if len(errors) > 0 {
		return &bindings.ErrorResponse{Code: bindings.ValidationError, Data: errors}
	}

	if err := s.repo.Save(ctx, id, form); err != nil {
		return bindings.NewInternalError("Error while saving track info: %v", err)
	}

	return nil
}
