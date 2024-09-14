package dataexchange

type Revision struct {
	// An optional comment about the revision.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// The dataset id.
	DataSetId string `json:"dataSetId,omitempty" yaml:"dataSetId,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
