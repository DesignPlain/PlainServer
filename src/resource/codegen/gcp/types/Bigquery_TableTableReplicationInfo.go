package types

type Bigquery_TableTableReplicationInfo struct {
	// The interval at which the source materialized view is polled for updates. The default is 300000.
	ReplicationIntervalMs int `json:"replicationIntervalMs,omitempty" yaml:"replicationIntervalMs,omitempty"`

	// The ID of the source dataset.
	SourceDatasetId string `json:"sourceDatasetId,omitempty" yaml:"sourceDatasetId,omitempty"`

	// The ID of the source project.
	SourceProjectId string `json:"sourceProjectId,omitempty" yaml:"sourceProjectId,omitempty"`

	// The ID of the source materialized view.
	SourceTableId string `json:"sourceTableId,omitempty" yaml:"sourceTableId,omitempty"`
}
