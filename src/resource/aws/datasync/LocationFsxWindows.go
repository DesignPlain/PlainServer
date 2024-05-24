package datasync

type LocationFsxWindows struct {
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns []string `json:"securityGroupArns,omitempty" yaml:"securityGroupArns,omitempty"`

	// Subdirectory to perform actions as source or destination.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User string `json:"user,omitempty" yaml:"user,omitempty"`

	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn string `json:"fsxFilesystemArn,omitempty" yaml:"fsxFilesystemArn,omitempty"`

	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}
