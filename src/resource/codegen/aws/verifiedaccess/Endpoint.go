package verifiedaccess

import types "libds/aws/types"

type Endpoint struct {
	// The policy document that is associated with this resource.
	PolicyDocument string `json:"policyDocument,omitempty" yaml:"policyDocument,omitempty"`

	// Key-value tags for the Verified Access Endpoint. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The DNS name for users to reach your application.
	ApplicationDomain string `json:"applicationDomain,omitempty" yaml:"applicationDomain,omitempty"`

	// A description for the Verified Access endpoint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ARN of the public TLS/SSL certificate in AWS Certificate Manager to associate with the endpoint. The CN in the certificate must match the DNS name your end users will use to reach your application.
	DomainCertificateArn string `json:"domainCertificateArn,omitempty" yaml:"domainCertificateArn,omitempty"`

	// The network interface details. This parameter is required if the endpoint type is `network-interface`.
	NetworkInterfaceOptions types.Verifiedaccess_EndpointNetworkInterfaceOptions `json:"networkInterfaceOptions,omitempty" yaml:"networkInterfaceOptions,omitempty"`

	// List of the the security groups IDs to associate with the Verified Access endpoint.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The options in use for server side encryption.
	SseSpecification types.Verifiedaccess_EndpointSseSpecification `json:"sseSpecification,omitempty" yaml:"sseSpecification,omitempty"`

	/*
	   The ID of the Verified Access group to associate the endpoint with.

	   The following arguments are optional:
	*/
	VerifiedAccessGroupId string `json:"verifiedAccessGroupId,omitempty" yaml:"verifiedAccessGroupId,omitempty"`

	// The type of attachment. Currently, only `vpc` is supported.
	AttachmentType string `json:"attachmentType,omitempty" yaml:"attachmentType,omitempty"`

	// A custom identifier that is prepended to the DNS name that is generated for the endpoint.
	EndpointDomainPrefix string `json:"endpointDomainPrefix,omitempty" yaml:"endpointDomainPrefix,omitempty"`

	// The type of Verified Access endpoint to create. Currently `load-balancer` or `network-interface` are supported.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`

	// The load balancer details. This parameter is required if the endpoint type is `load-balancer`.
	LoadBalancerOptions types.Verifiedaccess_EndpointLoadBalancerOptions `json:"loadBalancerOptions,omitempty" yaml:"loadBalancerOptions,omitempty"`
}
