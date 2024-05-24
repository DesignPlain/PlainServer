package emrserverless

import types "DesignSphere_Server/src/resource/aws/types"

type Application struct {
	// The image configuration applied to all worker types.
	ImageConfiguration types.Emrserverless_ApplicationImageConfiguration `json:"imageConfiguration,omitempty" yaml:"imageConfiguration,omitempty"`

	// The name of the application.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The network configuration for customer VPC connectivity.
	NetworkConfiguration types.Emrserverless_ApplicationNetworkConfiguration `json:"networkConfiguration,omitempty" yaml:"networkConfiguration,omitempty"`

	// The type of application you want to start, such as `spark` or `hive`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
	MaximumCapacity types.Emrserverless_ApplicationMaximumCapacity `json:"maximumCapacity,omitempty" yaml:"maximumCapacity,omitempty"`

	// The EMR release version associated with the application.
	ReleaseLabel string `json:"releaseLabel,omitempty" yaml:"releaseLabel,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The CPU architecture of an application. Valid values are `ARM64` or `X86_64`. Default value is `X86_64`.
	Architecture string `json:"architecture,omitempty" yaml:"architecture,omitempty"`

	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration types.Emrserverless_ApplicationAutoStartConfiguration `json:"autoStartConfiguration,omitempty" yaml:"autoStartConfiguration,omitempty"`

	// The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration types.Emrserverless_ApplicationAutoStopConfiguration `json:"autoStopConfiguration,omitempty" yaml:"autoStopConfiguration,omitempty"`

	// The capacity to initialize when the application is created.
	InitialCapacities []types.Emrserverless_ApplicationInitialCapacity `json:"initialCapacities,omitempty" yaml:"initialCapacities,omitempty"`
}
