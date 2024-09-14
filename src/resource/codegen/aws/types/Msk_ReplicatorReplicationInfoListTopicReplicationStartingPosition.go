package types

type Msk_ReplicatorReplicationInfoListTopicReplicationStartingPosition struct {
	// The type of replication starting position. Supports `LATEST` and `EARLIEST`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
