package vertex

import types "DesignSphere_Server/src/resource/gcp/types"

type AiMetadataStore struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the Metadata Store. eg us-central1
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Description of the MetadataStore.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Customer-managed encryption key spec for a MetadataStore. If set, this MetadataStore and all sub-resources of this MetadataStore will be secured by this key.
	   Structure is documented below.
	*/
	EncryptionSpec types.Vertex_AiMetadataStoreEncryptionSpec `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`

	// The name of the MetadataStore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
