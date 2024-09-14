package types

type Bigquery_TableTableConstraints struct {
	/*
	   Present only if the table has a foreign key.
	   The foreign key is not enforced.
	   Structure is documented below.
	*/
	ForeignKeys []Bigquery_TableTableConstraintsForeignKey `json:"foreignKeys,omitempty" yaml:"foreignKeys,omitempty"`

	/*
	   Represents the primary key constraint
	   on a table's columns. Present only if the table has a primary key.
	   The primary key is not enforced.
	   Structure is documented below.
	*/
	PrimaryKey Bigquery_TableTableConstraintsPrimaryKey `json:"primaryKey,omitempty" yaml:"primaryKey,omitempty"`
}
