package types

type Compute_InstanceFromTemplateBootDiskInitializeParams struct {
	// The size of the image in gigabytes.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// A flag to enable confidential compute mode on boot disk
	EnableConfidentialCompute bool `json:"enableConfidentialCompute,omitempty" yaml:"enableConfidentialCompute,omitempty"`

	// The image from which this disk was initialised.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// A set of key/value label pairs assigned to the disk.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle.
	ProvisionedIops int `json:"provisionedIops,omitempty" yaml:"provisionedIops,omitempty"`

	// Indicates how much throughput to provision for the disk. This sets the number of throughput mb per second that the disk can handle.
	ProvisionedThroughput int `json:"provisionedThroughput,omitempty" yaml:"provisionedThroughput,omitempty"`

	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty" yaml:"resourceManagerTags,omitempty"`
}
