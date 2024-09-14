package ec2

type DedicatedHost struct {
	// The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
	OutpostArn string `json:"outpostArn,omitempty" yaml:"outpostArn,omitempty"`

	// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
	AssetId string `json:"assetId,omitempty" yaml:"assetId,omitempty"`

	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
	AutoPlacement string `json:"autoPlacement,omitempty" yaml:"autoPlacement,omitempty"`

	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
	HostRecovery string `json:"hostRecovery,omitempty" yaml:"hostRecovery,omitempty"`

	// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
	InstanceFamily string `json:"instanceFamily,omitempty" yaml:"instanceFamily,omitempty"`

	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
}
