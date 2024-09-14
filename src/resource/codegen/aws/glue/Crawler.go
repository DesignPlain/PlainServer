package glue

import types "libds/aws/types"

type Crawler struct {
	// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
	Classifiers []string `json:"classifiers,omitempty" yaml:"classifiers,omitempty"`

	// Description of the crawler.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of nested Iceberg target arguments. See Iceberg Target below.
	IcebergTargets []types.Glue_CrawlerIcebergTarget `json:"icebergTargets,omitempty" yaml:"icebergTargets,omitempty"`

	// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.. See Recrawl Policy below.
	RecrawlPolicy types.Glue_CrawlerRecrawlPolicy `json:"recrawlPolicy,omitempty" yaml:"recrawlPolicy,omitempty"`

	// The name of Security Configuration to be used by the crawler
	SecurityConfiguration string `json:"securityConfiguration,omitempty" yaml:"securityConfiguration,omitempty"`

	// The table prefix used for catalog tables that are created.
	TablePrefix string `json:"tablePrefix,omitempty" yaml:"tablePrefix,omitempty"`

	// List of nested AWS Glue Data Catalog target arguments. See Catalog Target below.
	CatalogTargets []types.Glue_CrawlerCatalogTarget `json:"catalogTargets,omitempty" yaml:"catalogTargets,omitempty"`

	// List of nested Delta Lake target arguments. See Delta Target below.
	DeltaTargets []types.Glue_CrawlerDeltaTarget `json:"deltaTargets,omitempty" yaml:"deltaTargets,omitempty"`

	// List of nested DynamoDB target arguments. See Dynamodb Target below.
	DynamodbTargets []types.Glue_CrawlerDynamodbTarget `json:"dynamodbTargets,omitempty" yaml:"dynamodbTargets,omitempty"`

	// List of nested MongoDB target arguments. See MongoDB Target below.
	MongodbTargets []types.Glue_CrawlerMongodbTarget `json:"mongodbTargets,omitempty" yaml:"mongodbTargets,omitempty"`

	// List of nested Amazon S3 target arguments. See S3 Target below.
	S3Targets []types.Glue_CrawlerS3Target `json:"s3Targets,omitempty" yaml:"s3Targets,omitempty"`

	// List of nested Hudi target arguments. See Iceberg Target below.
	HudiTargets []types.Glue_CrawlerHudiTarget `json:"hudiTargets,omitempty" yaml:"hudiTargets,omitempty"`

	// List of nested JDBC target arguments. See JDBC Target below.
	JdbcTargets []types.Glue_CrawlerJdbcTarget `json:"jdbcTargets,omitempty" yaml:"jdbcTargets,omitempty"`

	// Specifies Lake Formation configuration settings for the crawler. See Lake Formation Configuration below.
	LakeFormationConfiguration types.Glue_CrawlerLakeFormationConfiguration `json:"lakeFormationConfiguration,omitempty" yaml:"lakeFormationConfiguration,omitempty"`

	// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	// A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 - - ? -)`.
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Policy for the crawler's update and deletion behavior. See Schema Change Policy below.
	SchemaChangePolicy types.Glue_CrawlerSchemaChangePolicy `json:"schemaChangePolicy,omitempty" yaml:"schemaChangePolicy,omitempty"`

	// JSON string of configuration information. For more details see [Setting Crawler Configuration Options](https://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html).
	Configuration string `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// Glue database where results are written.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Specifies data lineage configuration settings for the crawler. See Lineage Configuration below.
	LineageConfiguration types.Glue_CrawlerLineageConfiguration `json:"lineageConfiguration,omitempty" yaml:"lineageConfiguration,omitempty"`

	// Name of the crawler.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
