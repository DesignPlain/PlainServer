package types

type Eks_getClusterOutpostConfigControlPlanePlacement struct {
	// The name of the placement group for the Kubernetes control plane instances.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
}
