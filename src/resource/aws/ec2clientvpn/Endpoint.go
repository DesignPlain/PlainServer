package ec2clientvpn

import types "DesignSphere_Server/src/resource/aws/types"

type Endpoint struct {
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock string `json:"clientCidrBlock,omitempty" yaml:"clientCidrBlock,omitempty"`

	// Information about the client connection logging options.
	ConnectionLogOptions types.Ec2clientvpn_EndpointConnectionLogOptions `json:"connectionLogOptions,omitempty" yaml:"connectionLogOptions,omitempty"`

	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the connecting device is used.
	DnsServers []string `json:"dnsServers,omitempty" yaml:"dnsServers,omitempty"`

	// The maximum session duration is a trigger by which end-users are required to re-authenticate prior to establishing a VPN session. Default value is `24` - Valid values: `8 | 10 | 12 | 24`
	SessionTimeoutHours int `json:"sessionTimeoutHours,omitempty" yaml:"sessionTimeoutHours,omitempty"`

	// The ID of the VPC to associate with the Client VPN endpoint. If no security group IDs are specified in the request, the default security group for the VPC is applied.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
	ClientLoginBannerOptions types.Ec2clientvpn_EndpointClientLoginBannerOptions `json:"clientLoginBannerOptions,omitempty" yaml:"clientLoginBannerOptions,omitempty"`

	// A brief description of the Client VPN endpoint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The IDs of one or more security groups to apply to the target network. You must also specify the ID of the VPC that contains the security groups.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The ARN of the ACM server certificate.
	ServerCertificateArn string `json:"serverCertificateArn,omitempty" yaml:"serverCertificateArn,omitempty"`

	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel bool `json:"splitTunnel,omitempty" yaml:"splitTunnel,omitempty"`

	// The port number for the Client VPN endpoint. Valid values are `443` and `1194`. Default value is `443`.
	VpnPort int `json:"vpnPort,omitempty" yaml:"vpnPort,omitempty"`

	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions []types.Ec2clientvpn_EndpointAuthenticationOption `json:"authenticationOptions,omitempty" yaml:"authenticationOptions,omitempty"`

	// The options for managing connection authorization for new client connections.
	ClientConnectOptions types.Ec2clientvpn_EndpointClientConnectOptions `json:"clientConnectOptions,omitempty" yaml:"clientConnectOptions,omitempty"`

	// Specify whether to enable the self-service portal for the Client VPN endpoint. Values can be `enabled` or `disabled`. Default value is `disabled`.
	SelfServicePortal string `json:"selfServicePortal,omitempty" yaml:"selfServicePortal,omitempty"`

	// A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol string `json:"transportProtocol,omitempty" yaml:"transportProtocol,omitempty"`
}
