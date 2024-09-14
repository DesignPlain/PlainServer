package types

type Imagebuilder_getDistributionConfigurationDistributionFastLaunchConfiguration struct {
	// The maximum number of parallel instances that are launched for creating resources.
	MaxParallelLaunches int `json:"maxParallelLaunches,omitempty" yaml:"maxParallelLaunches,omitempty"`

	// Nested list of configurations for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.
	SnapshotConfigurations []Imagebuilder_getDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfiguration `json:"snapshotConfigurations,omitempty" yaml:"snapshotConfigurations,omitempty"`

	// The account ID that this configuration applies to.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// A Boolean that represents the current state of faster launching for the Windows AMI.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Nested list of launch templates that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.
	LaunchTemplates []Imagebuilder_getDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate `json:"launchTemplates,omitempty" yaml:"launchTemplates,omitempty"`
}
