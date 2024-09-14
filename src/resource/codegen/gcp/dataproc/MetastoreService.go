package dataproc

import types "libds/gcp/types"

type MetastoreService struct {
	/*
	   Represents the scaling configuration of a metastore service.
	   Structure is documented below.
	*/
	ScalingConfig types.Dataproc_MetastoreServiceScalingConfig `json:"scalingConfig,omitempty" yaml:"scalingConfig,omitempty"`

	/*
	   The database type that the Metastore service stores its data.
	   Default value is `MYSQL`.
	   Possible values are: `MYSQL`, `SPANNER`.
	*/
	DatabaseType string `json:"databaseType,omitempty" yaml:"databaseType,omitempty"`

	/*
	   The setting that defines how metastore metadata should be integrated with external services and systems.
	   Structure is documented below.
	*/
	MetadataIntegration types.Dataproc_MetastoreServiceMetadataIntegration `json:"metadataIntegration,omitempty" yaml:"metadataIntegration,omitempty"`

	// The TCP port at which the metastore service is reached. Default: 9083.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   The release channel of the service. If unspecified, defaults to `STABLE`.
	   Default value is `STABLE`.
	   Possible values are: `CANARY`, `STABLE`.
	*/
	ReleaseChannel string `json:"releaseChannel,omitempty" yaml:"releaseChannel,omitempty"`

	/*
	   Information used to configure the Dataproc Metastore service to encrypt
	   customer data at rest.
	   Structure is documented below.
	*/
	EncryptionConfig types.Dataproc_MetastoreServiceEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	   and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	   3 and 63 characters.


	   - - -
	*/
	ServiceId string `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`

	/*
	   The tier of the service.
	   Possible values are: `DEVELOPER`, `ENTERPRISE`.
	*/
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`

	/*
	   The one hour maintenance window of the metastore service.
	   This specifies when the service can be restarted for maintenance purposes in UTC time.
	   Maintenance window is not needed for services with the `SPANNER` database type.
	   Structure is documented below.
	*/
	MaintenanceWindow types.Dataproc_MetastoreServiceMaintenanceWindow `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	/*
	   The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	   Structure is documented below.
	*/
	TelemetryConfig types.Dataproc_MetastoreServiceTelemetryConfig `json:"telemetryConfig,omitempty" yaml:"telemetryConfig,omitempty"`

	/*
	   The configuration specifying the network settings for the Dataproc Metastore service.
	   Structure is documented below.
	*/
	NetworkConfig types.Dataproc_MetastoreServiceNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	/*
	   Configuration information specific to running Hive metastore software as the metastore service.
	   Structure is documented below.
	*/
	HiveMetastoreConfig types.Dataproc_MetastoreServiceHiveMetastoreConfig `json:"hiveMetastoreConfig,omitempty" yaml:"hiveMetastoreConfig,omitempty"`

	/*
	   User-defined labels for the metastore service.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location where the metastore service should reside.
	   The default value is `global`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	   "projects/{projectNumber}/global/networks/{network_id}".
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
