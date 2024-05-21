package types

type Workbench_InstanceUpgradeHistory struct {
	/*
	   Use a container image to start the workbench instance.
	   Structure is documented below.
	*/
	ContainerImage string `json:"containerImage,omitempty" yaml:"containerImage,omitempty"`

	/*
	   An RFC3339 timestamp in UTC time. This in the format of yyyy-MM-ddTHH:mm:ss.SSSZ.
	   The milliseconds portion (".SSS") is optional.
	*/
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	/*
	   Definition of a custom Compute Engine virtual machine image for starting
	   a workbench instance with the environment installed directly on the VM.
	   Structure is documented below.
	*/
	VmImage string `json:"vmImage,omitempty" yaml:"vmImage,omitempty"`

	// Optional. Action. Rolloback or Upgrade.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// Optional. The snapshot of the boot disk of this workbench instance before upgrade.
	Snapshot string `json:"snapshot,omitempty" yaml:"snapshot,omitempty"`

	/*
	   (Output)
	   Output only. The state of this instance upgrade history entry.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Optional. Target VM Version, like m63.
	TargetVersion string `json:"targetVersion,omitempty" yaml:"targetVersion,omitempty"`

	// Optional. The version of the workbench instance before this upgrade.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Optional. The framework of this workbench instance.
	Framework string `json:"framework,omitempty" yaml:"framework,omitempty"`
}
