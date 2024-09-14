package logging

type FolderSettings struct {
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly provided.
	StorageLocation string `json:"storageLocation,omitempty" yaml:"storageLocation,omitempty"`

	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The _Default sink can be re-enabled manually if needed.
	DisableDefaultSink bool `json:"disableDefaultSink,omitempty" yaml:"disableDefaultSink,omitempty"`

	/*
	   The folder for which to retrieve settings.


	   - - -
	*/
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	// The resource name for the configured Cloud KMS key.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`
}
