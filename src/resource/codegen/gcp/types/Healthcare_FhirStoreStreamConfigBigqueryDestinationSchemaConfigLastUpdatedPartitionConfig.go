package types

type Healthcare_FhirStoreStreamConfigBigqueryDestinationSchemaConfigLastUpdatedPartitionConfig struct {
	// Number of milliseconds for which to keep the storage for a partition.
	ExpirationMs string `json:"expirationMs,omitempty" yaml:"expirationMs,omitempty"`

	/*
	   Type of partitioning.
	   Possible values are: `PARTITION_TYPE_UNSPECIFIED`, `HOUR`, `DAY`, `MONTH`, `YEAR`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
