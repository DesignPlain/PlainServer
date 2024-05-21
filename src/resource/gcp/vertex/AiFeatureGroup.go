package vertex

import types "DesignSphere_Server/src/resource/gcp/types"

type AiFeatureGroup struct {
	// The resource name of the Feature Group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of feature group. eg us-central1
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Indicates that features for this group come from BigQuery Table/View. By default treats the source as a sparse time series source, which is required to have an entityId and a feature_timestamp column in the source.
	   Structure is documented below.
	*/
	BigQuery types.Vertex_AiFeatureGroupBigQuery `json:"bigQuery,omitempty" yaml:"bigQuery,omitempty"`

	// The description of the FeatureGroup.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The labels with user-defined metadata to organize your FeatureGroup.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
