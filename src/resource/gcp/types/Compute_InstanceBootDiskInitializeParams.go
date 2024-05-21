package types

type Compute_InstanceBootDiskInitializeParams struct {
	/*
	   The size of the image in gigabytes. If not specified, it
	   will inherit the size of its base image.
	*/
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// The GCE disk type. Such as pd-standard, pd-balanced or pd-ssd.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   Whether this disk is using confidential compute mode.
	   Note: Only supported on hyperdisk skus, disk_encryption_key is required when setting to true.
	*/
	EnableConfidentialCompute bool `json:"enableConfidentialCompute,omitempty" yaml:"enableConfidentialCompute,omitempty"`

	/*
	   The image from which to initialize this disk. This can be
	   one of: the image's `self_link`, `projects/{project}/global/images/{image}`,
	   `projects/{project}/global/images/family/{family}`, `global/images/{image}`,
	   `global/images/family/{family}`, `family/{family}`, `{project}/{family}`,
	   `{project}/{image}`, `{family}`, or `{image}`. If referred by family, the
	   images names must include the family name. If they don't, use the
	   [gcp.compute.Image data source](https://www.terraform.io/docs/providers/google/d/compute_image.html).
	   For instance, the image `centos-6-v20180104` includes its family name `centos-6`.
	   These images can be referred by family name here.
	*/
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	/*
	   A set of key/value label pairs assigned to the disk. This
	   field is only applicable for persistent disks.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Indicates how many IOPS to provision for the disk.
	   This sets the number of I/O operations per second that the disk can handle.
	   For more details,see the [Hyperdisk documentation](https://cloud.google.com/compute/docs/disks/hyperdisks).
	   Note: Updating currently is only supported for hyperdisk skus via disk update
	   api/gcloud without the need to delete and recreate the disk, hyperdisk allows
	   for an update of IOPS every 4 hours. To update your hyperdisk more frequently,
	   you'll need to manually delete and recreate it.
	*/
	ProvisionedIops int `json:"provisionedIops,omitempty" yaml:"provisionedIops,omitempty"`

	/*
	   Indicates how much throughput to provision for the disk.
	   This sets the number of throughput mb per second that the disk can handle.
	   For more details,see the [Hyperdisk documentation](https://cloud.google.com/compute/docs/disks/hyperdisks).
	   Note: Updating currently is only supported for hyperdisk skus via disk update
	   api/gcloud without the need to delete and recreate the disk, hyperdisk allows
	   for an update of throughput every 4 hours. To update your hyperdisk more
	   frequently, you'll need to manually delete and recreate it.
	*/
	ProvisionedThroughput int `json:"provisionedThroughput,omitempty" yaml:"provisionedThroughput,omitempty"`

	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty" yaml:"resourceManagerTags,omitempty"`
}
