package types

type Route53_ResolverRuleTargetIp struct {
	// One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses.
	Ip string `json:"ip,omitempty" yaml:"ip,omitempty"`

	// The port at `ip` that you want to forward DNS queries to. Default value is `53`
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
