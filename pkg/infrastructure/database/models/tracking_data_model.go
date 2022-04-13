package models

import "github.com/Guilospanck/igocqlx/table"

type TrackingDataTable struct {
	Table *table.Table
}

func NewTrackingDataTable() *TrackingDataTable {
	trackingDataMetadata := table.Metadata{
		Name: "tracking_data",
		Columns: []string{
			"first_name", "last_name", "timestamp", "heat",
			"location", "speed", "telepathy_powers",
		},
		PartKey: []string{"first_name", "last_name"},
		SortKey: []string{"timestamp"},
	}

	trackingDataTable := table.New(trackingDataMetadata)

	return &TrackingDataTable{
		Table: trackingDataTable,
	}
}
