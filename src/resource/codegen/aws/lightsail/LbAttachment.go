package lightsail

type LbAttachment struct {
	// The name of the instance to attach to the load balancer.
	InstanceName string `json:"instanceName,omitempty" yaml:"instanceName,omitempty"`

	// The name of the Lightsail load balancer.
	LbName string `json:"lbName,omitempty" yaml:"lbName,omitempty"`
}
