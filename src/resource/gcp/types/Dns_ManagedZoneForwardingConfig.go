package types

type Dns_ManagedZoneForwardingConfig struct {
	/*
	   List of target name servers to forward to. Cloud DNS will
	   select the best available name server if more than
	   one target is given.
	   Structure is documented below.
	*/
	TargetNameServers []Dns_ManagedZoneForwardingConfigTargetNameServer `json:"targetNameServers,omitempty" yaml:"targetNameServers,omitempty"`
}
