package types

type Auditmanager_FrameworkControlSet struct {
	// List of controls within the control set. See `controls` below.
	Controls []Auditmanager_FrameworkControlSetControl `json:"controls,omitempty" yaml:"controls,omitempty"`

	// Unique identifier of the control.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the control set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
