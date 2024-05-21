package types

type Biglake_TableHiveOptions struct {
	/*
	   Stores user supplied Hive table parameters. An object containing a
	   list of "key": value pairs.
	   Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	*/
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	/*
	   Stores physical storage information on the data.
	   Structure is documented below.
	*/
	StorageDescriptor Biglake_TableHiveOptionsStorageDescriptor `json:"storageDescriptor,omitempty" yaml:"storageDescriptor,omitempty"`

	// Hive table type. For example, MANAGED_TABLE, EXTERNAL_TABLE.
	TableType string `json:"tableType,omitempty" yaml:"tableType,omitempty"`
}
