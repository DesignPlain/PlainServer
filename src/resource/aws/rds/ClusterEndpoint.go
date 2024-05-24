package rds

type ClusterEndpoint struct {
	// List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excluded_members`.
	StaticMembers []string `json:"staticMembers,omitempty" yaml:"staticMembers,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
	ClusterEndpointIdentifier string `json:"clusterEndpointIdentifier,omitempty" yaml:"clusterEndpointIdentifier,omitempty"`

	// The cluster identifier.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The type of the endpoint. One of: READER , ANY .
	CustomEndpointType string `json:"customEndpointType,omitempty" yaml:"customEndpointType,omitempty"`

	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `static_members`.
	ExcludedMembers []string `json:"excludedMembers,omitempty" yaml:"excludedMembers,omitempty"`
}
