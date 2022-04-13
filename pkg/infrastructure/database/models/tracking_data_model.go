package models

import (
	igocqlxtable "github.com/Guilospanck/igocqlx/table"
	"github.com/scylladb/gocqlx/v2/table"
)

type TrackingDataTable struct {
	Table igocqlxtable.ITable
}

func NewTrackingDataTable() *TrackingDataTable {
	trackingDataMetadata := igocqlxtable.Metadata{
		M: &table.Metadata{
			Name: "tracking_data",
			Columns: []string{
				"first_name", "last_name", "timestamp", "heat",
				"location", "speed", "telepathy_powers",
			},
			PartKey: []string{"first_name", "last_name"},
			SortKey: []string{"timestamp"},
		},
	}

	trackingDataTable := igocqlxtable.New(*trackingDataMetadata.M)

	return &TrackingDataTable{
		Table: trackingDataTable,
	}
}
