package types



type Bigquery_TableTableConstraintsForeignKey struct {
	/*
	   The pair of the foreign key column and primary key column.
	   Structure is documented below.
	*/
	ColumnReferences Bigquery_TableTableConstraintsForeignKeyColumnReferences `json:"columnReferences,omitempty" yaml:"columnReferences,omitempty"`

	// Set only if the foreign key constraint is named.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The table that holds the primary key
	   and is referenced by this foreign key.
	   Structure is documented below.
	*/
	ReferencedTable Bigquery_TableTableConstraintsForeignKeyReferencedTable `json:"referencedTable,omitempty" yaml:"referencedTable,omitempty"`
}
