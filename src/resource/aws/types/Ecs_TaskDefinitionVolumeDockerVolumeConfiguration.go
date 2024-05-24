package types

type Ecs_TaskDefinitionVolumeDockerVolumeConfiguration struct {
	// Scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as `shared` persist after the task stops.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// If this value is `true`, the Docker volume is created if it does not already exist. -Note-: This field is only used if the scope is `shared`.
	Autoprovision bool `json:"autoprovision,omitempty" yaml:"autoprovision,omitempty"`

	// Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
	Driver string `json:"driver,omitempty" yaml:"driver,omitempty"`

	// Map of Docker driver specific options.
	DriverOpts map[string]string `json:"driverOpts,omitempty" yaml:"driverOpts,omitempty"`

	// Map of custom metadata to add to your Docker volume.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
