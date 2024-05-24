package servicediscovery

type Instance struct {
	// A map contains the attributes of the instance. Check the [doc](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html#API_RegisterInstance_RequestSyntax) for the supported attributes and syntax.
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	// The ID of the service instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// The ID of the service that you want to use to create the instance.
	ServiceId string `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`
}
