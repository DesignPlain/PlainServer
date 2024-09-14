package types

type Lambda_FunctionFileSystemConfig struct {
	// Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Path where the function can access the file system, starting with /mnt/.
	LocalMountPath string `json:"localMountPath,omitempty" yaml:"localMountPath,omitempty"`
}
