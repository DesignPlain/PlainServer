package types

type Networkservices_EndpointPolicyEndpointMatcher struct {
	/*
	   The matcher is based on node metadata presented by xDS clients.
	   Structure is documented below.
	*/
	MetadataLabelMatcher Networkservices_EndpointPolicyEndpointMatcherMetadataLabelMatcher `json:"metadataLabelMatcher,omitempty" yaml:"metadataLabelMatcher,omitempty"`
}
