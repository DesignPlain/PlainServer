package alb

import types "libds/aws/types"

type TargetGroup struct {
	// The type of IP addresses used by the target group, only supported when target type is set to `ip`. Possible values are `ipv4` or `ipv6`.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	// Name of the target group. If omitted, this provider will assign a random, unique name. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Port on which targets receive traffic, unless overridden when registering a specific target. Required when `target_type` is `instance`, `ip` or `alb`. Does not apply when `target_type` is `lambda`.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
	ProxyProtocolV2 bool `json:"proxyProtocolV2,omitempty" yaml:"proxyProtocolV2,omitempty"`

	/*
	   Type of target that you must specify when registering targets with this target group.
	   See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_CreateTargetGroup.html) for supported values.
	   The default is `instance`.

	   Note that you can't specify targets for a target group using both instance IDs and IP addresses.

	   If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.

	   Network Load Balancers do not support the `lambda` target type.

	   Application Load Balancers do not support the `alb` target type.
	*/
	TargetType string `json:"targetType,omitempty" yaml:"targetType,omitempty"`

	// Target failover block. Only applicable for Gateway Load Balancer target groups. See target_failover for more information.
	TargetFailovers []types.Alb_TargetGroupTargetFailover `json:"targetFailovers,omitempty" yaml:"targetFailovers,omitempty"`

	// Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#deregistration-delay) for more information. Default is `false`.
	ConnectionTermination bool `json:"connectionTermination,omitempty" yaml:"connectionTermination,omitempty"`

	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay int `json:"deregistrationDelay,omitempty" yaml:"deregistrationDelay,omitempty"`

	// Health Check configuration block. Detailed below.
	HealthCheck types.Alb_TargetGroupHealthCheck `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`

	// Determines whether to enable target anomaly mitigation.  Target anomaly mitigation is only supported by the `weighted_random` load balancing algorithm type.  See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#automatic-target-weights) for more information.  The value is `"on"` or `"off"`. The default is `"off"`.
	LoadBalancingAnomalyMitigation string `json:"loadBalancingAnomalyMitigation,omitempty" yaml:"loadBalancingAnomalyMitigation,omitempty"`

	// Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify `GRPC` to send requests to targets using gRPC. Specify `HTTP2` to send requests to targets using HTTP/2. The default is `HTTP1`, which sends requests to targets using HTTP/1.1
	ProtocolVersion string `json:"protocolVersion,omitempty" yaml:"protocolVersion,omitempty"`

	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart int `json:"slowStart,omitempty" yaml:"slowStart,omitempty"`

	// Indicates whether cross zone load balancing is enabled. The value is `"true"`, `"false"` or `"use_load_balancer_configuration"`. The default is `"use_load_balancer_configuration"`.
	LoadBalancingCrossZoneEnabled string `json:"loadBalancingCrossZoneEnabled,omitempty" yaml:"loadBalancingCrossZoneEnabled,omitempty"`

	// Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
	PreserveClientIp string `json:"preserveClientIp,omitempty" yaml:"preserveClientIp,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Target health requirements block. See target_group_health for more information.
	TargetGroupHealth types.Alb_TargetGroupTargetGroupHealth `json:"targetGroupHealth,omitempty" yaml:"targetGroupHealth,omitempty"`

	// Target health state block. Only applicable for Network Load Balancer target groups when `protocol` is `TCP` or `TLS`. See target_health_state for more information.
	TargetHealthStates []types.Alb_TargetGroupTargetHealthState `json:"targetHealthStates,omitempty" yaml:"targetHealthStates,omitempty"`

	// Identifier of the VPC in which to create the target group. Required when `target_type` is `instance`, `ip` or `alb`. Does not apply when `target_type` is `lambda`.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `target_type` is `lambda`. Default is `false`.
	LambdaMultiValueHeadersEnabled bool `json:"lambdaMultiValueHeadersEnabled,omitempty" yaml:"lambdaMultiValueHeadersEnabled,omitempty"`

	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `round_robin`, `least_outstanding_requests`, or `weighted_random`. The default is `round_robin`.
	LoadBalancingAlgorithmType string `json:"loadBalancingAlgorithmType,omitempty" yaml:"loadBalancingAlgorithmType,omitempty"`

	/*
	   Protocol to use for routing traffic to the targets.
	   Should be one of `GENEVE`, `HTTP`, `HTTPS`, `TCP`, `TCP_UDP`, `TLS`, or `UDP`.
	   Required when `target_type` is `instance`, `ip`, or `alb`.
	   Does not apply when `target_type` is `lambda`.
	*/
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Stickiness configuration block. Detailed below.
	Stickiness types.Alb_TargetGroupStickiness `json:"stickiness,omitempty" yaml:"stickiness,omitempty"`
}
