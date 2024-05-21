package types

type Datastream_StreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemaPostgresqlTablePostgresqlColumn struct {
	/*
	   (Output)
	   Column precision.
	*/
	Precision int `json:"precision,omitempty" yaml:"precision,omitempty"`

	// Whether or not the column represents a primary key.
	PrimaryKey bool `json:"primaryKey,omitempty" yaml:"primaryKey,omitempty"`

	/*
	   (Output)
	   Column scale.
	*/
	Scale int `json:"scale,omitempty" yaml:"scale,omitempty"`

	// Column name.
	Column string `json:"column,omitempty" yaml:"column,omitempty"`

	/*
	   The PostgreSQL data type. Full data types list can be found here:
	   https://www.postgresql.org/docs/current/datatype.html
	*/
	DataType string `json:"dataType,omitempty" yaml:"dataType,omitempty"`

	/*
	   (Output)
	   Column length.
	*/
	Length int `json:"length,omitempty" yaml:"length,omitempty"`

	// Whether or not the column can accept a null value.
	Nullable bool `json:"nullable,omitempty" yaml:"nullable,omitempty"`

	// The ordinal position of the column in the table.
	OrdinalPosition int `json:"ordinalPosition,omitempty" yaml:"ordinalPosition,omitempty"`
}
