package codeartifact

type DomainPermissions struct {
	// The name of the domain on which to set the resource policy.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// The account number of the AWS account that owns the domain.
	DomainOwner string `json:"domainOwner,omitempty" yaml:"domainOwner,omitempty"`

	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument string `json:"policyDocument,omitempty" yaml:"policyDocument,omitempty"`

	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision string `json:"policyRevision,omitempty" yaml:"policyRevision,omitempty"`
}
