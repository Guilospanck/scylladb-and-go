package usecases

import "base/pkg/domain/dtos"

type IFindAllTrackingDataByPartitionKeyUsecase interface {
	Perform(data dtos.TrackingDataPartitionKeyDTO) ([]*dtos.TrackingDataDTO, error)
}
