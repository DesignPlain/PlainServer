package types

type Storage_BucketVersioning struct {
	// While set to `true`, versioning is fully enabled for this bucket.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
