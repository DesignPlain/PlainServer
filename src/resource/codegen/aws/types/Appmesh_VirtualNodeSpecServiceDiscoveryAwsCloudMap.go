package types

type Appmesh_VirtualNodeSpecServiceDiscoveryAwsCloudMap struct {
	// String map that contains attributes with values that you can use to filter instances by any custom attribute that you specified when you registered the instance. Only instances that match all of the specified key/value pairs will be returned.
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	/*
	   Name of the AWS Cloud Map namespace to use.
	   Use the `aws.servicediscovery.HttpNamespace` resource to configure a Cloud Map namespace. Must be between 1 and 1024 characters in length.
	*/
	NamespaceName string `json:"namespaceName,omitempty" yaml:"namespaceName,omitempty"`

	// Name of the AWS Cloud Map service to use. Use the `aws.servicediscovery.Service` resource to configure a Cloud Map service. Must be between 1 and 1024 characters in length.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`
}
