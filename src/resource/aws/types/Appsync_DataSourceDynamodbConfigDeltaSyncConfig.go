package types

type Appsync_DataSourceDynamodbConfigDeltaSyncConfig struct {
	// The table name.
	DeltaSyncTableName string `json:"deltaSyncTableName,omitempty" yaml:"deltaSyncTableName,omitempty"`

	// The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.
	DeltaSyncTableTtl int `json:"deltaSyncTableTtl,omitempty" yaml:"deltaSyncTableTtl,omitempty"`

	// The number of minutes that an Item is stored in the data source.
	BaseTableTtl int `json:"baseTableTtl,omitempty" yaml:"baseTableTtl,omitempty"`
}
