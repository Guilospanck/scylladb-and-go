package models

import (
	igocqlxtable "github.com/Guilospanck/igocqlx/table"
	"github.com/scylladb/gocqlx/v2/table"
)

type MutantDataTable struct {
	Table igocqlxtable.ITable
}

func NewMutantDataTable() *MutantDataTable {
	mutantDataMetadata := igocqlxtable.Metadata{
		M: &table.Metadata{
			Name:    "mutant_data",
			Columns: []string{"first_name", "last_name", "address", "picture_location"},
			PartKey: []string{"first_name", "last_name"},
		},
	}

	mutantDataTable := igocqlxtable.New(*mutantDataMetadata.M)

	return &MutantDataTable{
		Table: mutantDataTable,
	}
}
