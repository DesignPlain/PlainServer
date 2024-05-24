package lightsail

type Disk_attachment struct {
	// The name of the Lightsail Instance to attach to.
	InstanceName string `json:"instanceName,omitempty" yaml:"instanceName,omitempty"`

	// The name of the Lightsail Disk.
	DiskName string `json:"diskName,omitempty" yaml:"diskName,omitempty"`

	// The disk path to expose to the instance.
	DiskPath string `json:"diskPath,omitempty" yaml:"diskPath,omitempty"`
}
