package types

type Networkservices_EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabel struct {
	// Required. Label name presented as key in xDS Node Metadata.
	LabelName string `json:"labelName,omitempty" yaml:"labelName,omitempty"`

	/*
	   Required. Label value presented as value corresponding to the above key, in xDS Node Metadata.

	   - - -
	*/
	LabelValue string `json:"labelValue,omitempty" yaml:"labelValue,omitempty"`
}
