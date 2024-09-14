package types

type Alloydb_InstanceReadPoolConfig struct {
	// Read capacity, i.e. number of nodes in a read pool instance.
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
}
