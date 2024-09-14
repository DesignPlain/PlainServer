package lightsail

type StaticIpAttachment struct {
	// The name of the allocated static IP
	StaticIpName string `json:"staticIpName,omitempty" yaml:"staticIpName,omitempty"`

	// The name of the Lightsail instance to attach the IP to
	InstanceName string `json:"instanceName,omitempty" yaml:"instanceName,omitempty"`
}
