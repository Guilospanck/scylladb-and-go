package usecases

import "base/pkg/domain/dtos"

type IFindAllTrackingDataUsecase interface {
	Perform() ([]*dtos.TrackingDataDTO, error)
}
