package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type Workteam struct {
	// The name of the Workteam (must be unique).
	WorkforceName string `json:"workforceName,omitempty" yaml:"workforceName,omitempty"`

	// The name of the workforce.
	WorkteamName string `json:"workteamName,omitempty" yaml:"workteamName,omitempty"`

	// A description of the work team.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognito_member_definition`. For workforces created using your own OIDC identity provider (IdP) use `oidc_member_definition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	MemberDefinitions []types.Sagemaker_WorkteamMemberDefinition `json:"memberDefinitions,omitempty" yaml:"memberDefinitions,omitempty"`

	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	NotificationConfiguration types.Sagemaker_WorkteamNotificationConfiguration `json:"notificationConfiguration,omitempty" yaml:"notificationConfiguration,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
