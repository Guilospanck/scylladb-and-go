package usecases

import "base/pkg/domain/dtos"

type ICreateTrackingDataUsecase interface {
	Perform(data dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error)
}
