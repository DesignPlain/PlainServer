package types

type Globalaccelerator_EndpointGroupPortOverride struct {
	// The endpoint port that you want a listener port to be mapped to. This is the port on the endpoint, such as the Application Load Balancer or Amazon EC2 instance.
	EndpointPort int `json:"endpointPort,omitempty" yaml:"endpointPort,omitempty"`

	// The listener port that you want to map to a specific endpoint port. This is the port that user traffic arrives to the Global Accelerator on.
	ListenerPort int `json:"listenerPort,omitempty" yaml:"listenerPort,omitempty"`
}
