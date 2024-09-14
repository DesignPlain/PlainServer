package m2

import types "libds/aws/types"

type Environment struct {
	//
	ApplyChangesDuringMaintenanceWindow bool `json:"applyChangesDuringMaintenanceWindow,omitempty" yaml:"applyChangesDuringMaintenanceWindow,omitempty"`

	//
	StorageConfiguration types.M2_EnvironmentStorageConfiguration `json:"storageConfiguration,omitempty" yaml:"storageConfiguration,omitempty"`

	// The specific version of the engine for the Environment.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	//
	HighAvailabilityConfig types.M2_EnvironmentHighAvailabilityConfig `json:"highAvailabilityConfig,omitempty" yaml:"highAvailabilityConfig,omitempty"`

	// ARN of the KMS key to use for the Environment.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// List of subnet ids to deploy environment to.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Key-value tags for the place index. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Force update the environment even if applications are running.
	ForceUpdate bool `json:"forceUpdate,omitempty" yaml:"forceUpdate,omitempty"`

	// Name of the runtime environment. Must be unique within the account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configures the maintenance window that you want for the runtime environment. The maintenance window must have the format `ddd:hh24:mi-ddd:hh24:mi` and must be less than 24 hours. If not provided a random value will be used.
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// Allow applications deployed to this environment to be publicly accessible.
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`

	// List of security group ids.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	//
	Timeouts types.M2_EnvironmentTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Engine type must be `microfocus` or `bluage`.
	EngineType string `json:"engineType,omitempty" yaml:"engineType,omitempty"`

	/*
	   M2 Instance Type.

	   The following arguments are optional:
	*/
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
}
