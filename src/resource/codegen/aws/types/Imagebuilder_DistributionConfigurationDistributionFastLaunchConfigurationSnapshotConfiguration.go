package types

type Imagebuilder_DistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfiguration struct {
	// The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.
	TargetResourceCount int `json:"targetResourceCount,omitempty" yaml:"targetResourceCount,omitempty"`
}
