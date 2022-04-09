package usecases

import "base/pkg/domain/dtos"

type IDeleteTrackingDataByPartitionKeyUsecase interface {
	Perform(data dtos.TrackingDataPartitionKeyDTO) error
}
