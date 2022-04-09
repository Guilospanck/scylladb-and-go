package usecases

import "base/pkg/domain/dtos"

type IFindTrackingDataByPrimaryKeyUsecase interface {
	Perform(data dtos.TrackingDataPrimaryKeyDTO) (*dtos.TrackingDataDTO, error)
}
