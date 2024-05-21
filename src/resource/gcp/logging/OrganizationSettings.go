package logging

type OrganizationSettings struct {
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The _Default sink can be re-enabled manually if needed.
	DisableDefaultSink bool `json:"disableDefaultSink,omitempty" yaml:"disableDefaultSink,omitempty"`

	// The resource name for the configured Cloud KMS key.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   The organization for which to retrieve or configure settings.


	   - - -
	*/
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly provided.
	StorageLocation string `json:"storageLocation,omitempty" yaml:"storageLocation,omitempty"`
}
