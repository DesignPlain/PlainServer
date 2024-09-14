package oam

import types "libds/aws/types"

type Link struct {
	// Human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.
	LabelTemplate string `json:"labelTemplate,omitempty" yaml:"labelTemplate,omitempty"`

	// Configuration for creating filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account. See `link_configuration` Block for details.
	LinkConfiguration types.Oam_LinkLinkConfiguration `json:"linkConfiguration,omitempty" yaml:"linkConfiguration,omitempty"`

	// Types of data that the source account shares with the monitoring account.
	ResourceTypes []string `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`

	/*
	   Identifier of the sink to use to create this link.

	   The following arguments are optional:
	*/
	SinkIdentifier string `json:"sinkIdentifier,omitempty" yaml:"sinkIdentifier,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
