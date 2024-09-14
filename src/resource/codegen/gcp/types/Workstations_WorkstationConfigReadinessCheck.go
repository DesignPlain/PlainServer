package types

type Workstations_WorkstationConfigReadinessCheck struct {
	// Path to which the request should be sent.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Port to which the request should be sent.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
