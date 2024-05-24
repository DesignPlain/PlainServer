package appsync

import types "DesignSphere_Server/src/resource/aws/types"

type DataSource struct {
	// API ID for the GraphQL API for the data source.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// DynamoDB settings. See DynamoDB Config
	DynamodbConfig types.Appsync_DataSourceDynamodbConfig `json:"dynamodbConfig,omitempty" yaml:"dynamodbConfig,omitempty"`

	// AWS EventBridge settings. See Event Bridge Config
	EventBridgeConfig types.Appsync_DataSourceEventBridgeConfig `json:"eventBridgeConfig,omitempty" yaml:"eventBridgeConfig,omitempty"`

	// HTTP settings. See HTTP Config
	HttpConfig types.Appsync_DataSourceHttpConfig `json:"httpConfig,omitempty" yaml:"httpConfig,omitempty"`

	// Description of the data source.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Amazon Elasticsearch settings. See ElasticSearch Config
	ElasticsearchConfig types.Appsync_DataSourceElasticsearchConfig `json:"elasticsearchConfig,omitempty" yaml:"elasticsearchConfig,omitempty"`

	// AWS Lambda settings. See Lambda Config
	LambdaConfig types.Appsync_DataSourceLambdaConfig `json:"lambdaConfig,omitempty" yaml:"lambdaConfig,omitempty"`

	// User-supplied name for the data source.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Amazon OpenSearch Service settings. See OpenSearch Service Config
	OpensearchserviceConfig types.Appsync_DataSourceOpensearchserviceConfig `json:"opensearchserviceConfig,omitempty" yaml:"opensearchserviceConfig,omitempty"`

	// AWS RDS settings. See Relational Database Config
	RelationalDatabaseConfig types.Appsync_DataSourceRelationalDatabaseConfig `json:"relationalDatabaseConfig,omitempty" yaml:"relationalDatabaseConfig,omitempty"`

	// IAM service role ARN for the data source. Required if `type` is specified as `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `AMAZON_EVENTBRIDGE`, or `AMAZON_OPENSEARCH_SERVICE`.
	ServiceRoleArn string `json:"serviceRoleArn,omitempty" yaml:"serviceRoleArn,omitempty"`

	// Type of the Data Source. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`, `RELATIONAL_DATABASE`, `AMAZON_EVENTBRIDGE`, `AMAZON_OPENSEARCH_SERVICE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
