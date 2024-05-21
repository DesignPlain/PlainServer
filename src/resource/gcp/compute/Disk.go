package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type Disk struct {
	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Indicates how much Throughput must be provisioned for the disk.
	   Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	   allows for an update of Throughput every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	*/
	ProvisionedThroughput int `json:"provisionedThroughput,omitempty" yaml:"provisionedThroughput,omitempty"`

	/*
	   Resource policies applied to this disk for automatic snapshot creations.
	   ~>--NOTE-- This value does not support updating the
	   resource policy, as resource policies can not be updated more than
	   one at a time. Use
	   `gcp.compute.DiskResourcePolicyAttachment`
	   to allow for updating the resource policy attached to the disk.
	*/
	ResourcePolicies []string `json:"resourcePolicies,omitempty" yaml:"resourcePolicies,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Whether this disk is using confidential compute mode.
	   Note: Only supported on hyperdisk skus, disk_encryption_key is required when setting to true
	*/
	EnableConfidentialCompute bool `json:"enableConfidentialCompute,omitempty" yaml:"enableConfidentialCompute,omitempty"`

	/*
	   A list of features to enable on the guest operating system.
	   Applicable only for bootable disks.
	   Structure is documented below.
	*/
	GuestOsFeatures []types.Compute_DiskGuestOsFeature `json:"guestOsFeatures,omitempty" yaml:"guestOsFeatures,omitempty"`

	// Any applicable license URI.
	Licenses []string `json:"licenses,omitempty" yaml:"licenses,omitempty"`

	/*
	   The customer-supplied encryption key of the source image. Required if
	   the source image is protected by a customer-supplied encryption key.
	   Structure is documented below.
	*/
	SourceImageEncryptionKey types.Compute_DiskSourceImageEncryptionKey `json:"sourceImageEncryptionKey,omitempty" yaml:"sourceImageEncryptionKey,omitempty"`

	/*
	   The customer-supplied encryption key of the source snapshot. Required
	   if the source snapshot is protected by a customer-supplied encryption
	   key.
	   Structure is documented below.
	*/
	SourceSnapshotEncryptionKey types.Compute_DiskSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty" yaml:"sourceSnapshotEncryptionKey,omitempty"`

	// A reference to the zone where the disk resides.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   Encrypts the disk using a customer-supplied encryption key.
	   After you encrypt a disk with a customer-supplied key, you must
	   provide the same key if you use the disk later (e.g. to create a disk
	   snapshot or an image, or to attach the disk to a virtual machine).
	   Customer-supplied encryption keys do not protect access to metadata of
	   the disk.
	   If you do not provide an encryption key when creating the disk, then
	   the disk will be encrypted using an automatically generated key and
	   you do not need to provide a key to use the disk later.
	   Structure is documented below.
	*/
	DiskEncryptionKey types.Compute_DiskDiskEncryptionKey `json:"diskEncryptionKey,omitempty" yaml:"diskEncryptionKey,omitempty"`

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
	   Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.

	   > --Warning:-- `interface` is deprecated and will be removed in a future major release. This field is no longer used and can be safely removed from your configurations; disk interfaces are automatically determined on attachment.
	*/
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`

	/*
	   Indicates how many IOPS must be provisioned for the disk.
	   Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	   allows for an update of IOPS every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	*/
	ProvisionedIops int `json:"provisionedIops,omitempty" yaml:"provisionedIops,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	AsyncPrimaryDisk types.Compute_DiskAsyncPrimaryDisk `json:"asyncPrimaryDisk,omitempty" yaml:"asyncPrimaryDisk,omitempty"`

	// Indicates whether or not the disk can be read/write attached to more than one instance.
	MultiWriter bool `json:"multiWriter,omitempty" yaml:"multiWriter,omitempty"`

	/*
	   The source disk used to create this disk. You can provide this as a partial or full URL to the resource.
	   For example, the following are valid values:
	   - https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/disks/{disk}
	   - https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/disks/{disk}
	   - projects/{project}/zones/{zone}/disks/{disk}
	   - projects/{project}/regions/{region}/disks/{disk}
	   - zones/{zone}/disks/{disk}
	   - regions/{region}/disks/{disk}
	*/
	SourceDisk string `json:"sourceDisk,omitempty" yaml:"sourceDisk,omitempty"`

	/*
	   URL of the disk type resource describing which disk type to use to
	   create the disk. Provide this when creating the disk.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   Labels to apply to this disk.  A list of key->value pairs.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Physical block size of the persistent disk, in bytes. If not present
	   in a request, a default value is used. Currently supported sizes
	   are 4096 and 16384, other sizes may be added in the future.
	   If an unsupported value is requested, the error message will list
	   the supported values for the caller's project.
	*/
	PhysicalBlockSizeBytes int `json:"physicalBlockSizeBytes,omitempty" yaml:"physicalBlockSizeBytes,omitempty"`

	/*
	   Size of the persistent disk, specified in GB. You can specify this
	   field when creating a persistent disk using the `image` or
	   `snapshot` parameter, or specify it alone to create an empty
	   persistent disk.
	   If you specify this field along with `image` or `snapshot`,
	   the value must not be less than the size of the image
	   or the size of the snapshot.
	   ~>--NOTE-- If you change the size, the provider updates the disk size
	   if upsizing is detected but recreates the disk if downsizing is requested.
	   You can add `lifecycle.prevent_destroy` in the config to prevent destroying
	   and recreating.
	*/
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	/*
	   The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the
	   snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:
	   - 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' -
	   'projects/project/global/snapshots/snapshot' - 'global/snapshots/snapshot' - 'snapshot'
	*/
	Snapshot string `json:"snapshot,omitempty" yaml:"snapshot,omitempty"`
}
