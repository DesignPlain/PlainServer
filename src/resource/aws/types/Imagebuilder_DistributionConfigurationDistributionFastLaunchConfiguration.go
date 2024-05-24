package types

type Imagebuilder_DistributionConfigurationDistributionFastLaunchConfiguration struct {
	// A Boolean that represents the current state of faster launching for the Windows AMI. Set to `true` to start using Windows faster launching, or `false` to stop using it.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Configuration block for the launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots. Detailed below.
	LaunchTemplate Imagebuilder_DistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate `json:"launchTemplate,omitempty" yaml:"launchTemplate,omitempty"`

	// The maximum number of parallel instances that are launched for creating resources.
	MaxParallelLaunches int `json:"maxParallelLaunches,omitempty" yaml:"maxParallelLaunches,omitempty"`

	// Configuration block for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled. Detailed below.
	SnapshotConfiguration Imagebuilder_DistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfiguration `json:"snapshotConfiguration,omitempty" yaml:"snapshotConfiguration,omitempty"`

	// The owner account ID for the fast-launch enabled Windows AMI.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}
