package usecases

import "base/pkg/domain/dtos"

type IDeleteeTrackingDataByPrimaryKeyUsecase interface {
	Perform(data dtos.TrackingDataPrimaryKeyDTO) error
}
