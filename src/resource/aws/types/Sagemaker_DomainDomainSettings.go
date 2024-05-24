package types

type Sagemaker_DomainDomainSettings struct {
	// The security groups for the Amazon Virtual Private Cloud that the Domain uses for communication between Domain-level apps and user apps.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The configuration for attaching a SageMaker user profile name to the execution role as a sts:SourceIdentity key [AWS Docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_monitor.html). Valid values are `USER_PROFILE_NAME` and `DISABLED`.
	ExecutionRoleIdentityConfig string `json:"executionRoleIdentityConfig,omitempty" yaml:"executionRoleIdentityConfig,omitempty"`

	// A collection of settings that configure the RStudioServerPro Domain-level app. see `r_studio_server_pro_domain_settings` Block below.
	RStudioServerProDomainSettings Sagemaker_DomainDomainSettingsRStudioServerProDomainSettings `json:"rStudioServerProDomainSettings,omitempty" yaml:"rStudioServerProDomainSettings,omitempty"`
}
