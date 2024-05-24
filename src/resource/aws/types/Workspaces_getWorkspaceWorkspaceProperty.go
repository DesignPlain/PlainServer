package types

type Workspaces_getWorkspaceWorkspaceProperty struct {
	// Running mode. For more information, see [Manage the WorkSpace Running Mode](https://docs.aws.amazon.com/workspaces/latest/adminguide/running-mode.html). Valid values are `AUTO_STOP` and `ALWAYS_ON`.
	RunningMode string `json:"runningMode,omitempty" yaml:"runningMode,omitempty"`

	// Time after a user logs off when WorkSpaces are automatically stopped. Configured in 60-minute intervals.
	RunningModeAutoStopTimeoutInMinutes int `json:"runningModeAutoStopTimeoutInMinutes,omitempty" yaml:"runningModeAutoStopTimeoutInMinutes,omitempty"`

	// Size of the user storage.
	UserVolumeSizeGib int `json:"userVolumeSizeGib,omitempty" yaml:"userVolumeSizeGib,omitempty"`

	// Compute type. For more information, see [Amazon WorkSpaces Bundles](http://aws.amazon.com/workspaces/details/#Amazon_WorkSpaces_Bundles). Valid values are `VALUE`, `STANDARD`, `PERFORMANCE`, `POWER`, `GRAPHICS`, `POWERPRO` and `GRAPHICSPRO`.
	ComputeTypeName string `json:"computeTypeName,omitempty" yaml:"computeTypeName,omitempty"`

	// Size of the root volume.
	RootVolumeSizeGib int `json:"rootVolumeSizeGib,omitempty" yaml:"rootVolumeSizeGib,omitempty"`
}
