package vpclattice

type ServiceNetworkServiceAssociation struct {
	// The ID or Amazon Resource Identifier (ARN) of the service.
	ServiceIdentifier string `json:"serviceIdentifier,omitempty" yaml:"serviceIdentifier,omitempty"`

	/*
	   The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
	   The following arguments are optional:
	*/
	ServiceNetworkIdentifier string `json:"serviceNetworkIdentifier,omitempty" yaml:"serviceNetworkIdentifier,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
