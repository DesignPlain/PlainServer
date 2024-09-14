package vertex

import types "libds/gcp/types"

type AiDataset struct {
	/*
	   Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
	   Structure is documented below.
	*/
	EncryptionSpec types.Vertex_AiDatasetEncryptionSpec `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`

	/*
	   A set of key/value label pairs to assign to this Workflow.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.


	   - - -
	*/
	MetadataSchemaUri string `json:"metadataSchemaUri,omitempty" yaml:"metadataSchemaUri,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the dataset. eg us-central1
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The user-defined name of the Dataset. The name can be up to 128 characters long and can be consist of any UTF-8 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
