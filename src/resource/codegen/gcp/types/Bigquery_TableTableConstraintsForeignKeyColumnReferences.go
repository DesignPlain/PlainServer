package types

type Bigquery_TableTableConstraintsForeignKeyColumnReferences struct {
	/*
	   The column in the primary key that are
	   referenced by the referencingColumn
	*/
	ReferencedColumn string `json:"referencedColumn,omitempty" yaml:"referencedColumn,omitempty"`

	// The column that composes the foreign key.
	ReferencingColumn string `json:"referencingColumn,omitempty" yaml:"referencingColumn,omitempty"`
}
