package lightsail

type Lb struct {
	//
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	// The name of the Lightsail load balancer.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The health check path of the load balancer. Default value "/".
	HealthCheckPath string `json:"healthCheckPath,omitempty" yaml:"healthCheckPath,omitempty"`

	// The instance port the load balancer will connect.
	InstancePort int `json:"instancePort,omitempty" yaml:"instancePort,omitempty"`
}
