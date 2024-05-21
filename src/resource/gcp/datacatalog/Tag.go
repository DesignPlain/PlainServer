package datacatalog

import types "DesignSphere_Server/src/resource/gcp/types"

type Tag struct {
	/*
	   Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
	   individual column based on that schema.
	   For attaching a tag to a nested column, use `.` to separate the column names. Example:
	   `outer_column.inner_column`
	*/
	Column string `json:"column,omitempty" yaml:"column,omitempty"`

	/*
	   This maps the ID of a tag field to the value of and additional information about that field.
	   Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.
	   Structure is documented below.
	*/
	Fields []types.Datacatalog_TagField `json:"fields,omitempty" yaml:"fields,omitempty"`

	/*
	   The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
	   all entries in that group.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   The resource name of the tag template that this tag uses. Example:
	   projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
	   This field cannot be modified after creation.
	*/
	Template string `json:"template,omitempty" yaml:"template,omitempty"`
}
