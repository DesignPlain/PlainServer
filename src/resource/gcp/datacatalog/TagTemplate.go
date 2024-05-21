package datacatalog

import types "DesignSphere_Server/src/resource/gcp/types"

type TagTemplate struct {
	// Template location region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The id of the tag template to create.
	TagTemplateId string `json:"tagTemplateId,omitempty" yaml:"tagTemplateId,omitempty"`

	// The display name for this template.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields. The change of field_id will be resulting in re-creating of field. The change of primitive_type will be resulting in re-creating of field, however if the field is a required, you cannot update it.
	   Structure is documented below.
	*/
	Fields []types.Datacatalog_TagTemplateField `json:"fields,omitempty" yaml:"fields,omitempty"`

	// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
