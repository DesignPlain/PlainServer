package datasync

type LocationFsxLustre struct {
	// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
	FsxFilesystemArn string `json:"fsxFilesystemArn,omitempty" yaml:"fsxFilesystemArn,omitempty"`

	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
	SecurityGroupArns []string `json:"securityGroupArns,omitempty" yaml:"securityGroupArns,omitempty"`

	// Subdirectory to perform actions as source or destination.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
