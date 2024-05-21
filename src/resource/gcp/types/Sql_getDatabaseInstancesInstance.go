package types

type Sql_getDatabaseInstancesInstance struct {
	// The ID of the project in which the resources belong. If it is not provided, the provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	RestoreBackupContexts []Sql_getDatabaseInstancesInstanceRestoreBackupContext `json:"restoreBackupContexts,omitempty" yaml:"restoreBackupContexts,omitempty"`

	// The URI of the created resource.
	SelfLink string `json:"selfLink,omitempty" yaml:"selfLink,omitempty"`

	// The service account email address assigned to the instance.
	ServiceAccountEmailAddress string `json:"serviceAccountEmailAddress,omitempty" yaml:"serviceAccountEmailAddress,omitempty"`

	// Available Maintenance versions.
	AvailableMaintenanceVersions []string `json:"availableMaintenanceVersions,omitempty" yaml:"availableMaintenanceVersions,omitempty"`

	// To filter out the Cloud SQL instances which are of the specified database version.
	DatabaseVersion string `json:"databaseVersion,omitempty" yaml:"databaseVersion,omitempty"`

	//
	FirstIpAddress string `json:"firstIpAddress,omitempty" yaml:"firstIpAddress,omitempty"`

	//
	IpAddresses []Sql_getDatabaseInstancesInstanceIpAddress `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`

	// Maintenance version.
	MaintenanceVersion string `json:"maintenanceVersion,omitempty" yaml:"maintenanceVersion,omitempty"`

	// The name of the instance that will act as the master in the replication setup. Note, this requires the master to have binary_log_enabled set, as well as existing backups.
	MasterInstanceName string `json:"masterInstanceName,omitempty" yaml:"masterInstanceName,omitempty"`

	// The link to service attachment of PSC instance.
	PscServiceAttachmentLink string `json:"pscServiceAttachmentLink,omitempty" yaml:"pscServiceAttachmentLink,omitempty"`

	// The configuration for replication.
	ReplicaConfigurations []Sql_getDatabaseInstancesInstanceReplicaConfiguration `json:"replicaConfigurations,omitempty" yaml:"replicaConfigurations,omitempty"`

	// Initial root password. Required for MS SQL Server.
	RootPassword string `json:"rootPassword,omitempty" yaml:"rootPassword,omitempty"`

	// The settings to use for the database. The configuration is detailed below.
	Settings []Sql_getDatabaseInstancesInstanceSetting `json:"settings,omitempty" yaml:"settings,omitempty"`

	// The dns name of the instance.
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	// The type of the instance. The valid values are:- 'SQL_INSTANCE_TYPE_UNSPECIFIED', 'CLOUD_SQL_INSTANCE', 'ON_PREMISES_INSTANCE' and 'READ_REPLICA_INSTANCE'.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	//
	PrivateIpAddress string `json:"privateIpAddress,omitempty" yaml:"privateIpAddress,omitempty"`

	//
	PublicIpAddress string `json:"publicIpAddress,omitempty" yaml:"publicIpAddress,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// To filter out the Cloud SQL instances which are located in the specified region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	//
	ServerCaCerts []Sql_getDatabaseInstancesInstanceServerCaCert `json:"serverCaCerts,omitempty" yaml:"serverCaCerts,omitempty"`

	// Configuration for creating a new instance as a clone of another instance.
	Clones []Sql_getDatabaseInstancesInstanceClone `json:"clones,omitempty" yaml:"clones,omitempty"`

	// The connection name of the instance to be used in connection strings. For example, when connecting with Cloud SQL Proxy.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	//
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	//
	EncryptionKeyName string `json:"encryptionKeyName,omitempty" yaml:"encryptionKeyName,omitempty"`
}
