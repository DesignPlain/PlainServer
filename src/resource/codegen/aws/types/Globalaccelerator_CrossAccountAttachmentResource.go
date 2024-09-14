package types

type Globalaccelerator_CrossAccountAttachmentResource struct {
	// The AWS Region where a shared endpoint resource is located.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// IP address range, in CIDR format, that is specified as resource.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// The endpoint ID for the endpoint that is specified as a AWS resource.
	EndpointId string `json:"endpointId,omitempty" yaml:"endpointId,omitempty"`
}
