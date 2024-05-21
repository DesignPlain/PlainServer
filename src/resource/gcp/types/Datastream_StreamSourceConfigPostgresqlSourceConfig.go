package types



type Datastream_StreamSourceConfigPostgresqlSourceConfig struct {
	/*
	   The name of the publication that includes the set of all tables
	   that are defined in the stream's include_objects.
	*/
	Publication string `json:"publication,omitempty" yaml:"publication,omitempty"`

	/*
	   The name of the logical replication slot that's configured with
	   the pgoutput plugin.
	*/
	ReplicationSlot string `json:"replicationSlot,omitempty" yaml:"replicationSlot,omitempty"`

	/*
	   PostgreSQL objects to exclude from the stream.
	   Structure is documented below.
	*/
	ExcludeObjects Datastream_StreamSourceConfigPostgresqlSourceConfigExcludeObjects `json:"excludeObjects,omitempty" yaml:"excludeObjects,omitempty"`

	/*
	   PostgreSQL objects to retrieve from the source.
	   Structure is documented below.
	*/
	IncludeObjects Datastream_StreamSourceConfigPostgresqlSourceConfigIncludeObjects `json:"includeObjects,omitempty" yaml:"includeObjects,omitempty"`

	/*
	   Maximum number of concurrent backfill tasks. The number should be non
	   negative. If not set (or set to 0), the system's default value will be used.
	*/
	MaxConcurrentBackfillTasks int `json:"maxConcurrentBackfillTasks,omitempty" yaml:"maxConcurrentBackfillTasks,omitempty"`
}
