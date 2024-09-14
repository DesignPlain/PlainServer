package types

type Workspaces_getDirectorySelfServicePermission struct {
	// Whether WorkSpaces directory users can restart their workspace.
	RestartWorkspace bool `json:"restartWorkspace,omitempty" yaml:"restartWorkspace,omitempty"`

	// Whether WorkSpaces directory users can switch the running mode of their workspace.
	SwitchRunningMode bool `json:"switchRunningMode,omitempty" yaml:"switchRunningMode,omitempty"`

	// Whether WorkSpaces directory users can change the compute type (bundle) for their workspace.
	ChangeComputeType bool `json:"changeComputeType,omitempty" yaml:"changeComputeType,omitempty"`

	// Whether WorkSpaces directory users can increase the volume size of the drives on their workspace.
	IncreaseVolumeSize bool `json:"increaseVolumeSize,omitempty" yaml:"increaseVolumeSize,omitempty"`

	// Whether WorkSpaces directory users can rebuild the operating system of a workspace to its original state.
	RebuildWorkspace bool `json:"rebuildWorkspace,omitempty" yaml:"rebuildWorkspace,omitempty"`
}
