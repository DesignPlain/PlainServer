package s3

import types "libds/aws/types"

type Inventory struct {
	// Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
	Filter types.S3_InventoryFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Object versions to include in the inventory list. Valid values: `All`, `Current`.
	IncludedObjectVersions string `json:"includedObjectVersions,omitempty" yaml:"includedObjectVersions,omitempty"`

	// Unique identifier of the inventory configuration for the bucket.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of optional fields that are included in the inventory results. Please refer to the S3 [documentation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_InventoryConfiguration.html#AmazonS3-Type-InventoryConfiguration-OptionalFields) for more details.
	OptionalFields []string `json:"optionalFields,omitempty" yaml:"optionalFields,omitempty"`

	// Specifies the schedule for generating inventory results (documented below).
	Schedule types.S3_InventorySchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Name of the source bucket that inventory lists the objects for.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Contains information about where to publish the inventory results (documented below).
	Destination types.S3_InventoryDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Specifies whether the inventory is enabled or disabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
