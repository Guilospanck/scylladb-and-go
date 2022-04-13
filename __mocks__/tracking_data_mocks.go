package mocks

import (
	"base/pkg/domain/dtos"
	"base/pkg/infrastructure/database/entities"
	"time"
)

var (
	timeLayout   = "2006-01-02 15:04:05 -0700 MST"
	timestamp, _ = time.Parse(timeLayout, "2017-11-11 12:05:00 +0000 UTC")

	CompleteDataEntity = entities.TrackingDataEntity{
		FirstName:       "Guilherme",
		LastName:        "Rodrigues",
		Timestamp:       timestamp,
		Location:        "Brazil",
		Speed:           50.5,
		Heat:            40,
		TelepathyPowers: 10,
	}

	CompleteDataDTO = dtos.TrackingDataDTO{
		FirstName:       "Guilherme",
		LastName:        "Rodrigues",
		Timestamp:       timestamp.String(),
		Location:        "Brazil",
		Speed:           50.5,
		Heat:            40,
		TelepathyPowers: 10,
	}

	PrimaryKeyDataDTO = dtos.TrackingDataPrimaryKeyDTO{
		FirstName: "Jim",
		LastName:  "Jeffries",
		Timestamp: timestamp.String(),
	}

	PartitionKeyDataDTO = dtos.TrackingDataPartitionKeyDTO{
		FirstName: "Jim",
		LastName:  "Jeffries",
	}
)

/* Delete */
// timeLayout := "2006-01-02 15:04:05 -0700 MST"
// timestamp, err := time.Parse(timeLayout, "2017-11-11 11:05:00 +0000 UTC")
// if err != nil {
// 	logger.Info(err.Error())
// }

// dataToBeDeleted := entities.TrackingDataEntity{
// 	FirstName: "Bob",
// 	LastName:  "Loblaw",
// 	Timestamp: timestamp,
// }
// querybuilder.Delete(&dataToBeDeleted)

/* Delete All */
// dataToBeDeleted := entities.TrackingDataEntity{
//   FirstName: "Bob",
//   LastName:  "Loblaw",
// }
// querybuilder.DeleteAllFromPartitioningKey(&dataToBeDeleted)

/* Select */
// dataToBeSearched := entities.TrackingDataEntity{
// 	FirstName: "Bob",
// 	LastName:  "Loblaw",
// }
// ShowSelect[entities.TrackingDataEntity](querybuilder, logger, &dataToBeSearched)

/* Get */
// timeLayout := "2006-01-02 15:04:05 -0700 MST"
// timestamp, err := time.Parse(timeLayout, "2017-11-11 10:05:00 +0000 UTC")
// if err != nil {
// 	logger.Info(err.Error())
// }
// dataToBeSearched := entities.TrackingDataEntity{
// 	FirstName: "Bob",
// 	LastName:  "Loblaw",
// 	Timestamp: timestamp,
// }
// ShowGet[entities.TrackingDataEntity](querybuilder, logger, &dataToBeSearched)

/* Select All */
// ShowValuesSelectAll[entities.TrackingDataEntity](querybuilder, logger)
