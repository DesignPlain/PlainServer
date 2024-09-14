package types

type Composer_getEnvironmentStorageConfig struct {
	// Optional. Name of an existing Cloud Storage bucket to be used by the environment.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
