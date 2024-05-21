package filestore

type Snapshot struct {
	// A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The resource name of the filestore instance.


	   - - -
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   Resource labels to represent user-provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The resource name of the snapshot. The name must be unique within the specified instance.
	   The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
