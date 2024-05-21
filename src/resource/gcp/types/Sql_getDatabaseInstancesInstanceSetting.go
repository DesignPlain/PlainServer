package types

type Sql_getDatabaseInstancesInstanceSetting struct {
	// The size of data disk, in GB. Size of a running instance cannot be reduced but can be increased. The minimum value is 10GB.
	DiskSize int `json:"diskSize,omitempty" yaml:"diskSize,omitempty"`

	// This specifies when the instance should be active. Can be either ALWAYS, NEVER or ON_DEMAND.
	ActivationPolicy string `json:"activationPolicy,omitempty" yaml:"activationPolicy,omitempty"`

	//
	ActiveDirectoryConfigs []Sql_getDatabaseInstancesInstanceSettingActiveDirectoryConfig `json:"activeDirectoryConfigs,omitempty" yaml:"activeDirectoryConfigs,omitempty"`

	// Data cache configurations.
	DataCacheConfigs []Sql_getDatabaseInstancesInstanceSettingDataCacheConfig `json:"dataCacheConfigs,omitempty" yaml:"dataCacheConfigs,omitempty"`

	//
	IpConfigurations []Sql_getDatabaseInstancesInstanceSettingIpConfiguration `json:"ipConfigurations,omitempty" yaml:"ipConfigurations,omitempty"`

	// Declares a one-hour maintenance window when an Instance can automatically restart to apply updates. The maintenance window is specified in UTC time.
	MaintenanceWindows []Sql_getDatabaseInstancesInstanceSettingMaintenanceWindow `json:"maintenanceWindows,omitempty" yaml:"maintenanceWindows,omitempty"`

	//
	PasswordValidationPolicies []Sql_getDatabaseInstancesInstanceSettingPasswordValidationPolicy `json:"passwordValidationPolicies,omitempty" yaml:"passwordValidationPolicies,omitempty"`

	/*
	   The availability type of the Cloud SQL instance, high availability
	   (REGIONAL) or single zone (ZONAL). For all instances, ensure that
	   settings.backup_configuration.enabled is set to true.
	   For MySQL instances, ensure that settings.backup_configuration.binary_log_enabled is set to true.
	   For Postgres instances, ensure that settings.backup_configuration.point_in_time_recovery_enabled
	   is set to true. Defaults to ZONAL.
	*/
	AvailabilityType string `json:"availabilityType,omitempty" yaml:"availabilityType,omitempty"`

	//
	DenyMaintenancePeriods []Sql_getDatabaseInstancesInstanceSettingDenyMaintenancePeriod `json:"denyMaintenancePeriods,omitempty" yaml:"denyMaintenancePeriods,omitempty"`

	// To filter out the Cloud SQL instances based on the tier(or machine type) of the database instances.
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`

	// Specifies if connections must use Cloud SQL connectors.
	ConnectorEnforcement string `json:"connectorEnforcement,omitempty" yaml:"connectorEnforcement,omitempty"`

	//
	AdvancedMachineFeatures []Sql_getDatabaseInstancesInstanceSettingAdvancedMachineFeature `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`

	// Configuration to protect against accidental instance deletion.
	DeletionProtectionEnabled bool `json:"deletionProtectionEnabled,omitempty" yaml:"deletionProtectionEnabled,omitempty"`

	// Pricing plan for this instance, can only be PER_USE.
	PricingPlan string `json:"pricingPlan,omitempty" yaml:"pricingPlan,omitempty"`

	// A set of key/value user label pairs to assign to the instance.
	UserLabels map[string]string `json:"userLabels,omitempty" yaml:"userLabels,omitempty"`

	// Used to make sure changes to the settings block are atomic.
	Version int `json:"version,omitempty" yaml:"version,omitempty"`

	// The type of data disk: PD_SSD or PD_HDD. Defaults to PD_SSD.
	DiskType string `json:"diskType,omitempty" yaml:"diskType,omitempty"`

	//
	BackupConfigurations []Sql_getDatabaseInstancesInstanceSettingBackupConfiguration `json:"backupConfigurations,omitempty" yaml:"backupConfigurations,omitempty"`

	// The name of server instance collation.
	Collation string `json:"collation,omitempty" yaml:"collation,omitempty"`

	//
	DatabaseFlags []Sql_getDatabaseInstancesInstanceSettingDatabaseFlag `json:"databaseFlags,omitempty" yaml:"databaseFlags,omitempty"`

	// The edition of the instance, can be ENTERPRISE or ENTERPRISE_PLUS.
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`

	// The time_zone to be used by the database engine (supported only for SQL Server), in SQL Server timezone format.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	//
	LocationPreferences []Sql_getDatabaseInstancesInstanceSettingLocationPreference `json:"locationPreferences,omitempty" yaml:"locationPreferences,omitempty"`

	// Enables auto-resizing of the storage size. Defaults to true.
	DiskAutoresize bool `json:"diskAutoresize,omitempty" yaml:"diskAutoresize,omitempty"`

	// The maximum size, in GB, to which storage capacity can be automatically increased. The default value is 0, which specifies that there is no limit.
	DiskAutoresizeLimit int `json:"diskAutoresizeLimit,omitempty" yaml:"diskAutoresizeLimit,omitempty"`

	// Configuration of Query Insights.
	InsightsConfigs []Sql_getDatabaseInstancesInstanceSettingInsightsConfig `json:"insightsConfigs,omitempty" yaml:"insightsConfigs,omitempty"`

	//
	SqlServerAuditConfigs []Sql_getDatabaseInstancesInstanceSettingSqlServerAuditConfig `json:"sqlServerAuditConfigs,omitempty" yaml:"sqlServerAuditConfigs,omitempty"`
}
