package models

import "github.com/scylladb/gocqlx/v2/table"

type MutantDataTable struct {
	Table *table.Table
}

func NewMutantDataTable() *MutantDataTable {
	mutantDataMetadata := table.Metadata{
		Name:    "mutant_data",
		Columns: []string{"first_name", "last_name", "address", "picture_location"},
		PartKey: []string{"first_name", "last_name"},
	}

	mutantDataTable := table.New(mutantDataMetadata)

	return &MutantDataTable{
		Table: mutantDataTable,
	}
}
