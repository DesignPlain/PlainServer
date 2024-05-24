package docdb

type ClusterInstance struct {
	// The identifier of the `aws.docdb.Cluster` in which to launch this instance.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	/*
	   The instance class to use. For details on CPU and memory, see [Scaling for DocumentDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance).
	   DocumentDB currently supports the below instance classes.
	   Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	   - db.r6g.large
	   - db.r6g.xlarge
	   - db.r6g.2xlarge
	   - db.r6g.4xlarge
	   - db.r6g.8xlarge
	   - db.r6g.12xlarge
	   - db.r6g.16xlarge
	   - db.r5.large
	   - db.r5.xlarge
	   - db.r5.2xlarge
	   - db.r5.4xlarge
	   - db.r5.12xlarge
	   - db.r5.24xlarge
	   - db.r4.large
	   - db.r4.xlarge
	   - db.r4.2xlarge
	   - db.r4.4xlarge
	   - db.r4.8xlarge
	   - db.r4.16xlarge
	   - db.t4g.medium
	   - db.t3.medium
	*/
	InstanceClass string `json:"instanceClass,omitempty" yaml:"instanceClass,omitempty"`

	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	// Copy all DB instance `tags` to snapshots. Default is `false`.
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" yaml:"copyTagsToSnapshot,omitempty"`

	/*
	   The window to perform maintenance in.
	   Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	*/
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier int `json:"promotionTier,omitempty" yaml:"promotionTier,omitempty"`

	// A map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Specifies whether any database modifications
	   are applied immediately, or during the next maintenance window. Default is`false`.
	*/
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The identifier of the certificate authority (CA) certificate for the DB instance.
	CaCertIdentifier string `json:"caCertIdentifier,omitempty" yaml:"caCertIdentifier,omitempty"`

	// The name of the database engine to be used for the DocumentDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix string `json:"identifierPrefix,omitempty" yaml:"identifierPrefix,omitempty"`

	// A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
	EnablePerformanceInsights bool `json:"enablePerformanceInsights,omitempty" yaml:"enablePerformanceInsights,omitempty"`

	// The identifier for the DocumentDB instance, if omitted, the provider will assign a random, unique identifier.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
	PerformanceInsightsKmsKeyId string `json:"performanceInsightsKmsKeyId,omitempty" yaml:"performanceInsightsKmsKeyId,omitempty"`
}
