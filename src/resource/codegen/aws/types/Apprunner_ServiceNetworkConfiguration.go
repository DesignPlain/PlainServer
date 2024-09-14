package types

type Apprunner_ServiceNetworkConfiguration struct {
	// Network configuration settings for outbound message traffic. See Egress Configuration below for more details.
	EgressConfiguration Apprunner_ServiceNetworkConfigurationEgressConfiguration `json:"egressConfiguration,omitempty" yaml:"egressConfiguration,omitempty"`

	// Network configuration settings for inbound network traffic. See Ingress Configuration below for more details.
	IngressConfiguration Apprunner_ServiceNetworkConfigurationIngressConfiguration `json:"ingressConfiguration,omitempty" yaml:"ingressConfiguration,omitempty"`

	// App Runner provides you with the option to choose between Internet Protocol version 4 (IPv4) and dual stack (IPv4 and IPv6) for your incoming public network configuration. Valid values: `IPV4`, `DUAL_STACK`. Default: `IPV4`.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`
}
