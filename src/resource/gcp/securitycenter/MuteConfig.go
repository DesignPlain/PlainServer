package securitycenter

type MuteConfig struct {
	// A description of the mute config.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   An expression that defines the filter to apply across create/update
	   events of findings. While creating a filter string, be mindful of
	   the scope in which the mute configuration is being created. E.g.,
	   If a filter contains project = X but is created under the
	   project = Y scope, it might not match any findings.
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Unique identifier provided by the client within the parent scope.
	MuteConfigId string `json:"muteConfigId,omitempty" yaml:"muteConfigId,omitempty"`

	/*
	   Resource name of the new mute configs's parent. Its format is
	   "organizations/[organization_id]", "folders/[folder_id]", or
	   "projects/[project_id]".


	   - - -
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
