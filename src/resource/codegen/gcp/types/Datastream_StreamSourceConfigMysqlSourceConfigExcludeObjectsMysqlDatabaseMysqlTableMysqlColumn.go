package types

type Datastream_StreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabaseMysqlTableMysqlColumn struct {
	// Column collation.
	Collation string `json:"collation,omitempty" yaml:"collation,omitempty"`

	// Column name.
	Column string `json:"column,omitempty" yaml:"column,omitempty"`

	/*
	   The MySQL data type. Full data types list can be found here:
	   https://dev.mysql.com/doc/refman/8.0/en/data-types.html
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

	// Whether or not the column represents a primary key.
	PrimaryKey bool `json:"primaryKey,omitempty" yaml:"primaryKey,omitempty"`
}
