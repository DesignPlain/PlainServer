package chatbot

import types "libds/aws/types"

type SlackChannelConfiguration struct {
	// List of IAM policy ARNs that are applied as channel guardrails. The AWS managed `AdministratorAccess` policy is applied by default if this is not set.
	GuardrailPolicyArns []string `json:"guardrailPolicyArns,omitempty" yaml:"guardrailPolicyArns,omitempty"`

	// Logging levels include `ERROR`, `INFO`, or `NONE`.
	LoggingLevel string `json:"loggingLevel,omitempty" yaml:"loggingLevel,omitempty"`

	/*
	   ID of the Slack workspace authorized with AWS Chatbot. For example, `T07EA123LEP`.

	   The following arguments are optional:
	*/
	SlackTeamId string `json:"slackTeamId,omitempty" yaml:"slackTeamId,omitempty"`

	// ARNs of the SNS topics that deliver notifications to AWS Chatbot.
	SnsTopicArns []string `json:"snsTopicArns,omitempty" yaml:"snsTopicArns,omitempty"`

	// Map of tags assigned to the resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Chatbot_SlackChannelConfigurationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Enables use of a user role requirement in your chat configuration.
	UserAuthorizationRequired bool `json:"userAuthorizationRequired,omitempty" yaml:"userAuthorizationRequired,omitempty"`

	// Name of the Slack channel configuration.
	ConfigurationName string `json:"configurationName,omitempty" yaml:"configurationName,omitempty"`

	// User-defined role that AWS Chatbot assumes. This is not the service-linked role.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// ID of the Slack channel. For example, `C07EZ1ABC23`.
	SlackChannelId string `json:"slackChannelId,omitempty" yaml:"slackChannelId,omitempty"`
}
