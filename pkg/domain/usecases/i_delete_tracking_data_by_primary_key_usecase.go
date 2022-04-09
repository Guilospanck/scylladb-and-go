package usecases

import "base/pkg/domain/dtos"

type IDeleteTrackingDataByPrimaryKeyUsecase interface {
	Perform(data dtos.TrackingDataPrimaryKeyDTO) error
}
