package types

type Cloudrun_getServiceTemplateSpecVolume struct {
	// A filesystem specified by the Container Storage Interface (CSI).
	Csis []Cloudrun_getServiceTemplateSpecVolumeCsi `json:"csis,omitempty" yaml:"csis,omitempty"`

	// Ephemeral storage which can be backed by real disks (HD, SSD), network storage or memory (i.e. tmpfs). For now only in memory (tmpfs) is supported. It is ephemeral in the sense that when the sandbox is taken down, the data is destroyed with it (it does not persist across sandbox runs).
	EmptyDirs []Cloudrun_getServiceTemplateSpecVolumeEmptyDir `json:"emptyDirs,omitempty" yaml:"emptyDirs,omitempty"`

	// The name of the Cloud Run Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The secret's value will be presented as the content of a file whose
	   name is defined in the item path. If no items are defined, the name of
	   the file is the secret_name.
	*/
	Secrets []Cloudrun_getServiceTemplateSpecVolumeSecret `json:"secrets,omitempty" yaml:"secrets,omitempty"`
}
