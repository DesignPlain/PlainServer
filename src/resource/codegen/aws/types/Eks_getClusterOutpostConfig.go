package types

type Eks_getClusterOutpostConfig struct {
	// List of ARNs of the Outposts hosting the EKS cluster. Only a single ARN is supported currently.
	OutpostArns []string `json:"outpostArns,omitempty" yaml:"outpostArns,omitempty"`

	// The Amazon EC2 instance type for all Kubernetes control plane instances.
	ControlPlaneInstanceType string `json:"controlPlaneInstanceType,omitempty" yaml:"controlPlaneInstanceType,omitempty"`

	// An object representing the placement configuration for all the control plane instances of your local Amazon EKS cluster on AWS Outpost.
	ControlPlanePlacements []Eks_getClusterOutpostConfigControlPlanePlacement `json:"controlPlanePlacements,omitempty" yaml:"controlPlanePlacements,omitempty"`
}
