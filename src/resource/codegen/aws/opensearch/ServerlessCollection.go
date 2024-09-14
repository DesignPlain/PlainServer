package opensearch

import types "libds/aws/types"

type ServerlessCollection struct {
	// Type of collection. One of `SEARCH`, `TIMESERIES`, or `VECTORSEARCH`. Defaults to `TIMESERIES`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Description of the collection.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the collection.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Indicates whether standby replicas should be used for a collection. One of `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
	StandbyReplicas string `json:"standbyReplicas,omitempty" yaml:"standbyReplicas,omitempty"`

	// A map of tags to assign to the collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Opensearch_ServerlessCollectionTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
