package rekognition

import types "libds/aws/types"

type Collection struct {
	// Map of tags assigned to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Rekognition_CollectionTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	/*
	   The name of the collection

	   The following arguments are optional:
	*/
	CollectionId string `json:"collectionId,omitempty" yaml:"collectionId,omitempty"`
}
