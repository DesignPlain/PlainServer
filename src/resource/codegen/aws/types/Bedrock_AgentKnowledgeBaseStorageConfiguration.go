package types

type Bedrock_AgentKnowledgeBaseStorageConfiguration struct {
	// The storage configuration of the knowledge base in Amazon OpenSearch Service. See `opensearch_serverless_configuration` block for details.
	OpensearchServerlessConfiguration Bedrock_AgentKnowledgeBaseStorageConfigurationOpensearchServerlessConfiguration `json:"opensearchServerlessConfiguration,omitempty" yaml:"opensearchServerlessConfiguration,omitempty"`

	// The storage configuration of the knowledge base in Pinecone. See `pinecone_configuration` block for details.
	PineconeConfiguration Bedrock_AgentKnowledgeBaseStorageConfigurationPineconeConfiguration `json:"pineconeConfiguration,omitempty" yaml:"pineconeConfiguration,omitempty"`

	// Details about the storage configuration of the knowledge base in Amazon RDS. For more information, see [Create a vector index in Amazon RDS](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup.html). See `rds_configuration` block for details.
	RdsConfiguration Bedrock_AgentKnowledgeBaseStorageConfigurationRdsConfiguration `json:"rdsConfiguration,omitempty" yaml:"rdsConfiguration,omitempty"`

	// The storage configuration of the knowledge base in Redis Enterprise Cloud. See `redis_enterprise_cloud_configuration` block for details.
	RedisEnterpriseCloudConfiguration Bedrock_AgentKnowledgeBaseStorageConfigurationRedisEnterpriseCloudConfiguration `json:"redisEnterpriseCloudConfiguration,omitempty" yaml:"redisEnterpriseCloudConfiguration,omitempty"`

	// Vector store service in which the knowledge base is stored. Valid Values: `OPENSEARCH_SERVERLESS`, `PINECONE`, `REDIS_ENTERPRISE_CLOUD`, `RDS`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
