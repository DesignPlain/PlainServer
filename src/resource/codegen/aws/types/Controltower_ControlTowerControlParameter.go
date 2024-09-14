package types

type Controltower_ControlTowerControlParameter struct {
	// The value of the parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The name of the parameter.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
