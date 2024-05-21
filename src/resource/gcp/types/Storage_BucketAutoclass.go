package types

type Storage_BucketAutoclass struct {
	// While set to `true`, autoclass automatically transitions objects in your bucket to appropriate storage classes based on each object's access pattern.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The storage class that objects in the bucket eventually transition to if they are not read for a certain length of time. Supported values include: `NEARLINE`, `ARCHIVE`.
	TerminalStorageClass string `json:"terminalStorageClass,omitempty" yaml:"terminalStorageClass,omitempty"`
}
