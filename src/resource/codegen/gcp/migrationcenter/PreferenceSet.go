package migrationcenter

import types "libds/gcp/types"

type PreferenceSet struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   VirtualMachinePreferences enables you to create sets of assumptions, for example, a geographical location and pricing track, for your migrated virtual machines. The set of preferences influence recommendations for migrating virtual machine assets.
	   Structure is documented below.
	*/
	VirtualMachinePreferences types.Migrationcenter_PreferenceSetVirtualMachinePreferences `json:"virtualMachinePreferences,omitempty" yaml:"virtualMachinePreferences,omitempty"`

	// A description of the preference set.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// User-friendly display name. Maximum length is 63 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Part of `parent`. See documentation of `projectsId`.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Required. User specified ID for the preference set. It will become the last component of the preference set name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression `a-z?`.


	   - - -
	*/
	PreferenceSetId string `json:"preferenceSetId,omitempty" yaml:"preferenceSetId,omitempty"`
}
