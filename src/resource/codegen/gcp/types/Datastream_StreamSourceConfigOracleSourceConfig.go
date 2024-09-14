package types

type Datastream_StreamSourceConfigOracleSourceConfig struct {
	// Configuration to drop large object values.
	StreamLargeObjects Datastream_StreamSourceConfigOracleSourceConfigStreamLargeObjects `json:"streamLargeObjects,omitempty" yaml:"streamLargeObjects,omitempty"`

	// Configuration to drop large object values.
	DropLargeObjects Datastream_StreamSourceConfigOracleSourceConfigDropLargeObjects `json:"dropLargeObjects,omitempty" yaml:"dropLargeObjects,omitempty"`

	/*
	   Oracle objects to exclude from the stream.
	   Structure is documented below.
	*/
	ExcludeObjects Datastream_StreamSourceConfigOracleSourceConfigExcludeObjects `json:"excludeObjects,omitempty" yaml:"excludeObjects,omitempty"`

	/*
	   Oracle objects to retrieve from the source.
	   Structure is documented below.
	*/
	IncludeObjects Datastream_StreamSourceConfigOracleSourceConfigIncludeObjects `json:"includeObjects,omitempty" yaml:"includeObjects,omitempty"`

	/*
	   Maximum number of concurrent backfill tasks. The number should be non negative.
	   If not set (or set to 0), the system's default value will be used.
	*/
	MaxConcurrentBackfillTasks int `json:"maxConcurrentBackfillTasks,omitempty" yaml:"maxConcurrentBackfillTasks,omitempty"`

	/*
	   Maximum number of concurrent CDC tasks. The number should be non negative.
	   If not set (or set to 0), the system's default value will be used.
	*/
	MaxConcurrentCdcTasks int `json:"maxConcurrentCdcTasks,omitempty" yaml:"maxConcurrentCdcTasks,omitempty"`
}
