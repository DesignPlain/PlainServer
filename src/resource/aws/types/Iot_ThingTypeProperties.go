package types

type Iot_ThingTypeProperties struct {
	// The description of the thing type.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A list of searchable thing attribute names.
	SearchableAttributes []string `json:"searchableAttributes,omitempty" yaml:"searchableAttributes,omitempty"`
}
