package types

type Datastream_StreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemaOracleTableOracleColumn struct {
	/*
	   (Output)
	   Column length.
	*/
	Length int `json:"length,omitempty" yaml:"length,omitempty"`

	/*
	   (Output)
	   Whether or not the column can accept a null value.
	*/
	Nullable bool `json:"nullable,omitempty" yaml:"nullable,omitempty"`

	/*
	   (Output)
	   The ordinal position of the column in the table.
	*/
	OrdinalPosition int `json:"ordinalPosition,omitempty" yaml:"ordinalPosition,omitempty"`

	/*
	   (Output)
	   Column precision.
	*/
	Precision int `json:"precision,omitempty" yaml:"precision,omitempty"`

	// Column name.
	Column string `json:"column,omitempty" yaml:"column,omitempty"`

	/*
	   The Oracle data type. Full data types list can be found here:
	   https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Data-Types.html
	*/
	DataType string `json:"dataType,omitempty" yaml:"dataType,omitempty"`

	/*
	   (Output)
	   Column encoding.
	*/
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`

	/*
	   (Output)
	   Whether or not the column represents a primary key.
	*/
	PrimaryKey bool `json:"primaryKey,omitempty" yaml:"primaryKey,omitempty"`

	/*
	   (Output)
	   Column scale.
	*/
	Scale int `json:"scale,omitempty" yaml:"scale,omitempty"`
}
