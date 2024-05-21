package securitycenter

import types "DesignSphere_Server/src/resource/gcp/types"

type ProjectCustomModule struct {
	/*
	   The user specified custom configuration for the module.
	   Structure is documented below.
	*/
	CustomConfig types.Securitycenter_ProjectCustomModuleCustomConfig `json:"customConfig,omitempty" yaml:"customConfig,omitempty"`

	/*
	   The display name of the Security Health Analytics custom module. This
	   display name becomes the finding category for all findings that are
	   returned by this custom module. The display name must be between 1 and
	   128 characters, start with a lowercase letter, and contain alphanumeric
	   characters or underscores only.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The enablement state of the custom module.
	   Possible values are: `ENABLED`, `DISABLED`.
	*/
	EnablementState string `json:"enablementState,omitempty" yaml:"enablementState,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
