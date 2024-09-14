package osconfig

import types "libds/gcp/types"

type PatchDeployment struct {
	/*
	   Duration of the patch. After the duration ends, the patch times out.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	*/
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`

	/*
	   Schedule a one-time execution.
	   Structure is documented below.
	*/
	OneTimeSchedule types.Osconfig_PatchDeploymentOneTimeSchedule `json:"oneTimeSchedule,omitempty" yaml:"oneTimeSchedule,omitempty"`

	/*
	   Patch configuration that is applied.
	   Structure is documented below.
	*/
	PatchConfig types.Osconfig_PatchDeploymentPatchConfig `json:"patchConfig,omitempty" yaml:"patchConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Schedule recurring executions.
	   Structure is documented below.
	*/
	RecurringSchedule types.Osconfig_PatchDeploymentRecurringSchedule `json:"recurringSchedule,omitempty" yaml:"recurringSchedule,omitempty"`

	/*
	   Rollout strategy of the patch job.
	   Structure is documented below.
	*/
	Rollout types.Osconfig_PatchDeploymentRollout `json:"rollout,omitempty" yaml:"rollout,omitempty"`

	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   VM instances to patch.
	   Structure is documented below.
	*/
	InstanceFilter types.Osconfig_PatchDeploymentInstanceFilter `json:"instanceFilter,omitempty" yaml:"instanceFilter,omitempty"`

	/*
	   A name for the patch deployment in the project. When creating a name the following rules apply:
	   - Must contain only lowercase letters, numbers, and hyphens.
	   - Must start with a letter.
	   - Must be between 1-63 characters.
	   - Must end with a number or a letter.
	   - Must be unique within the project.
	*/
	PatchDeploymentId string `json:"patchDeploymentId,omitempty" yaml:"patchDeploymentId,omitempty"`
}
