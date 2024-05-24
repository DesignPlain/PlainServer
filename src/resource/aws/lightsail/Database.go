package lightsail

type Database struct {
	// The name of the database snapshot created if skip final snapshot is false, which is the default value for that parameter.
	FinalSnapshotName string `json:"finalSnapshotName,omitempty" yaml:"finalSnapshotName,omitempty"`

	// Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`

	// The name to use for your new Lightsail database resource. Names be unique within each AWS Region in your Lightsail account.
	RelationalDatabaseName string `json:"relationalDatabaseName,omitempty" yaml:"relationalDatabaseName,omitempty"`

	// The name of the master database created when the Lightsail database resource is created.
	MasterDatabaseName string `json:"masterDatabaseName,omitempty" yaml:"masterDatabaseName,omitempty"`

	// The weekly time range during which system maintenance can occur on your new database. Must be in the ddd:hh24:mi-ddd:hh24:mi format. Specified in Coordinated Universal Time (UTC). Example: `Tue:17:00-Tue:17:30`
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// Determines whether a final database snapshot is created before your database is deleted. If true is specified, no database snapshot is created. If false is specified, a database snapshot is created before your database is deleted. You must specify the final relational database snapshot name parameter if the skip final snapshot parameter is false.
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" yaml:"skipFinalSnapshot,omitempty"`

	// When true , applies changes immediately. When false , applies changes during the preferred maintenance window. Some changes may cause an outage.
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// When true, enables automated backup retention for your database. When false, disables automated backup retention for your database. Disabling backup retention deletes all automated database backups. Before disabling this, you may want to create a snapshot of your database.
	BackupRetentionEnabled bool `json:"backupRetentionEnabled,omitempty" yaml:"backupRetentionEnabled,omitempty"`

	// The blueprint ID for your new database. A blueprint describes the major engine version of a database. You can get a list of database blueprints IDs by using the AWS CLI command: `aws lightsail get-relational-database-blueprints`
	BlueprintId string `json:"blueprintId,omitempty" yaml:"blueprintId,omitempty"`

	// The master user name for your new database.
	MasterUsername string `json:"masterUsername,omitempty" yaml:"masterUsername,omitempty"`

	// The daily time range during which automated backups are created for your new database if automated backups are enabled. Must be in the hh24:mi-hh24:mi format. Example: `16:00-16:30`. Specified in Coordinated Universal Time (UTC).
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" yaml:"preferredBackupWindow,omitempty"`

	// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The bundle ID for your new database. A bundle describes the performance specifications for your database (see list below). You can get a list of database bundle IDs by using the AWS CLI command: `aws lightsail get-relational-database-bundles`.
	BundleId string `json:"bundleId,omitempty" yaml:"bundleId,omitempty"`

	// The password for the master user of your new database. The password can include any printable ASCII character except "/", """, or "@".
	MasterPassword string `json:"masterPassword,omitempty" yaml:"masterPassword,omitempty"`
}
