package chatbot

import types "libds/aws/types"

type TeamsChannelConfiguration struct {
	// ID of the Microsoft Teams channel.
	ChannelId string `json:"channelId,omitempty" yaml:"channelId,omitempty"`

	// ID of the Microsoft Team authorized with AWS Chatbot. To get the team ID, you must perform the initial authorization flow with Microsoft Teams in the AWS Chatbot console. Then you can copy and paste the team ID from the console.
	TeamId string `json:"teamId,omitempty" yaml:"teamId,omitempty"`

	/*
	   ID of the Microsoft Teams tenant.

	   The following arguments are optional:
	*/
	TenantId string `json:"tenantId,omitempty" yaml:"tenantId,omitempty"`

	//
	Timeouts types.Chatbot_TeamsChannelConfigurationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Logging levels include `ERROR`, `INFO`, or `NONE`.
	LoggingLevel string `json:"loggingLevel,omitempty" yaml:"loggingLevel,omitempty"`

	// ARNs of the SNS topics that deliver notifications to AWS Chatbot.
	SnsTopicArns []string `json:"snsTopicArns,omitempty" yaml:"snsTopicArns,omitempty"`

	// Map of tags assigned to the resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Name of the Microsoft Teams team.
	TeamName string `json:"teamName,omitempty" yaml:"teamName,omitempty"`

	// Name of the Microsoft Teams channel.
	ChannelName string `json:"channelName,omitempty" yaml:"channelName,omitempty"`

	// Name of the Microsoft Teams channel configuration.
	ConfigurationName string `json:"configurationName,omitempty" yaml:"configurationName,omitempty"`

	// List of IAM policy ARNs that are applied as channel guardrails. The AWS managed `AdministratorAccess` policy is applied by default if this is not set.
	GuardrailPolicyArns []string `json:"guardrailPolicyArns,omitempty" yaml:"guardrailPolicyArns,omitempty"`

	// ARN of the IAM role that defines the permissions for AWS Chatbot. This is a user-defined role that AWS Chatbot will assume. This is not the service-linked role.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// Enables use of a user role requirement in your chat configuration.
	UserAuthorizationRequired bool `json:"userAuthorizationRequired,omitempty" yaml:"userAuthorizationRequired,omitempty"`
}
