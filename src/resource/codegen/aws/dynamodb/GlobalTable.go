package dynamodb

import types "libds/aws/types"

type GlobalTable struct {
	// The name of the global table. Must match underlying DynamoDB Table names in all regions.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Underlying DynamoDB Table. At least 1 replica must be defined. See below.
	Replicas []types.Dynamodb_GlobalTableReplica `json:"replicas,omitempty" yaml:"replicas,omitempty"`
}
