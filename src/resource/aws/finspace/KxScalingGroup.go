package finspace

type KxScalingGroup struct {
	// A unique identifier for the kdb environment, where you want to create the scaling group.
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	/*
	   The memory and CPU capabilities of the scaling group host on which FinSpace Managed kdb clusters will be placed.

	   The following arguments are optional:
	*/
	HostType string `json:"hostType,omitempty" yaml:"hostType,omitempty"`

	// Unique name for the scaling group that you want to create.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. You can add up to 50 tags to a scaling group.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The availability zone identifiers for the requested regions.
	AvailabilityZoneId string `json:"availabilityZoneId,omitempty" yaml:"availabilityZoneId,omitempty"`
}
