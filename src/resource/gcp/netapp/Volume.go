package netapp

import types "DesignSphere_Server/src/resource/gcp/types"

type Volume struct {
	// Capacity of the volume (in GiB).
	CapacityGib string `json:"capacityGib,omitempty" yaml:"capacityGib,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Export policy of the volume for NFSV3 and/or NFSV4.1 access.
	   Structure is documented below.
	*/
	ExportPolicy types.Netapp_VolumeExportPolicy `json:"exportPolicy,omitempty" yaml:"exportPolicy,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Share name (SMB) or export path (NFS) of the volume. Needs to be unique per location.
	ShareName string `json:"shareName,omitempty" yaml:"shareName,omitempty"`

	/*
	   Security Style of the Volume. Use UNIX to use UNIX or NFSV4 ACLs for file permissions.
	   Use NTFS to use NTFS ACLs for file permissions. Can only be set for volumes which use SMB together with NFS as protocol.
	   Possible values are: `NTFS`, `UNIX`.
	*/
	SecurityStyle string `json:"securityStyle,omitempty" yaml:"securityStyle,omitempty"`

	// If enabled, a NFS volume will contain a read-only .snapshot directory which provides access to each of the volume's snapshots. Will enable "Previous Versions" support for SMB.
	SnapshotDirectory bool `json:"snapshotDirectory,omitempty" yaml:"snapshotDirectory,omitempty"`

	// Name of the storage pool to create the volume in. Pool needs enough spare capacity to accomodate the volume.
	StoragePool string `json:"storagePool,omitempty" yaml:"storagePool,omitempty"`

	/*
	   The protocol of the volume. Allowed combinations are `['NFSV3']`, `['NFSV4']`, `['SMB']`, `['NFSV3', 'NFSV4']`, `['SMB', 'NFSV3']` and `['SMB', 'NFSV4']`.
	   Each value may be one of: `NFSV3`, `NFSV4`, `SMB`.
	*/
	Protocols []string `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	/*
	   List of actions that are restricted on this volume.
	   Each value may be one of: `DELETE`.
	*/
	RestrictedActions []string `json:"restrictedActions,omitempty" yaml:"restrictedActions,omitempty"`

	/*
	   Settings for volumes with SMB access.
	   Each value may be one of: `ENCRYPT_DATA`, `BROWSABLE`, `CHANGE_NOTIFY`, `NON_BROWSABLE`, `OPLOCKS`, `SHOW_SNAPSHOT`, `SHOW_PREVIOUS_VERSIONS`, `ACCESS_BASED_ENUMERATION`, `CONTINUOUSLY_AVAILABLE`.
	*/
	SmbSettings []string `json:"smbSettings,omitempty" yaml:"smbSettings,omitempty"`

	/*
	   Snapshot policy defines the schedule for automatic snapshot creation.
	   To disable automatic snapshot creation you have to remove the whole snapshot_policy block.
	   Structure is documented below.
	*/
	SnapshotPolicy types.Netapp_VolumeSnapshotPolicy `json:"snapshotPolicy,omitempty" yaml:"snapshotPolicy,omitempty"`

	/*
	   Policy to determine if the volume should be deleted forcefully.
	   Volumes may have nested snapshot resources. Deleting such a volume will fail.
	   Setting this parameter to FORCE will delete volumes including nested snapshots.
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	// Flag indicating if the volume is a kerberos volume or not, export policy rules control kerberos security modes (krb5, krb5i, krb5p).
	KerberosEnabled bool `json:"kerberosEnabled,omitempty" yaml:"kerberosEnabled,omitempty"`

	/*
	   Labels as key value pairs. Example: `{ "owner": "Bob", "department": "finance", "purpose": "testing" }`.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Name of the pool location. Usually a region name, expect for some STANDARD service level pools which require a zone name.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The name of the volume. Needs to be unique per location.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Used to create this volume from a snapshot (= cloning) or an backup.
	   Structure is documented below.
	*/
	RestoreParameters types.Netapp_VolumeRestoreParameters `json:"restoreParameters,omitempty" yaml:"restoreParameters,omitempty"`

	// Unix permission the mount point will be created with. Default is 0770. Applicable for UNIX security style volumes only.
	UnixPermissions string `json:"unixPermissions,omitempty" yaml:"unixPermissions,omitempty"`
}
