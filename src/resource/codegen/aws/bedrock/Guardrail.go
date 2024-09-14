package bedrock

import types "libds/aws/types"

type Guardrail struct {
	// Message to return when the guardrail blocks a prompt.
	BlockedInputMessaging string `json:"blockedInputMessaging,omitempty" yaml:"blockedInputMessaging,omitempty"`

	// Description of the guardrail or its version.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	Timeouts types.Bedrock_GuardrailTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Topic policy config for a guardrail. See Topic Policy Config for more information.
	TopicPolicyConfig types.Bedrock_GuardrailTopicPolicyConfig `json:"topicPolicyConfig,omitempty" yaml:"topicPolicyConfig,omitempty"`

	// Message to return when the guardrail blocks a model response.
	BlockedOutputsMessaging string `json:"blockedOutputsMessaging,omitempty" yaml:"blockedOutputsMessaging,omitempty"`

	// Content policy config for a guardrail. See Content Policy Config for more information.
	ContentPolicyConfig types.Bedrock_GuardrailContentPolicyConfig `json:"contentPolicyConfig,omitempty" yaml:"contentPolicyConfig,omitempty"`

	// Contextual grounding policy config for a guardrail. See Contextual Grounding Policy Config for more information.
	ContextualGroundingPolicyConfig types.Bedrock_GuardrailContextualGroundingPolicyConfig `json:"contextualGroundingPolicyConfig,omitempty" yaml:"contextualGroundingPolicyConfig,omitempty"`

	// The KMS key with which the guardrail was encrypted at rest.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	/*
	   Name of the guardrail.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Sensitive information policy config for a guardrail. See Sensitive Information Policy Config for more information.
	SensitiveInformationPolicyConfig types.Bedrock_GuardrailSensitiveInformationPolicyConfig `json:"sensitiveInformationPolicyConfig,omitempty" yaml:"sensitiveInformationPolicyConfig,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Word policy config for a guardrail. See Word Policy Config for more information.
	WordPolicyConfig types.Bedrock_GuardrailWordPolicyConfig `json:"wordPolicyConfig,omitempty" yaml:"wordPolicyConfig,omitempty"`
}
