package vertex

import types "DesignSphere_Server/src/resource/gcp/types"

type AiTensorboard struct {
	// The region of the tensorboard. eg us-central1
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Description of this Tensorboard.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   User provided name of this Tensorboard.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	   Structure is documented below.
	*/
	EncryptionSpec types.Vertex_AiTensorboardEncryptionSpec `json:"encryptionSpec,omitempty" yaml:"encryptionSpec,omitempty"`

	/*
	   The labels with user-defined metadata to organize your Tensorboards.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
