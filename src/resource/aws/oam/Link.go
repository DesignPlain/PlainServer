package oam

type Link struct {
	// Types of data that the source account shares with the monitoring account.
	ResourceTypes []string `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`

	/*
	   Identifier of the sink to use to create this link.

	   The following arguments are optional:
	*/
	SinkIdentifier string `json:"sinkIdentifier,omitempty" yaml:"sinkIdentifier,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
	LabelTemplate string `json:"labelTemplate,omitempty" yaml:"labelTemplate,omitempty"`
}
