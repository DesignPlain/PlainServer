package types

type Lb_TargetGroupTargetFailover struct {
	// Indicates how the GWLB handles existing flows when a target is deregistered. Possible values are `rebalance` and `no_rebalance`. Must match the attribute value set for `on_unhealthy`. Default: `no_rebalance`.
	OnDeregistration string `json:"onDeregistration,omitempty" yaml:"onDeregistration,omitempty"`

	// Indicates how the GWLB handles existing flows when a target is unhealthy. Possible values are `rebalance` and `no_rebalance`. Must match the attribute value set for `on_deregistration`. Default: `no_rebalance`.
	OnUnhealthy string `json:"onUnhealthy,omitempty" yaml:"onUnhealthy,omitempty"`
}
