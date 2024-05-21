package vertex

type AiFeatureGroupFeature struct {
	// The name of the BigQuery Table/View column hosting data for this version. If no value is provided, will use featureId.
	VersionColumnName string `json:"versionColumnName,omitempty" yaml:"versionColumnName,omitempty"`

	// The description of the FeatureGroup.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the Feature Group.
	FeatureGroup string `json:"featureGroup,omitempty" yaml:"featureGroup,omitempty"`

	/*
	   The labels with user-defined metadata to organize your FeatureGroup.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The resource name of the Feature Group Feature.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The region for the resource. It should be the same as the feature group's region.


	   - - -
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
