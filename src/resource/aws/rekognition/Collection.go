package rekognition

import types "DesignSphere_Server/src/resource/aws/types"

type Collection struct {
	/*
	   The name of the collection

	   The following arguments are optional:
	*/
	CollectionId string `json:"collectionId,omitempty" yaml:"collectionId,omitempty"`

	// A map of tags assigned to the WorkSpaces Connection Alias. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Rekognition_CollectionTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
