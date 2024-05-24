package types

type Opensearch_DomainClusterConfig struct {
	// Whether dedicated main nodes are enabled for the cluster.
	DedicatedMasterEnabled bool `json:"dedicatedMasterEnabled,omitempty" yaml:"dedicatedMasterEnabled,omitempty"`

	// Number of instances in the cluster.
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`

	// Instance type of data nodes in the cluster.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Number of warm nodes in the cluster. Valid values are between `2` and `150`. `warm_count` can be only and must be set when `warm_enabled` is set to `true`.
	WarmCount int `json:"warmCount,omitempty" yaml:"warmCount,omitempty"`

	// Configuration block containing zone awareness settings. Detailed below.
	ZoneAwarenessConfig Opensearch_DomainClusterConfigZoneAwarenessConfig `json:"zoneAwarenessConfig,omitempty" yaml:"zoneAwarenessConfig,omitempty"`

	// Configuration block containing cold storage configuration. Detailed below.
	ColdStorageOptions Opensearch_DomainClusterConfigColdStorageOptions `json:"coldStorageOptions,omitempty" yaml:"coldStorageOptions,omitempty"`

	// Number of dedicated main nodes in the cluster.
	DedicatedMasterCount int `json:"dedicatedMasterCount,omitempty" yaml:"dedicatedMasterCount,omitempty"`

	// Instance type of the dedicated main nodes in the cluster.
	DedicatedMasterType string `json:"dedicatedMasterType,omitempty" yaml:"dedicatedMasterType,omitempty"`

	// Whether a multi-AZ domain is turned on with a standby AZ. For more information, see [Configuring a multi-AZ domain in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html).
	MultiAzWithStandbyEnabled bool `json:"multiAzWithStandbyEnabled,omitempty" yaml:"multiAzWithStandbyEnabled,omitempty"`

	// Whether to enable warm storage.
	WarmEnabled bool `json:"warmEnabled,omitempty" yaml:"warmEnabled,omitempty"`

	// Instance type for the OpenSearch cluster's warm nodes. Valid values are `ultrawarm1.medium.search`, `ultrawarm1.large.search` and `ultrawarm1.xlarge.search`. `warm_type` can be only and must be set when `warm_enabled` is set to `true`.
	WarmType string `json:"warmType,omitempty" yaml:"warmType,omitempty"`

	// Whether zone awareness is enabled, set to `true` for multi-az deployment. To enable awareness with three Availability Zones, the `availability_zone_count` within the `zone_awareness_config` must be set to `3`.
	ZoneAwarenessEnabled bool `json:"zoneAwarenessEnabled,omitempty" yaml:"zoneAwarenessEnabled,omitempty"`
}
