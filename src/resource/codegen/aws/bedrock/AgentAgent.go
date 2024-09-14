package bedrock

import types "libds/aws/types"

type AgentAgent struct {
	// Configurations to override prompt templates in different parts of an agent sequence. For more information, see [Advanced prompts](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html). See `prompt_override_configuration` Block for details.
	PromptOverrideConfigurations []types.Bedrock_AgentAgentPromptOverrideConfiguration `json:"promptOverrideConfigurations,omitempty" yaml:"promptOverrideConfigurations,omitempty"`

	//
	Timeouts types.Bedrock_AgentAgentTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Name of the agent.
	AgentName string `json:"agentName,omitempty" yaml:"agentName,omitempty"`

	// ARN of the AWS KMS key that encrypts the agent.
	CustomerEncryptionKeyArn string `json:"customerEncryptionKeyArn,omitempty" yaml:"customerEncryptionKeyArn,omitempty"`

	// Whether to prepare the agent after creation or modification. Defaults to `true`.
	PrepareAgent bool `json:"prepareAgent,omitempty" yaml:"prepareAgent,omitempty"`

	// Number of seconds for which Amazon Bedrock keeps information about a user's conversation with the agent. A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Bedrock deletes any data provided before the timeout.
	IdleSessionTtlInSeconds int `json:"idleSessionTtlInSeconds,omitempty" yaml:"idleSessionTtlInSeconds,omitempty"`

	// Instructions that tell the agent what it should do and how it should interact with users.
	Instruction string `json:"instruction,omitempty" yaml:"instruction,omitempty"`

	// Whether the in-use check is skipped when deleting the agent.
	SkipResourceInUseCheck bool `json:"skipResourceInUseCheck,omitempty" yaml:"skipResourceInUseCheck,omitempty"`

	// Map of tags assigned to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// ARN of the IAM role with permissions to invoke API operations on the agent.
	AgentResourceRoleArn string `json:"agentResourceRoleArn,omitempty" yaml:"agentResourceRoleArn,omitempty"`

	// Description of the agent.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Foundation model used for orchestration by the agent.

	   The following arguments are optional:
	*/
	FoundationModel string `json:"foundationModel,omitempty" yaml:"foundationModel,omitempty"`
}
