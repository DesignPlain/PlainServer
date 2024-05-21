package types

type Networkservices_EndpointPolicyEndpointMatcherMetadataLabelMatcher struct {
	/*
	   Specifies how matching should be done.
	   Possible values are: `MATCH_ANY`, `MATCH_ALL`.
	*/
	MetadataLabelMatchCriteria string `json:"metadataLabelMatchCriteria,omitempty" yaml:"metadataLabelMatchCriteria,omitempty"`

	/*
	   The list of label value pairs that must match labels in the provided metadata based on filterMatchCriteria
	   Structure is documented below.
	*/
	MetadataLabels []Networkservices_EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabel `json:"metadataLabels,omitempty" yaml:"metadataLabels,omitempty"`
}
