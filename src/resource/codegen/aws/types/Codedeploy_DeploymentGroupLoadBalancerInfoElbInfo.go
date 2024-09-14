package types

type Codedeploy_DeploymentGroupLoadBalancerInfoElbInfo struct {
	// The name of the load balancer that will be used to route traffic from original instances to replacement instances in a blue/green deployment. For in-place deployments, the name of the load balancer that instances are deregistered from so they are not serving traffic during a deployment, and then re-registered with after the deployment completes.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
