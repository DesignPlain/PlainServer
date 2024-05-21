package datafusion

import types "DesignSphere_Server/src/resource/gcp/types"

type Instance struct {
	/*
	   The resource labels for instance to use to annotate any related underlying resources,
	   such as Compute Engine VMs.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The region of the Data Fusion instance.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   List of accelerators enabled for this CDF instance.
	   If accelerators are enabled it is possible a permadiff will be created with the Options field.
	   Users will need to either manually update their state file to include these diffed options, or include the field in a lifecycle ignore changes block.
	   Structure is documented below.
	*/
	Accelerators []types.Datafusion_InstanceAccelerator `json:"accelerators,omitempty" yaml:"accelerators,omitempty"`

	/*
	   The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
	   Structure is documented below.
	*/
	CryptoKeyConfig types.Datafusion_InstanceCryptoKeyConfig `json:"cryptoKeyConfig,omitempty" yaml:"cryptoKeyConfig,omitempty"`

	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	DataprocServiceAccount string `json:"dataprocServiceAccount,omitempty" yaml:"dataprocServiceAccount,omitempty"`

	// An optional description of the instance.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging bool `json:"enableStackdriverLogging,omitempty" yaml:"enableStackdriverLogging,omitempty"`

	/*
	   Option to enable and pass metadata for event publishing.
	   Structure is documented below.
	*/
	EventPublishConfig types.Datafusion_InstanceEventPublishConfig `json:"eventPublishConfig,omitempty" yaml:"eventPublishConfig,omitempty"`

	// Current version of the Data Fusion.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	/*
	   Specifies whether the Data Fusion instance should be private. If set to
	   true, all Data Fusion nodes will have private IP addresses and will not be
	   able to access the public internet.
	*/
	PrivateInstance bool `json:"privateInstance,omitempty" yaml:"privateInstance,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Represents the type of Data Fusion instance. Each type is configured with
	   the default settings for processing and memory.
	   - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
	   using point and click UI. However, there are certain limitations, such as fewer number
	   of concurrent pipelines, no support for streaming pipelines, etc.
	   - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
	   available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	   - DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
	   with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration
	   pipelines at low cost.
	   Possible values are: `BASIC`, `ENTERPRISE`, `DEVELOPER`.


	   - - -
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Display name for an instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Option to enable granular role-based access control.
	EnableRbac bool `json:"enableRbac,omitempty" yaml:"enableRbac,omitempty"`

	/*
	   Network configuration options. These are required when a private Data Fusion instance is to be created.
	   Structure is documented below.
	*/
	NetworkConfig types.Datafusion_InstanceNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring bool `json:"enableStackdriverMonitoring,omitempty" yaml:"enableStackdriverMonitoring,omitempty"`

	// The ID of the instance or a fully qualified identifier for the instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options map[string]string `json:"options,omitempty" yaml:"options,omitempty"`

	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
