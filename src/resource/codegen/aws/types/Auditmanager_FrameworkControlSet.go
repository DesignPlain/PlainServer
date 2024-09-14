package types

type Auditmanager_FrameworkControlSet struct {
	// Unique identifier for the framework.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the control set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block(s) for the controls within the control set. See `controls` Block below for details.
	Controls []Auditmanager_FrameworkControlSetControl `json:"controls,omitempty" yaml:"controls,omitempty"`
}
