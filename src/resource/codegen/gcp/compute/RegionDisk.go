package compute

import types "libds/gcp/types"

type RegionDisk struct {
	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

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
	   field when creating a persistent disk using the sourceImage or
	   sourceSnapshot parameter, or specify it alone to create an empty
	   persistent disk.
	   If you specify this field along with sourceImage or sourceSnapshot,
	   the value of sizeGb must not be less than the size of the sourceImage
	   or the size of the snapshot.
	*/
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	/*
	   Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.

	   > --Warning:-- `interface` is deprecated and will be removed in a future major release. This field is no longer used and can be safely removed from your configurations; disk interfaces are automatically determined on attachment.
	*/
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For
	   example, the following are valid values: -
	   'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' -
	   'projects/project/global/snapshots/snapshot' - 'global/snapshots/snapshot' - 'snapshot'
	*/
	Snapshot string `json:"snapshot,omitempty" yaml:"snapshot,omitempty"`

	/*
	   A list of features to enable on the guest operating system.
	   Applicable only for bootable disks.
	   Structure is documented below.
	*/
	GuestOsFeatures []types.Compute_RegionDiskGuestOsFeature `json:"guestOsFeatures,omitempty" yaml:"guestOsFeatures,omitempty"`

	/*
	   Labels to apply to this disk.  A list of key->value pairs.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// A reference to the region where the disk resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

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
	   The customer-supplied encryption key of the source snapshot. Required
	   if the source snapshot is protected by a customer-supplied encryption
	   key.
	   Structure is documented below.
	*/
	SourceSnapshotEncryptionKey types.Compute_RegionDiskSourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty" yaml:"sourceSnapshotEncryptionKey,omitempty"`

	/*
	   URL of the disk type resource describing which disk type to use to
	   create the disk. Provide this when creating the disk.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	AsyncPrimaryDisk types.Compute_RegionDiskAsyncPrimaryDisk `json:"asyncPrimaryDisk,omitempty" yaml:"asyncPrimaryDisk,omitempty"`

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
	DiskEncryptionKey types.Compute_RegionDiskDiskEncryptionKey `json:"diskEncryptionKey,omitempty" yaml:"diskEncryptionKey,omitempty"`

	// Any applicable license URI.
	Licenses []string `json:"licenses,omitempty" yaml:"licenses,omitempty"`

	/*
	   URLs of the zones where the disk should be replicated to.


	   - - -
	*/
	ReplicaZones []string `json:"replicaZones,omitempty" yaml:"replicaZones,omitempty"`
}
