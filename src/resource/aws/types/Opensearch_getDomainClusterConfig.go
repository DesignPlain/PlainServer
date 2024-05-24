package types

type Opensearch_getDomainClusterConfig struct {
	// Indicates whether dedicated master nodes are enabled for the cluster.
	DedicatedMasterEnabled bool `json:"dedicatedMasterEnabled,omitempty" yaml:"dedicatedMasterEnabled,omitempty"`

	// Warm storage is enabled.
	WarmEnabled bool `json:"warmEnabled,omitempty" yaml:"warmEnabled,omitempty"`

	// Indicates whether zone awareness is enabled.
	ZoneAwarenessEnabled bool `json:"zoneAwarenessEnabled,omitempty" yaml:"zoneAwarenessEnabled,omitempty"`

	// Configuration block containing cold storage configuration.
	ColdStorageOptions []Opensearch_getDomainClusterConfigColdStorageOption `json:"coldStorageOptions,omitempty" yaml:"coldStorageOptions,omitempty"`

	// Number of dedicated master nodes in the cluster.
	DedicatedMasterCount int `json:"dedicatedMasterCount,omitempty" yaml:"dedicatedMasterCount,omitempty"`

	// Instance type of data nodes in the cluster.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Whether a multi-AZ domain is turned on with a standby AZ.
	MultiAzWithStandbyEnabled bool `json:"multiAzWithStandbyEnabled,omitempty" yaml:"multiAzWithStandbyEnabled,omitempty"`

	// Number of warm nodes in the cluster.
	WarmCount int `json:"warmCount,omitempty" yaml:"warmCount,omitempty"`

	// Instance type for the OpenSearch cluster's warm nodes.
	WarmType string `json:"warmType,omitempty" yaml:"warmType,omitempty"`

	// Configuration block containing zone awareness settings.
	ZoneAwarenessConfigs []Opensearch_getDomainClusterConfigZoneAwarenessConfig `json:"zoneAwarenessConfigs,omitempty" yaml:"zoneAwarenessConfigs,omitempty"`

	// Instance type of the dedicated master nodes in the cluster.
	DedicatedMasterType string `json:"dedicatedMasterType,omitempty" yaml:"dedicatedMasterType,omitempty"`

	// Number of instances in the cluster.
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`
}
