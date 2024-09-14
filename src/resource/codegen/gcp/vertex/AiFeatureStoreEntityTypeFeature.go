package vertex

type AiFeatureStoreEntityTypeFeature struct {
	// Description of the feature.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.


	   - - -
	*/
	Entitytype string `json:"entitytype,omitempty" yaml:"entitytype,omitempty"`

	/*
	   A set of key/value label pairs to assign to the feature.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
	ValueType string `json:"valueType,omitempty" yaml:"valueType,omitempty"`
}
