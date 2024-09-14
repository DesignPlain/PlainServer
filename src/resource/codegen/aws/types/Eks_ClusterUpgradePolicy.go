package types

type Eks_ClusterUpgradePolicy struct {
	// Support type to use for the cluster. If the cluster is set to `EXTENDED`, it will enter extended support at the end of standard support. If the cluster is set to `STANDARD`, it will be automatically upgraded at the end of standard support. Valid values are `EXTENDED`, `STANDARD`
	SupportType string `json:"supportType,omitempty" yaml:"supportType,omitempty"`
}
