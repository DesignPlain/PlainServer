package types

type Vpclattice_TargetGroupConfig struct {
	// The protocol to use for routing traffic to the targets. Valid Values are `HTTP` | `HTTPS`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// The protocol version. Valid Values are `HTTP1` | `HTTP2` | `GRPC`. Default value is `HTTP1`.
	ProtocolVersion string `json:"protocolVersion,omitempty" yaml:"protocolVersion,omitempty"`

	// The ID of the VPC.
	VpcIdentifier string `json:"vpcIdentifier,omitempty" yaml:"vpcIdentifier,omitempty"`

	// The health check configuration.
	HealthCheck Vpclattice_TargetGroupConfigHealthCheck `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`

	// The type of IP address used for the target group. Valid values: `IPV4` | `IPV6`.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	// The version of the event structure that the Lambda function receives. Supported only if `type` is `LAMBDA`. Valid Values are `V1` | `V2`.
	LambdaEventStructureVersion string `json:"lambdaEventStructureVersion,omitempty" yaml:"lambdaEventStructureVersion,omitempty"`

	// The port on which the targets are listening.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
