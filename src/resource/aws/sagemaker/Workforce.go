package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type Workforce struct {
	// Use this parameter to configure an Amazon Cognito private workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. Conflicts with `oidc_config`. see Cognito Config details below.
	CognitoConfig types.Sagemaker_WorkforceCognitoConfig `json:"cognitoConfig,omitempty" yaml:"cognitoConfig,omitempty"`

	// Use this parameter to configure a private workforce using your own OIDC Identity Provider. Conflicts with `cognito_config`. see OIDC Config details below.
	OidcConfig types.Sagemaker_WorkforceOidcConfig `json:"oidcConfig,omitempty" yaml:"oidcConfig,omitempty"`

	// A list of IP address ranges Used to create an allow list of IP addresses for a private workforce. By default, a workforce isn't restricted to specific IP addresses. see Source Ip Config details below.
	SourceIpConfig types.Sagemaker_WorkforceSourceIpConfig `json:"sourceIpConfig,omitempty" yaml:"sourceIpConfig,omitempty"`

	// The name of the Workforce (must be unique).
	WorkforceName string `json:"workforceName,omitempty" yaml:"workforceName,omitempty"`

	// configure a workforce using VPC. see Workforce VPC Config details below.
	WorkforceVpcConfig types.Sagemaker_WorkforceWorkforceVpcConfig `json:"workforceVpcConfig,omitempty" yaml:"workforceVpcConfig,omitempty"`
}
