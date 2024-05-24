package types

type Resourcegroupstaggingapi_getResourcesTagFilter struct {
	// One part of a key-value pair that makes up a tag.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Optional part of a key-value pair that make up a tag.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
