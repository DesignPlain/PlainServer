package types

type Dns_ManagedZonePeeringConfig struct {
	/*
	   The network with which to peer.
	   Structure is documented below.
	*/
	TargetNetwork Dns_ManagedZonePeeringConfigTargetNetwork `json:"targetNetwork,omitempty" yaml:"targetNetwork,omitempty"`
}
