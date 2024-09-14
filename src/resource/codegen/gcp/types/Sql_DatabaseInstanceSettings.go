package types

type Sql_DatabaseInstanceSettings struct {
	/*
	   This specifies when the instance should be
	   active. Can be either `ALWAYS`, `NEVER` or `ON_DEMAND`.
	*/
	ActivationPolicy string `json:"activationPolicy,omitempty" yaml:"activationPolicy,omitempty"`

	/*
	   The availability type of the Cloud SQL
	   instance, high availability (`REGIONAL`) or single zone (`ZONAL`).' For all instances, ensure that
	   `settings.backup_configuration.enabled` is set to `true`.
	   For MySQL instances, ensure that `settings.backup_configuration.binary_log_enabled` is set to `true`.
	   For Postgres and SQL Server instances, ensure that `settings.backup_configuration.point_in_time_recovery_enabled`
	   is set to `true`. Defaults to `ZONAL`.
	*/
	AvailabilityType string `json:"availabilityType,omitempty" yaml:"availabilityType,omitempty"`

	//
	ActiveDirectoryConfig Sql_DatabaseInstanceSettingsActiveDirectoryConfig `json:"activeDirectoryConfig,omitempty" yaml:"activeDirectoryConfig,omitempty"`

	// The type of data disk: PD_SSD or PD_HDD. Defaults to `PD_SSD`.
	DiskType string `json:"diskType,omitempty" yaml:"diskType,omitempty"`

	//
	LocationPreference Sql_DatabaseInstanceSettingsLocationPreference `json:"locationPreference,omitempty" yaml:"locationPreference,omitempty"`

	//
	PasswordValidationPolicy Sql_DatabaseInstanceSettingsPasswordValidationPolicy `json:"passwordValidationPolicy,omitempty" yaml:"passwordValidationPolicy,omitempty"`

	// The time_zone to be used by the database engine (supported only for SQL Server), in SQL Server timezone format.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	// Specifies if connections must use Cloud SQL connectors.
	ConnectorEnforcement string `json:"connectorEnforcement,omitempty" yaml:"connectorEnforcement,omitempty"`

	// Configuration to protect against accidental instance deletion.
	DeletionProtectionEnabled bool `json:"deletionProtectionEnabled,omitempty" yaml:"deletionProtectionEnabled,omitempty"`

	// The size of data disk, in GB. Size of a running instance cannot be reduced but can be increased. The minimum value is 10GB.
	DiskSize int `json:"diskSize,omitempty" yaml:"diskSize,omitempty"`

	// Configuration of Query Insights.
	InsightsConfig Sql_DatabaseInstanceSettingsInsightsConfig `json:"insightsConfig,omitempty" yaml:"insightsConfig,omitempty"`

	// Used to make sure changes to the settings block are atomic.
	Version int `json:"version,omitempty" yaml:"version,omitempty"`

	// The name of server instance collation.
	Collation string `json:"collation,omitempty" yaml:"collation,omitempty"`

	//
	DenyMaintenancePeriod Sql_DatabaseInstanceSettingsDenyMaintenancePeriod `json:"denyMaintenancePeriod,omitempty" yaml:"denyMaintenancePeriod,omitempty"`

	//
	IpConfiguration Sql_DatabaseInstanceSettingsIpConfiguration `json:"ipConfiguration,omitempty" yaml:"ipConfiguration,omitempty"`

	//
	BackupConfiguration Sql_DatabaseInstanceSettingsBackupConfiguration `json:"backupConfiguration,omitempty" yaml:"backupConfiguration,omitempty"`

	// Data cache configurations.
	DataCacheConfig Sql_DatabaseInstanceSettingsDataCacheConfig `json:"dataCacheConfig,omitempty" yaml:"dataCacheConfig,omitempty"`

	//
	DatabaseFlags []Sql_DatabaseInstanceSettingsDatabaseFlag `json:"databaseFlags,omitempty" yaml:"databaseFlags,omitempty"`

	// The maximum size to which storage capacity can be automatically increased. The default value is 0, which specifies that there is no limit.
	DiskAutoresizeLimit int `json:"diskAutoresizeLimit,omitempty" yaml:"diskAutoresizeLimit,omitempty"`

	// The edition of the instance, can be `ENTERPRISE` or `ENTERPRISE_PLUS`.
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`

	/*
	   The machine type to use. See [tiers](https://cloud.google.com/sql/docs/admin-api/v1beta4/tiers)
	   for more details and supported versions. Postgres supports only shared-core machine types,
	   and custom machine types such as `db-custom-2-13312`. See the [Custom Machine Type Documentation](https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#create) to learn about specifying custom machine types.
	*/
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`

	// A set of key/value user label pairs to assign to the instance.
	UserLabels map[string]string `json:"userLabels,omitempty" yaml:"userLabels,omitempty"`

	//
	AdvancedMachineFeatures Sql_DatabaseInstanceSettingsAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`

	//
	SqlServerAuditConfig Sql_DatabaseInstanceSettingsSqlServerAuditConfig `json:"sqlServerAuditConfig,omitempty" yaml:"sqlServerAuditConfig,omitempty"`

	// Enables auto-resizing of the storage size. Defaults to `true`.
	DiskAutoresize bool `json:"diskAutoresize,omitempty" yaml:"diskAutoresize,omitempty"`

	// Declares a one-hour maintenance window when an Instance can automatically restart to apply updates. The maintenance window is specified in UTC time.
	MaintenanceWindow Sql_DatabaseInstanceSettingsMaintenanceWindow `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	// Pricing plan for this instance, can only be `PER_USE`.
	PricingPlan string `json:"pricingPlan,omitempty" yaml:"pricingPlan,omitempty"`
}
