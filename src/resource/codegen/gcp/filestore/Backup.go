package filestore

type Backup struct {
	/*
	   Resource labels to represent user-provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The name of the location of the instance. This can be a region for ENTERPRISE tier instances.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The resource name of the backup. The name must be unique within the specified instance.
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

	// Name of the file share in the source Cloud Filestore instance that the backup is created from.
	SourceFileShare string `json:"sourceFileShare,omitempty" yaml:"sourceFileShare,omitempty"`

	// The resource name of the source Cloud Filestore instance, in the format projects/{projectId}/locations/{locationId}/instances/{instanceId}, used to create this backup.
	SourceInstance string `json:"sourceInstance,omitempty" yaml:"sourceInstance,omitempty"`

	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
