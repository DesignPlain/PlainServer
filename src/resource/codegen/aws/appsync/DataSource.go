package appsync

import types "libds/aws/types"

type DataSource struct {
	// Amazon Elasticsearch settings. See `elasticsearch_config` Block for details.
	ElasticsearchConfig types.Appsync_DataSourceElasticsearchConfig `json:"elasticsearchConfig,omitempty" yaml:"elasticsearchConfig,omitempty"`

	// Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// DynamoDB settings. See `dynamodb_config` Block for details.
	DynamodbConfig types.Appsync_DataSourceDynamodbConfig `json:"dynamodbConfig,omitempty" yaml:"dynamodbConfig,omitempty"`

	// Description of the data source.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// AWS EventBridge settings. See `event_bridge_config` Block for details.
	EventBridgeConfig types.Appsync_DataSourceEventBridgeConfig `json:"eventBridgeConfig,omitempty" yaml:"eventBridgeConfig,omitempty"`

	// HTTP settings. See `http_config` Block for details.
	HttpConfig types.Appsync_DataSourceHttpConfig `json:"httpConfig,omitempty" yaml:"httpConfig,omitempty"`

	// AWS Lambda settings. See `lambda_config` Block for details.
	LambdaConfig types.Appsync_DataSourceLambdaConfig `json:"lambdaConfig,omitempty" yaml:"lambdaConfig,omitempty"`

	// User-supplied name for the data source.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Amazon OpenSearch Service settings. See `opensearchservice_config` Block for details.
	OpensearchserviceConfig types.Appsync_DataSourceOpensearchserviceConfig `json:"opensearchserviceConfig,omitempty" yaml:"opensearchserviceConfig,omitempty"`

	// AWS RDS settings. See `relational_database_config` Block for details.
	RelationalDatabaseConfig types.Appsync_DataSourceRelationalDatabaseConfig `json:"relationalDatabaseConfig,omitempty" yaml:"relationalDatabaseConfig,omitempty"`

	// API ID for the GraphQL API for the data source.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
	ServiceRoleArn string `json:"serviceRoleArn,omitempty" yaml:"serviceRoleArn,omitempty"`
}
