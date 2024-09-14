package types

type Storage_getBucketVersioning struct {
	// While set to true, versioning is fully enabled for this bucket.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
