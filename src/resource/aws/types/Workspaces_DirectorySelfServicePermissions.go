package types

type Workspaces_DirectorySelfServicePermissions struct {
	// Whether WorkSpaces directory users can increase the volume size of the drives on their workspace. Default `false`.
	IncreaseVolumeSize bool `json:"increaseVolumeSize,omitempty" yaml:"increaseVolumeSize,omitempty"`

	// Whether WorkSpaces directory users can rebuild the operating system of a workspace to its original state. Default `false`.
	RebuildWorkspace bool `json:"rebuildWorkspace,omitempty" yaml:"rebuildWorkspace,omitempty"`

	// Whether WorkSpaces directory users can restart their workspace. Default `true`.
	RestartWorkspace bool `json:"restartWorkspace,omitempty" yaml:"restartWorkspace,omitempty"`

	// Whether WorkSpaces directory users can switch the running mode of their workspace. Default `false`.
	SwitchRunningMode bool `json:"switchRunningMode,omitempty" yaml:"switchRunningMode,omitempty"`

	// Whether WorkSpaces directory users can change the compute type (bundle) for their workspace. Default `false`.
	ChangeComputeType bool `json:"changeComputeType,omitempty" yaml:"changeComputeType,omitempty"`
}
