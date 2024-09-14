package types

type Composer_getEnvironmentConfig struct {
	// The URI of the Apache Airflow Web UI hosted within this environment.
	AirflowUri string `json:"airflowUri,omitempty" yaml:"airflowUri,omitempty"`

	// The recovery configuration settings for the Cloud Composer environment
	RecoveryConfigs []Composer_getEnvironmentConfigRecoveryConfig `json:"recoveryConfigs,omitempty" yaml:"recoveryConfigs,omitempty"`

	// The configuration setting for Airflow data retention mechanism. This field is supported for Cloud Composer environments in versions composer-2.0.32-airflow-2.1.4. or newer
	DataRetentionConfigs []Composer_getEnvironmentConfigDataRetentionConfig `json:"dataRetentionConfigs,omitempty" yaml:"dataRetentionConfigs,omitempty"`

	// The size of the Cloud Composer environment. This field is supported for Cloud Composer environments in versions composer-2.-.--airflow--.-.- and newer.
	EnvironmentSize string `json:"environmentSize,omitempty" yaml:"environmentSize,omitempty"`

	// The configuration for Cloud Composer maintenance window.
	MaintenanceWindows []Composer_getEnvironmentConfigMaintenanceWindow `json:"maintenanceWindows,omitempty" yaml:"maintenanceWindows,omitempty"`

	// Network-level access control policy for the Airflow web server.
	WebServerNetworkAccessControls []Composer_getEnvironmentConfigWebServerNetworkAccessControl `json:"webServerNetworkAccessControls,omitempty" yaml:"webServerNetworkAccessControls,omitempty"`

	// The workloads configuration settings for the GKE cluster associated with the Cloud Composer environment. Supported for Cloud Composer environments in versions composer-2.-.--airflow--.-.- and newer.
	WorkloadsConfigs []Composer_getEnvironmentConfigWorkloadsConfig `json:"workloadsConfigs,omitempty" yaml:"workloadsConfigs,omitempty"`

	// The configuration of Cloud SQL instance that is used by the Apache Airflow software. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	DatabaseConfigs []Composer_getEnvironmentConfigDatabaseConfig `json:"databaseConfigs,omitempty" yaml:"databaseConfigs,omitempty"`

	// Optional. If true, builds performed during operations that install Python packages have only private connectivity to Google services. If false, the builds also have access to the internet.
	EnablePrivateBuildsOnly bool `json:"enablePrivateBuildsOnly,omitempty" yaml:"enablePrivateBuildsOnly,omitempty"`

	// The encryption options for the Composer environment and its dependencies.
	EncryptionConfigs []Composer_getEnvironmentConfigEncryptionConfig `json:"encryptionConfigs,omitempty" yaml:"encryptionConfigs,omitempty"`

	// The number of nodes in the Kubernetes Engine cluster that will be used to run this environment. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`

	// The configuration used for the Private IP Cloud Composer environment.
	PrivateEnvironmentConfigs []Composer_getEnvironmentConfigPrivateEnvironmentConfig `json:"privateEnvironmentConfigs,omitempty" yaml:"privateEnvironmentConfigs,omitempty"`

	// The configuration settings for the Airflow web server App Engine instance. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	WebServerConfigs []Composer_getEnvironmentConfigWebServerConfig `json:"webServerConfigs,omitempty" yaml:"webServerConfigs,omitempty"`

	// The configuration settings for software inside the environment.
	SoftwareConfigs []Composer_getEnvironmentConfigSoftwareConfig `json:"softwareConfigs,omitempty" yaml:"softwareConfigs,omitempty"`

	// The Cloud Storage prefix of the DAGs for this environment. Although Cloud Storage objects reside in a flat namespace, a hierarchical file tree can be simulated using '/'-delimited object name prefixes. DAG objects for this environment reside in a simulated directory with this prefix.
	DagGcsPrefix string `json:"dagGcsPrefix,omitempty" yaml:"dagGcsPrefix,omitempty"`

	// Optional. If true, a private Composer environment will be created.
	EnablePrivateEnvironment bool `json:"enablePrivateEnvironment,omitempty" yaml:"enablePrivateEnvironment,omitempty"`

	// The Kubernetes Engine cluster used to run this environment.
	GkeCluster string `json:"gkeCluster,omitempty" yaml:"gkeCluster,omitempty"`

	// Configuration options for the master authorized networks feature. Enabled master authorized networks will disallow all external traffic to access Kubernetes master through HTTPS except traffic from the given CIDR blocks, Google Compute Engine Public IPs and Google Prod IPs.
	MasterAuthorizedNetworksConfigs []Composer_getEnvironmentConfigMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfigs,omitempty" yaml:"masterAuthorizedNetworksConfigs,omitempty"`

	// The configuration used for the Kubernetes Engine cluster.
	NodeConfigs []Composer_getEnvironmentConfigNodeConfig `json:"nodeConfigs,omitempty" yaml:"nodeConfigs,omitempty"`

	// Whether high resilience is enabled or not. This field is supported for Cloud Composer environments in versions composer-2.1.15-airflow--.-.- and newer.
	ResilienceMode string `json:"resilienceMode,omitempty" yaml:"resilienceMode,omitempty"`
}
