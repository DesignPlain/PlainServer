package alb

import types "libds/aws/types"

type LoadBalancer struct {
	// Whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
	DropInvalidHeaderFields bool `json:"dropInvalidHeaderFields,omitempty" yaml:"dropInvalidHeaderFields,omitempty"`

	// If true, cross-zone load balancing of the load balancer will be enabled. For `network` and `gateway` type load balancers, this feature is disabled by default (`false`). For `application` load balancer this feature is always enabled (`true`) and cannot be disabled. Defaults to `false`.
	EnableCrossZoneLoadBalancing bool `json:"enableCrossZoneLoadBalancing,omitempty" yaml:"enableCrossZoneLoadBalancing,omitempty"`

	// Whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. Defaults to `false`.
	EnableWafFailOpen bool `json:"enableWafFailOpen,omitempty" yaml:"enableWafFailOpen,omitempty"`

	// If true, the LB will be internal. Defaults to `false`.
	Internal bool `json:"internal,omitempty" yaml:"internal,omitempty"`

	// Subnet mapping block. See below. For Load Balancers of type `network` subnet mappings can only be added.
	SubnetMappings []types.Alb_LoadBalancerSubnetMapping `json:"subnetMappings,omitempty" yaml:"subnetMappings,omitempty"`

	// Determines how the load balancer modifies the `X-Forwarded-For` header in the HTTP request before sending the request to the target. The possible values are `append`, `preserve`, and `remove`. Only valid for Load Balancers of type `application`. The default is `append`.
	XffHeaderProcessingMode string `json:"xffHeaderProcessingMode,omitempty" yaml:"xffHeaderProcessingMode,omitempty"`

	// Connection Logs block. See below. Only valid for Load Balancers of type `application`.
	ConnectionLogs types.Alb_LoadBalancerConnectionLogs `json:"connectionLogs,omitempty" yaml:"connectionLogs,omitempty"`

	// Type of load balancer to create. Possible values are `application`, `gateway`, or `network`. The default value is `application`.
	LoadBalancerType string `json:"loadBalancerType,omitempty" yaml:"loadBalancerType,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Whether the Application Load Balancer should preserve the Host header in the HTTP request and send it to the target without any change. Defaults to `false`.
	PreserveHostHeader bool `json:"preserveHostHeader,omitempty" yaml:"preserveHostHeader,omitempty"`

	// List of security group IDs to assign to the LB. Only valid for Load Balancers of type `application` or `network`. For load balancers of type `network` security groups cannot be added if none are currently present, and cannot all be removed once added. If either of these conditions are met, this will force a recreation of the resource.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// Client keep alive value in seconds. The valid range is 60-604800 seconds. The default is 3600 seconds.
	ClientKeepAlive int `json:"clientKeepAlive,omitempty" yaml:"clientKeepAlive,omitempty"`

	// ID of the customer owned ipv4 pool to use for this load balancer.
	CustomerOwnedIpv4Pool string `json:"customerOwnedIpv4Pool,omitempty" yaml:"customerOwnedIpv4Pool,omitempty"`

	// How the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are `monitor`, `defensive` (default), `strictest`.
	DesyncMitigationMode string `json:"desyncMitigationMode,omitempty" yaml:"desyncMitigationMode,omitempty"`

	// If true, deletion of the load balancer will be disabled via the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection bool `json:"enableDeletionProtection,omitempty" yaml:"enableDeletionProtection,omitempty"`

	// Whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 bool `json:"enableHttp2,omitempty" yaml:"enableHttp2,omitempty"`

	// Whether the two headers (`x-amzn-tls-version` and `x-amzn-tls-cipher-suite`), which contain information about the negotiated TLS version and cipher suite, are added to the client request before sending it to the target. Only valid for Load Balancers of type `application`. Defaults to `false`
	EnableTlsVersionAndCipherSuiteHeaders bool `json:"enableTlsVersionAndCipherSuiteHeaders,omitempty" yaml:"enableTlsVersionAndCipherSuiteHeaders,omitempty"`

	// Whether inbound security group rules are enforced for traffic originating from a PrivateLink. Only valid for Load Balancers of type `network`. The possible values are `on` and `off`.
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic string `json:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic,omitempty" yaml:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic,omitempty"`

	// Access Logs block. See below.
	AccessLogs types.Alb_LoadBalancerAccessLogs `json:"accessLogs,omitempty" yaml:"accessLogs,omitempty"`

	// How traffic is distributed among the load balancer Availability Zones. Possible values are `any_availability_zone` (default), `availability_zone_affinity`, or `partial_availability_zone_affinity`. See   [Availability Zone DNS affinity](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#zonal-dns-affinity) for additional details. Only valid for `network` type load balancers.
	DnsRecordClientRoutingPolicy string `json:"dnsRecordClientRoutingPolicy,omitempty" yaml:"dnsRecordClientRoutingPolicy,omitempty"`

	// Whether the X-Forwarded-For header should preserve the source port that the client used to connect to the load balancer in `application` load balancers. Defaults to `false`.
	EnableXffClientPort bool `json:"enableXffClientPort,omitempty" yaml:"enableXffClientPort,omitempty"`

	// Time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout int `json:"idleTimeout,omitempty" yaml:"idleTimeout,omitempty"`

	// Type of IP addresses used by the subnets for your load balancer. The possible values depend upon the load balancer type: `ipv4` (all load balancer types), `dualstack` (all load balancer types), and `dualstack-without-public-ipv4` (type `application` only).
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	// Name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified, this provider will autogenerate a name beginning with `tf-lb`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of subnet IDs to attach to the LB. For Load Balancers of type `network` subnets can only be added (see [Availability Zones](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#availability-zones)), deleting a subnet for load balancers of type `network` will force a recreation of the resource.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
