package alloydb

import types "libds/gcp/types"

type Cluster struct {
	/*
	   The type of cluster. If not set, defaults to PRIMARY.
	   Default value is `PRIMARY`.
	   Possible values are: `PRIMARY`, `SECONDARY`.
	*/
	ClusterType string `json:"clusterType,omitempty" yaml:"clusterType,omitempty"`

	/*
	   The continuous backup config for this cluster.
	   If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
	   Structure is documented below.
	*/
	ContinuousBackupConfig types.Alloydb_ClusterContinuousBackupConfig `json:"continuousBackupConfig,omitempty" yaml:"continuousBackupConfig,omitempty"`

	/*
	   Initial user to setup during cluster creation.
	   Structure is documented below.
	*/
	InitialUser types.Alloydb_ClusterInitialUser `json:"initialUser,omitempty" yaml:"initialUser,omitempty"`

	/*
	   Metadata related to network configuration.
	   Structure is documented below.
	*/
	NetworkConfig types.Alloydb_ClusterNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	/*
	   Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// The ID of the alloydb cluster.
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	// User-settable and human-readable display name for the Cluster.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// For Resource freshness validation (https://google.aip.dev/154)
	Etag string `json:"etag,omitempty" yaml:"etag,omitempty"`

	/*
	   The location where the alloydb cluster should reside.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The database engine major version. This is an optional field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.
	DatabaseVersion string `json:"databaseVersion,omitempty" yaml:"databaseVersion,omitempty"`

	/*
	   EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
	   Structure is documented below.
	*/
	EncryptionConfig types.Alloydb_ClusterEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	/*
	   User-defined labels for the alloydb cluster.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   (Optional, Deprecated)
	   The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	   "projects/{projectNumber}/global/networks/{network_id}".

	   > --Warning:-- `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The source when restoring from a backup. Conflicts with 'restore_continuous_backup_source', both can't be set together.
	   Structure is documented below.
	*/
	RestoreBackupSource types.Alloydb_ClusterRestoreBackupSource `json:"restoreBackupSource,omitempty" yaml:"restoreBackupSource,omitempty"`

	/*
	   The source when restoring via point in time recovery (PITR). Conflicts with 'restore_backup_source', both can't be set together.
	   Structure is documented below.
	*/
	RestoreContinuousBackupSource types.Alloydb_ClusterRestoreContinuousBackupSource `json:"restoreContinuousBackupSource,omitempty" yaml:"restoreContinuousBackupSource,omitempty"`

	/*
	   The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
	   Structure is documented below.
	*/
	AutomatedBackupPolicy types.Alloydb_ClusterAutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty" yaml:"automatedBackupPolicy,omitempty"`

	/*
	   Policy to determine if the cluster should be deleted forcefully.
	   Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
	   Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = "FORCE" otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	/*
	   Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY.
	   Structure is documented below.
	*/
	SecondaryConfig types.Alloydb_ClusterSecondaryConfig `json:"secondaryConfig,omitempty" yaml:"secondaryConfig,omitempty"`
}
