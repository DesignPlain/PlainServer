package securitycenter

import types "libds/gcp/types"

type FolderCustomModule struct {
	// Numerical ID of the parent folder.
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	/*
	   The user specified custom configuration for the module.
	   Structure is documented below.
	*/
	CustomConfig types.Securitycenter_FolderCustomModuleCustomConfig `json:"customConfig,omitempty" yaml:"customConfig,omitempty"`

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
}
