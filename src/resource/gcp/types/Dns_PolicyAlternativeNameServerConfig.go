package types

type Dns_PolicyAlternativeNameServerConfig struct {
	/*
	   Sets an alternative name server for the associated networks. When specified,
	   all DNS queries are forwarded to a name server that you choose. Names such as .internal
	   are not available when an alternative name server is specified.
	   Structure is documented below.
	*/
	TargetNameServers []Dns_PolicyAlternativeNameServerConfigTargetNameServer `json:"targetNameServers,omitempty" yaml:"targetNameServers,omitempty"`
}
