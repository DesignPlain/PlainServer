package types

type Edgecontainer_VpnConnectionDetail struct {
	/*
	   (Output)
	   The error message. This is only populated when state=ERROR.
	*/
	Error string `json:"error,omitempty" yaml:"error,omitempty"`

	/*
	   (Output)
	   The current connection state.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   (Output)
	   The Cloud Router info.
	   Structure is documented below.
	*/
	CloudRouters []Edgecontainer_VpnConnectionDetailCloudRouter `json:"cloudRouters,omitempty" yaml:"cloudRouters,omitempty"`

	/*
	   (Output)
	   Each connection has multiple Cloud VPN gateways.
	   Structure is documented below.
	*/
	CloudVpns []Edgecontainer_VpnConnectionDetailCloudVpn `json:"cloudVpns,omitempty" yaml:"cloudVpns,omitempty"`
}
