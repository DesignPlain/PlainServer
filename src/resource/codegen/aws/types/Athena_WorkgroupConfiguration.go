package types

type Athena_WorkgroupConfiguration struct {
	// Role used in a notebook session for accessing the user's resources.
	ExecutionRole string `json:"executionRole,omitempty" yaml:"executionRole,omitempty"`

	// Boolean whether Amazon CloudWatch metrics are enabled for the workgroup. Defaults to `true`.
	PublishCloudwatchMetricsEnabled bool `json:"publishCloudwatchMetricsEnabled,omitempty" yaml:"publishCloudwatchMetricsEnabled,omitempty"`

	// If set to true , allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false , workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. The default is false . For more information about Requester Pays buckets, see [Requester Pays Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html) in the Amazon Simple Storage Service Developer Guide.
	RequesterPaysEnabled bool `json:"requesterPaysEnabled,omitempty" yaml:"requesterPaysEnabled,omitempty"`

	// Configuration block with result settings. See Result Configuration below.
	ResultConfiguration Athena_WorkgroupConfigurationResultConfiguration `json:"resultConfiguration,omitempty" yaml:"resultConfiguration,omitempty"`

	// Integer for the upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan. Must be at least `10485760`.
	BytesScannedCutoffPerQuery int `json:"bytesScannedCutoffPerQuery,omitempty" yaml:"bytesScannedCutoffPerQuery,omitempty"`

	// Boolean whether the settings for the workgroup override client-side settings. For more information, see [Workgroup Settings Override Client-Side Settings](https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html). Defaults to `true`.
	EnforceWorkgroupConfiguration bool `json:"enforceWorkgroupConfiguration,omitempty" yaml:"enforceWorkgroupConfiguration,omitempty"`

	// Configuration block for the Athena Engine Versioning. For more information, see [Athena Engine Versioning](https://docs.aws.amazon.com/athena/latest/ug/engine-versions.html). See Engine Version below.
	EngineVersion Athena_WorkgroupConfigurationEngineVersion `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`
}
