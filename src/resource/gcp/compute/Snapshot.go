package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type Snapshot struct {
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Labels to apply to this Snapshot.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Encrypts the snapshot using a customer-supplied encryption key.
	   After you encrypt a snapshot using a customer-supplied key, you must
	   provide the same key if you use the snapshot later. For example, you
	   must provide the encryption key when you create a disk from the
	   encrypted snapshot in a future request.
	   Customer-supplied encryption keys do not protect access to metadata of
	   the snapshot.
	   If you do not provide an encryption key when creating the snapshot,
	   then the snapshot will be encrypted using an automatically generated
	   key and you do not need to provide a key to use the snapshot later.
	   Structure is documented below.
	*/
	SnapshotEncryptionKey types.Compute_SnapshotSnapshotEncryptionKey `json:"snapshotEncryptionKey,omitempty" yaml:"snapshotEncryptionKey,omitempty"`

	/*
	   The customer-supplied encryption key of the source snapshot. Required
	   if the source snapshot is protected by a customer-supplied encryption
	   key.
	   Structure is documented below.
	*/
	SourceDiskEncryptionKey types.Compute_SnapshotSourceDiskEncryptionKey `json:"sourceDiskEncryptionKey,omitempty" yaml:"sourceDiskEncryptionKey,omitempty"`

	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations []string `json:"storageLocations,omitempty" yaml:"storageLocations,omitempty"`

	// A reference to the zone where the disk is hosted.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   Creates the new snapshot in the snapshot chain labeled with the
	   specified name. The chain name must be 1-63 characters long and
	   comply with RFC1035. This is an uncommon option only for advanced
	   service owners who needs to create separate snapshot chains, for
	   example, for chargeback tracking.  When you describe your snapshot
	   resource, this field is visible only if it has a non-empty value.
	*/
	ChainName string `json:"chainName,omitempty" yaml:"chainName,omitempty"`

	/*
	   Name of the resource; provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A reference to the disk used to create this snapshot.


	   - - -
	*/
	SourceDisk string `json:"sourceDisk,omitempty" yaml:"sourceDisk,omitempty"`
}
