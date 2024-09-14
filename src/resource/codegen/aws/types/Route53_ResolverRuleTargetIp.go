package types

type Route53_ResolverRuleTargetIp struct {
	// One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses.
	Ip string `json:"ip,omitempty" yaml:"ip,omitempty"`

	// The port at `ip` that you want to forward DNS queries to. Default value is `53`.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The protocol for the resolver endpoint. Valid values can be found in the [AWS documentation](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_TargetAddress.html). Default value is `Do53`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
