package types

type Datacatalog_TagTemplateField struct {
	// Whether this is a required field. Defaults to false.
	IsRequired bool `json:"isRequired,omitempty" yaml:"isRequired,omitempty"`

	/*
	   (Output)
	   The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The order of this field with respect to other fields in this tag template.
	   A higher value indicates a more important field. The value can be negative.
	   Multiple fields can have the same order, and field orders within a tag do not have to be sequential.
	*/
	Order int `json:"order,omitempty" yaml:"order,omitempty"`

	/*
	   The type of value this tag field can contain.
	   Structure is documented below.
	*/
	Type Datacatalog_TagTemplateFieldType `json:"type,omitempty" yaml:"type,omitempty"`

	// A description for this field.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The display name for this field.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The identifier for this object. Format specified above.
	FieldId string `json:"fieldId,omitempty" yaml:"fieldId,omitempty"`
}
