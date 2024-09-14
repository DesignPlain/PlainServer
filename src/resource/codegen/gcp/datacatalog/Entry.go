package datacatalog

import types "libds/gcp/types"

type Entry struct {
	// The name of the entry group this entry is in.
	EntryGroup string `json:"entryGroup,omitempty" yaml:"entryGroup,omitempty"`

	/*
	   The resource this metadata entry refers to.
	   For Google Cloud Platform resources, linkedResource is the full name of the resource.
	   For example, the linkedResource for a table resource from BigQuery is:
	   //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	   Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	   this field is optional and defaults to an empty string.
	*/
	LinkedResource string `json:"linkedResource,omitempty" yaml:"linkedResource,omitempty"`

	/*
	   Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	   attached to it. See
	   https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	   for what fields this schema can contain.
	*/
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`

	/*
	   The type of the entry. Only used for Entries with types in the EntryType enum.
	   Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
	   Possible values are: `FILESET`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	   When creating an entry, users should check the enum values first, if nothing matches the entry
	   to be created, then provide a custom value, for example "my_special_type".
	   userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	   numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	*/
	UserSpecifiedType string `json:"userSpecifiedType,omitempty" yaml:"userSpecifiedType,omitempty"`

	/*
	   Display information such as title and description. A short name to identify the entry,
	   for example, "Analytics Data - Jan 2011".
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The id of the entry to create.


	   - - -
	*/
	EntryId string `json:"entryId,omitempty" yaml:"entryId,omitempty"`

	/*
	   Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
	   Structure is documented below.
	*/
	GcsFilesetSpec types.Datacatalog_EntryGcsFilesetSpec `json:"gcsFilesetSpec,omitempty" yaml:"gcsFilesetSpec,omitempty"`

	/*
	   This field indicates the entry's source system that Data Catalog does not integrate with.
	   userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	   and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	*/
	UserSpecifiedSystem string `json:"userSpecifiedSystem,omitempty" yaml:"userSpecifiedSystem,omitempty"`

	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
