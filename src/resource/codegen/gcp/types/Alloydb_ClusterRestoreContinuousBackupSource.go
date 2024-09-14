package types

type Alloydb_ClusterRestoreContinuousBackupSource struct {
	// The name of the source cluster that this cluster is restored from.
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// The point in time that this cluster is restored to, in RFC 3339 format.
	PointInTime string `json:"pointInTime,omitempty" yaml:"pointInTime,omitempty"`
}
