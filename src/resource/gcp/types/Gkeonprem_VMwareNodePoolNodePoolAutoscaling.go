package types

type Gkeonprem_VMwareNodePoolNodePoolAutoscaling struct {
	// Maximum number of replicas in the NodePool.
	MaxReplicas int `json:"maxReplicas,omitempty" yaml:"maxReplicas,omitempty"`

	// Minimum number of replicas in the NodePool.
	MinReplicas int `json:"minReplicas,omitempty" yaml:"minReplicas,omitempty"`
}
