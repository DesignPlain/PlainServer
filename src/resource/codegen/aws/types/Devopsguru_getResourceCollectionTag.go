package types

type Devopsguru_getResourceCollectionTag struct {
	// An AWS tag key that is used to identify the AWS resources that DevOps Guru analyzes.
	AppBoundaryKey string `json:"appBoundaryKey,omitempty" yaml:"appBoundaryKey,omitempty"`

	// Array of tag values.
	TagValues []string `json:"tagValues,omitempty" yaml:"tagValues,omitempty"`
}
