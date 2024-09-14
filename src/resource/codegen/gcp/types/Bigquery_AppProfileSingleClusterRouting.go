package types

type Bigquery_AppProfileSingleClusterRouting struct {
	/*
	   If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
	   It is unsafe to send these requests to the same table/row/column in multiple clusters.
	*/
	AllowTransactionalWrites bool `json:"allowTransactionalWrites,omitempty" yaml:"allowTransactionalWrites,omitempty"`

	// The cluster to which read/write requests should be routed.
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
}
