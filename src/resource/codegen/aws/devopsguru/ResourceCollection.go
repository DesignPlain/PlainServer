package devopsguru

import types "libds/aws/types"

type ResourceCollection struct {
	/*
	   Type of AWS resource collection to create. Valid values are `AWS_CLOUD_FORMATION`, `AWS_SERVICE`, and `AWS_TAGS`.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// A collection of AWS CloudFormation stacks. See `cloudformation` below for additional details.
	Cloudformation types.Devopsguru_ResourceCollectionCloudformation `json:"cloudformation,omitempty" yaml:"cloudformation,omitempty"`

	// AWS tags used to filter the resources in the resource collection. See `tags` below for additional details.
	Tags types.Devopsguru_ResourceCollectionTags `json:"tags,omitempty" yaml:"tags,omitempty"`
}
