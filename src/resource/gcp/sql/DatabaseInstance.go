package sql

import types "DesignSphere_Server/src/resource/gcp/types"

type DatabaseInstance struct {
	/*
	   The configuration for replication. The
	   configuration is detailed below. Valid only for MySQL instances.
	*/
	ReplicaConfiguration types.Sql_DatabaseInstanceReplicaConfiguration `json:"replicaConfiguration,omitempty" yaml:"replicaConfiguration,omitempty"`

	/*
	   The context needed to restore the database to a backup run. This field will
	   cause the provider to trigger the database to restore from the backup run indicated. The configuration is detailed below.
	   --NOTE:-- Restoring from a backup is an imperative action and not recommended via this provider. Adding or modifying this
	   block during resource creation/update will trigger the restore action after the resource is created/updated.
	*/
	RestoreBackupContext types.Sql_DatabaseInstanceRestoreBackupContext `json:"restoreBackupContext,omitempty" yaml:"restoreBackupContext,omitempty"`

	// The current software version on the instance. This attribute can not be set during creation. Refer to `available_maintenance_versions` attribute to see what `maintenance_version` are available for upgrade. When this attribute gets updated, it will cause an instance restart. Setting a `maintenance_version` value that is older than the current one on the instance will be ignored.
	MaintenanceVersion string `json:"maintenanceVersion,omitempty" yaml:"maintenanceVersion,omitempty"`

	/*
	   The name of the existing instance that will
	   act as the master in the replication setup. Note, this requires the master to
	   have `binary_log_enabled` set, as well as existing backups.
	*/
	MasterInstanceName string `json:"masterInstanceName,omitempty" yaml:"masterInstanceName,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The context needed to create this instance as a clone of another instance. When this field is set during
	   resource creation, this provider will attempt to clone another instance as indicated in the context. The
	   configuration is detailed below.
	*/
	Clone types.Sql_DatabaseInstanceClone `json:"clone,omitempty" yaml:"clone,omitempty"`

	/*
	   Whether or not to allow the provider to destroy the instance. Unless this field is set to false
	   in state, a `destroy` or `update` command that deletes the instance will fail. Defaults to `true`.
	*/
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// Initial root password. Can be updated. Required for MS SQL Server.
	RootPassword string `json:"rootPassword,omitempty" yaml:"rootPassword,omitempty"`

	/*
	   The settings to use for the database. The
	   configuration is detailed below. Required if `clone` is not set.
	*/
	Settings types.Sql_DatabaseInstanceSettings `json:"settings,omitempty" yaml:"settings,omitempty"`

	/*
	   The full path to the encryption key used for the CMEK disk encryption.  Setting
	   up disk encryption currently requires manual steps outside of this provider.
	   The provided key must be in the same region as the SQL instance.  In order
	   to use this feature, a special kind of service account must be created and
	   granted permission on this key.  This step can currently only be done
	   manually, please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#service-account).
	   That service account needs the `Cloud KMS > Cloud KMS CryptoKey Encrypter/Decrypter` role on your
	   key - please see [this step](https://cloud.google.com/sql/docs/mysql/configure-cmek#grantkey).
	*/
	EncryptionKeyName string `json:"encryptionKeyName,omitempty" yaml:"encryptionKeyName,omitempty"`

	// The type of the instance. The supported values are `SQL_INSTANCE_TYPE_UNSPECIFIED`, `CLOUD_SQL_INSTANCE`, `ON_PREMISES_INSTANCE` and `READ_REPLICA_INSTANCE`.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	/*
	   The region the instance will sit in. If a region is not provided in the resource definition,
	   the provider region will be used instead.

	   - - -
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The MySQL, PostgreSQL or
	   SQL Server version to use. Supported values include `MYSQL_5_6`,
	   `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`,`POSTGRES_10`, `POSTGRES_11`,
	   `POSTGRES_12`, `POSTGRES_13`, `POSTGRES_14`, `POSTGRES_15`, `SQLSERVER_2017_STANDARD`,
	   `SQLSERVER_2017_ENTERPRISE`, `SQLSERVER_2017_EXPRESS`, `SQLSERVER_2017_WEB`.
	   `SQLSERVER_2019_STANDARD`, `SQLSERVER_2019_ENTERPRISE`, `SQLSERVER_2019_EXPRESS`,
	   `SQLSERVER_2019_WEB`.
	   [Database Version Policies](https://cloud.google.com/sql/docs/db-versions)
	   includes an up-to-date reference of supported versions.
	*/
	DatabaseVersion string `json:"databaseVersion,omitempty" yaml:"databaseVersion,omitempty"`

	/*
	   The name of the instance. If the name is left
	   blank, the provider will randomly generate one when the instance is first
	   created. This is done because after a name is used, it cannot be reused for
	   up to [one week](https://cloud.google.com/sql/docs/delete-instance).
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
