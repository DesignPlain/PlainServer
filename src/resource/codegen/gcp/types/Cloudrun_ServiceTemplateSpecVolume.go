package types

type Cloudrun_ServiceTemplateSpecVolume struct {
	/*
	   A filesystem specified by the Container Storage Interface (CSI).
	   Structure is documented below.
	*/
	Csi Cloudrun_ServiceTemplateSpecVolumeCsi `json:"csi,omitempty" yaml:"csi,omitempty"`

	/*
	   Ephemeral storage which can be backed by real disks (HD, SSD), network storage or memory (i.e. tmpfs). For now only in memory (tmpfs) is supported. It is ephemeral in the sense that when the sandbox is taken down, the data is destroyed with it (it does not persist across sandbox runs).
	   Structure is documented below.
	*/
	EmptyDir Cloudrun_ServiceTemplateSpecVolumeEmptyDir `json:"emptyDir,omitempty" yaml:"emptyDir,omitempty"`

	// Volume's name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The secret's value will be presented as the content of a file whose
	   name is defined in the item path. If no items are defined, the name of
	   the file is the secret_name.
	   Structure is documented below.
	*/
	Secret Cloudrun_ServiceTemplateSpecVolumeSecret `json:"secret,omitempty" yaml:"secret,omitempty"`
}
