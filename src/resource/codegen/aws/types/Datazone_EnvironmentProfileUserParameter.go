package types

type Datazone_EnvironmentProfileUserParameter struct {
	// Name of the environment profile parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the environment profile parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
