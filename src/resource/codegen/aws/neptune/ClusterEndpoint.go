package neptune

type ClusterEndpoint struct {
	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty.
	ExcludedMembers []string `json:"excludedMembers,omitempty" yaml:"excludedMembers,omitempty"`

	// List of DB instance identifiers that are part of the custom endpoint group.
	StaticMembers []string `json:"staticMembers,omitempty" yaml:"staticMembers,omitempty"`

	// A map of tags to assign to the Neptune cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The identifier of the endpoint.
	ClusterEndpointIdentifier string `json:"clusterEndpointIdentifier,omitempty" yaml:"clusterEndpointIdentifier,omitempty"`

	// The DB cluster identifier of the DB cluster associated with the endpoint.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The type of the endpoint. One of: `READER`, `WRITER`, `ANY`.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`
}
