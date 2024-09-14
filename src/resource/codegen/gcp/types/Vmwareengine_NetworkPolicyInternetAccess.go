package types

type Vmwareengine_NetworkPolicyInternetAccess struct {
	// True if the service is enabled; false otherwise.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   (Output)
	   State of the service. New values may be added to this enum when appropriate.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
