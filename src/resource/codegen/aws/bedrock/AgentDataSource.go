package bedrock

import types "libds/aws/types"

type AgentDataSource struct {
	//
	Timeouts types.Bedrock_AgentDataSourceTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Details about the configuration of the server-side encryption. See `vector_ingestion_configuration` block for details.
	VectorIngestionConfiguration types.Bedrock_AgentDataSourceVectorIngestionConfiguration `json:"vectorIngestionConfiguration,omitempty" yaml:"vectorIngestionConfiguration,omitempty"`

	// Data deletion policy for a data source. Valid values: `RETAIN`, `DELETE`.
	DataDeletionPolicy string `json:"dataDeletionPolicy,omitempty" yaml:"dataDeletionPolicy,omitempty"`

	// Details about how the data source is stored. See `data_source_configuration` block for details.
	DataSourceConfiguration types.Bedrock_AgentDataSourceDataSourceConfiguration `json:"dataSourceConfiguration,omitempty" yaml:"dataSourceConfiguration,omitempty"`

	// Description of the data source.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Unique identifier of the knowledge base to which the data source belongs.
	KnowledgeBaseId string `json:"knowledgeBaseId,omitempty" yaml:"knowledgeBaseId,omitempty"`

	/*
	   Name of the data source.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Details about the configuration of the server-side encryption. See `server_side_encryption_configuration` block for details.
	ServerSideEncryptionConfiguration types.Bedrock_AgentDataSourceServerSideEncryptionConfiguration `json:"serverSideEncryptionConfiguration,omitempty" yaml:"serverSideEncryptionConfiguration,omitempty"`
}
