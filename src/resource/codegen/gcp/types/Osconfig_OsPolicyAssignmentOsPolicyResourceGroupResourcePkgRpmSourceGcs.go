package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgRpmSourceGcs struct {
	// Bucket of the Cloud Storage object.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Generation number of the Cloud Storage object.
	Generation int `json:"generation,omitempty" yaml:"generation,omitempty"`

	// Name of the Cloud Storage object.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`
}
