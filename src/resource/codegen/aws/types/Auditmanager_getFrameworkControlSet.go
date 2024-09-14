package types

type Auditmanager_getFrameworkControlSet struct {
	//
	Controls []Auditmanager_getFrameworkControlSetControl `json:"controls,omitempty" yaml:"controls,omitempty"`

	//
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the framework.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
