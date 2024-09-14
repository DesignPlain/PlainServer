package types

type Compute_getBackendServiceLocalityLbPolicyCustomPolicy struct {
	/*
	   An optional, arbitrary JSON object with configuration data, understood
	   by a locally installed custom policy implementation.
	*/
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	/*
	   The name of the Backend Service.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
