package bedrock

import types "libds/aws/types"

type AgentKnowledgeBase struct {
	//
	Timeouts types.Bedrock_AgentKnowledgeBaseTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Description of the knowledge base.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Details about the embeddings configuration of the knowledge base. See `knowledge_base_configuration` block for details.
	KnowledgeBaseConfiguration types.Bedrock_AgentKnowledgeBaseKnowledgeBaseConfiguration `json:"knowledgeBaseConfiguration,omitempty" yaml:"knowledgeBaseConfiguration,omitempty"`

	// Name of the knowledge base.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ARN of the IAM role with permissions to invoke API operations on the knowledge base.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	/*
	   Details about the storage configuration of the knowledge base. See `storage_configuration` block for details.

	   The following arguments are optional:
	*/
	StorageConfiguration types.Bedrock_AgentKnowledgeBaseStorageConfiguration `json:"storageConfiguration,omitempty" yaml:"storageConfiguration,omitempty"`

	// Map of tags assigned to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
