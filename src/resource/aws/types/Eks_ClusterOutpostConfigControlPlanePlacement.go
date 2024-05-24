package types

type Eks_ClusterOutpostConfigControlPlanePlacement struct {
	// The name of the placement group for the Kubernetes control plane instances. This setting can't be changed after cluster creation.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
}
