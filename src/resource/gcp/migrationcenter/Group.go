package migrationcenter

type Group struct {
	/*
	   Labels as key value pairs.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location of the group.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Optional. The description of the group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Optional. User-friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Required. User specified ID for the group. It will become the last component of the group name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression: `a-z?`.


	   - - -
	*/
	GroupId string `json:"groupId,omitempty" yaml:"groupId,omitempty"`
}
