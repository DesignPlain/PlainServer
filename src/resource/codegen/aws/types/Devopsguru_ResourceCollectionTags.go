package types

type Devopsguru_ResourceCollectionTags struct {
	// An AWS tag key that is used to identify the AWS resources that DevOps Guru analyzes. All AWS resources in your account and Region tagged with this key make up your DevOps Guru application and analysis boundary. The key must begin with the prefix `DevOps-Guru-`. Any casing can be used for the prefix, but the associated tags __must use the same casing__ in their tag key.
	AppBoundaryKey string `json:"appBoundaryKey,omitempty" yaml:"appBoundaryKey,omitempty"`

	// Array of tag values. These can be used to further filter for specific resources within the application boundary. To analyze all resources tagged with the `app_boundary_key` regardless of the corresponding tag value, this array should be a single item containing a wildcard (`"-"`).
	TagValues []string `json:"tagValues,omitempty" yaml:"tagValues,omitempty"`
}
