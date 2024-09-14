package types

type Osconfig_PatchDeploymentPatchConfigWindowsUpdate struct {
	/*
	   Only apply updates of these windows update classifications. If empty, all updates are applied.
	   Each value may be one of: `CRITICAL`, `SECURITY`, `DEFINITION`, `DRIVER`, `FEATURE_PACK`, `SERVICE_PACK`, `TOOL`, `UPDATE_ROLLUP`, `UPDATE`.
	*/
	Classifications []string `json:"classifications,omitempty" yaml:"classifications,omitempty"`

	// List of KBs to exclude from update.
	Excludes []string `json:"excludes,omitempty" yaml:"excludes,omitempty"`

	/*
	   An exclusive list of kbs to be updated. These are the only patches that will be updated.
	   This field must not be used with other patch configurations.
	*/
	ExclusivePatches []string `json:"exclusivePatches,omitempty" yaml:"exclusivePatches,omitempty"`
}
