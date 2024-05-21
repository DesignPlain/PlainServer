package types

type Compute_ExternalVpnGatewayInterface struct {
	/*
	   The numeric ID for this interface. Allowed values are based on the redundancy type
	   of this external VPN gateway
	   - `0 - SINGLE_IP_INTERNALLY_REDUNDANT`
	   - `0, 1 - TWO_IPS_REDUNDANCY`
	   - `0, 1, 2, 3 - FOUR_IPS_REDUNDANCY`
	*/
	Id int `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   IP address of the interface in the external VPN gateway.
	   Only IPv4 is supported. This IP address can be either from
	   your on-premise gateway or another Cloud provider's VPN gateway,
	   it cannot be an IP address from Google Compute Engine.
	*/
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
}
