package lightsail

type Disk struct {
	// The Availability Zone in which to create your disk.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The name of the disk.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The size of the disk in GB.
	SizeInGb int `json:"sizeInGb,omitempty" yaml:"sizeInGb,omitempty"`

	// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
