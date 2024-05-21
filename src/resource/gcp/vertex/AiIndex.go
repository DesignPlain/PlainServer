package vertex

import types "DesignSphere_Server/src/resource/gcp/types"

type AiIndex struct {
	// The region of the index. eg us-central1
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The description of the Index.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The update method to use with this Index. The value must be the followings. If not set, BATCH_UPDATE will be used by default.
	   - BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.
	   - STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.
	*/
	IndexUpdateMethod string `json:"indexUpdateMethod,omitempty" yaml:"indexUpdateMethod,omitempty"`

	/*
	   The labels with user-defined metadata to organize your Indexes.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   An additional information about the Index
	   Structure is documented below.
	*/
	Metadata types.Vertex_AiIndexMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
