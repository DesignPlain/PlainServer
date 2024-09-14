package types

type Appfabric_IngestionDestinationProcessingConfigurationAuditLog struct {
	// The format in which the audit logs need to be formatted. Valid values: `json`, `parquet`.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// The event schema in which the audit logs need to be formatted. Valid values: `ocsf`, `raw`.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
