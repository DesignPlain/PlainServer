package lightsail

type Disk struct {
	// The Availability Zone in which to create your disk.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The name of the Lightsail load balancer.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The instance port the load balancer will connect.
	SizeInGb int `json:"sizeInGb,omitempty" yaml:"sizeInGb,omitempty"`

	// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
