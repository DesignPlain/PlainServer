package types

type Gkeonprem_BareMetalClusterLoadBalancerPortConfig struct {
	// The port that control plane hosted load balancers will listen on.
	ControlPlaneLoadBalancerPort int `json:"controlPlaneLoadBalancerPort,omitempty" yaml:"controlPlaneLoadBalancerPort,omitempty"`
}
