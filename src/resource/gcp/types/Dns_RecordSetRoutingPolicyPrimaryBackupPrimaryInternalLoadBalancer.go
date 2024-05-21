package types

type Dns_RecordSetRoutingPolicyPrimaryBackupPrimaryInternalLoadBalancer struct {
	// The configured port of the load balancer.
	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	// The ID of the project in which the load balancer belongs.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the load balancer. Only needed for regional load balancers.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The frontend IP address of the load balancer.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
	IpProtocol string `json:"ipProtocol,omitempty" yaml:"ipProtocol,omitempty"`

	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb", "regionalL7ilb", "globalL7ilb"]
	LoadBalancerType string `json:"loadBalancerType,omitempty" yaml:"loadBalancerType,omitempty"`

	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like `projects/{project}/global/networks/{network}` or `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`.
	NetworkUrl string `json:"networkUrl,omitempty" yaml:"networkUrl,omitempty"`
}
