package types

type Gkeonprem_BareMetalAdminClusterLoadBalancerPortConfig struct {
	// The port that control plane hosted load balancers will listen on.
	ControlPlaneLoadBalancerPort int `json:"controlPlaneLoadBalancerPort,omitempty" yaml:"controlPlaneLoadBalancerPort,omitempty"`
}
