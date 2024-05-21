package dataplex

import types "DesignSphere_Server/src/resource/gcp/types"

type Lake struct {
	// Optional. Description of the lake.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Optional. User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Optional. User-defined labels for the lake.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Optional. Settings to manage lake and Dataproc Metastore service instance association.
	Metastore types.Dataplex_LakeMetastore `json:"metastore,omitempty" yaml:"metastore,omitempty"`

	/*
	   The name of the lake.



	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
