package usecases

import "base/pkg/application/interfaces"

type createTrackingDataUsecase struct {
	repo   interfaces.ITrackingDataRepository
	logger interfaces.ILogger
}

func NewCreateTrackingDataUsecase(repo interfaces.ITrackingDataRepository, logger interfaces.ILogger) *createTrackingDataUsecase {
	return &createTrackingDataUsecase{
		repo,
		logger,
	}
}
