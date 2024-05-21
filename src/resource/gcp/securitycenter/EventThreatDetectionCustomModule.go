package securitycenter

type EventThreatDetectionCustomModule struct {
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   Config for the module. For the resident module, its config value is defined at this level.
	   For the inherited module, its config value is inherited from the ancestor module.
	*/
	Config string `json:"config,omitempty" yaml:"config,omitempty"`

	// The human readable name to be displayed for the module.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The state of enablement for the module at the given level of the hierarchy.
	   Possible values are: `ENABLED`, `DISABLED`.
	*/
	EnablementState string `json:"enablementState,omitempty" yaml:"enablementState,omitempty"`

	/*
	   Numerical ID of the parent organization.


	   - - -
	*/
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`
}
