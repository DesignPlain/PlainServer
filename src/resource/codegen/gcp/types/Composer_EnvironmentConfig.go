package types

type Composer_EnvironmentConfig struct {
	// The URI of the Apache Airflow Web UI hosted within this environment.
	AirflowUri string `json:"airflowUri,omitempty" yaml:"airflowUri,omitempty"`

	// Optional. If true, a private Composer environment will be created.
	EnablePrivateEnvironment bool `json:"enablePrivateEnvironment,omitempty" yaml:"enablePrivateEnvironment,omitempty"`

	// The Kubernetes Engine cluster used to run this environment.
	GkeCluster string `json:"gkeCluster,omitempty" yaml:"gkeCluster,omitempty"`

	// The recovery configuration settings for the Cloud Composer environment
	RecoveryConfig Composer_EnvironmentConfigRecoveryConfig `json:"recoveryConfig,omitempty" yaml:"recoveryConfig,omitempty"`

	// The workloads configuration settings for the GKE cluster associated with the Cloud Composer environment. Supported for Cloud Composer environments in versions composer-2.-.--airflow--.-.- and newer.
	WorkloadsConfig Composer_EnvironmentConfigWorkloadsConfig `json:"workloadsConfig,omitempty" yaml:"workloadsConfig,omitempty"`

	// The configuration of Cloud SQL instance that is used by the Apache Airflow software. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	DatabaseConfig Composer_EnvironmentConfigDatabaseConfig `json:"databaseConfig,omitempty" yaml:"databaseConfig,omitempty"`

	// Optional. If true, builds performed during operations that install Python packages have only private connectivity to Google services. If false, the builds also have access to the internet.
	EnablePrivateBuildsOnly bool `json:"enablePrivateBuildsOnly,omitempty" yaml:"enablePrivateBuildsOnly,omitempty"`

	// The encryption options for the Composer environment and its dependencies.
	EncryptionConfig Composer_EnvironmentConfigEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	// The configuration used for the Private IP Cloud Composer environment.
	PrivateEnvironmentConfig Composer_EnvironmentConfigPrivateEnvironmentConfig `json:"privateEnvironmentConfig,omitempty" yaml:"privateEnvironmentConfig,omitempty"`

	// Network-level access control policy for the Airflow web server.
	WebServerNetworkAccessControl Composer_EnvironmentConfigWebServerNetworkAccessControl `json:"webServerNetworkAccessControl,omitempty" yaml:"webServerNetworkAccessControl,omitempty"`

	// Configuration options for the master authorized networks feature. Enabled master authorized networks will disallow all external traffic to access Kubernetes master through HTTPS except traffic from the given CIDR blocks, Google Compute Engine Public IPs and Google Prod IPs.
	MasterAuthorizedNetworksConfig Composer_EnvironmentConfigMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty" yaml:"masterAuthorizedNetworksConfig,omitempty"`

	// The configuration settings for software inside the environment.
	SoftwareConfig Composer_EnvironmentConfigSoftwareConfig `json:"softwareConfig,omitempty" yaml:"softwareConfig,omitempty"`

	// The configuration settings for the Airflow web server App Engine instance. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	WebServerConfig Composer_EnvironmentConfigWebServerConfig `json:"webServerConfig,omitempty" yaml:"webServerConfig,omitempty"`

	// The number of nodes in the Kubernetes Engine cluster that will be used to run this environment. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`

	// Whether high resilience is enabled or not. This field is supported for Cloud Composer environments in versions composer-2.1.15-airflow--.-.- and newer.
	ResilienceMode string `json:"resilienceMode,omitempty" yaml:"resilienceMode,omitempty"`

	// The Cloud Storage prefix of the DAGs for this environment. Although Cloud Storage objects reside in a flat namespace, a hierarchical file tree can be simulated using '/'-delimited object name prefixes. DAG objects for this environment reside in a simulated directory with this prefix.
	DagGcsPrefix string `json:"dagGcsPrefix,omitempty" yaml:"dagGcsPrefix,omitempty"`

	// The configuration setting for Airflow data retention mechanism. This field is supported for Cloud Composer environments in versions composer-2.0.32-airflow-2.1.4. or newer
	DataRetentionConfig Composer_EnvironmentConfigDataRetentionConfig `json:"dataRetentionConfig,omitempty" yaml:"dataRetentionConfig,omitempty"`

	// The size of the Cloud Composer environment. This field is supported for Cloud Composer environments in versions composer-2.-.--airflow--.-.- and newer.
	EnvironmentSize string `json:"environmentSize,omitempty" yaml:"environmentSize,omitempty"`

	// The configuration for Cloud Composer maintenance window.
	MaintenanceWindow Composer_EnvironmentConfigMaintenanceWindow `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	// The configuration used for the Kubernetes Engine cluster.
	NodeConfig Composer_EnvironmentConfigNodeConfig `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`
}
