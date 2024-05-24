package timestreamwrite

import types "DesignSphere_Server/src/resource/aws/types"

type Table struct {
	// The name of the Timestream table.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name of the Timestream database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
	MagneticStoreWriteProperties types.Timestreamwrite_TableMagneticStoreWriteProperties `json:"magneticStoreWriteProperties,omitempty" yaml:"magneticStoreWriteProperties,omitempty"`

	// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
	RetentionProperties types.Timestreamwrite_TableRetentionProperties `json:"retentionProperties,omitempty" yaml:"retentionProperties,omitempty"`

	// The schema of the table. See Schema below for more details.
	Schema types.Timestreamwrite_TableSchema `json:"schema,omitempty" yaml:"schema,omitempty"`
}
