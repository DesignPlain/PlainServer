package cognito

import types "DesignSphere_Server/src/resource/aws/types"

type UserPool struct {
	// Configuration block to define which verified available method a user can use to recover their forgotten password. Detailed below.
	AccountRecoverySetting types.Cognito_UserPoolAccountRecoverySetting `json:"accountRecoverySetting,omitempty" yaml:"accountRecoverySetting,omitempty"`

	// Configuration block for configuring email. Detailed below.
	EmailConfiguration types.Cognito_UserPoolEmailConfiguration `json:"emailConfiguration,omitempty" yaml:"emailConfiguration,omitempty"`

	// String representing the SMS verification message. Conflicts with `verification_message_template` configuration block `sms_message` argument.
	SmsVerificationMessage string `json:"smsVerificationMessage,omitempty" yaml:"smsVerificationMessage,omitempty"`

	// Configuration block for user attribute update settings. Detailed below.
	UserAttributeUpdateSettings types.Cognito_UserPoolUserAttributeUpdateSettings `json:"userAttributeUpdateSettings,omitempty" yaml:"userAttributeUpdateSettings,omitempty"`

	// Configuration block for creating a new user profile. Detailed below.
	AdminCreateUserConfig types.Cognito_UserPoolAdminCreateUserConfig `json:"adminCreateUserConfig,omitempty" yaml:"adminCreateUserConfig,omitempty"`

	// Attributes to be auto-verified. Valid values: `email`, `phone_number`.
	AutoVerifiedAttributes []string `json:"autoVerifiedAttributes,omitempty" yaml:"autoVerifiedAttributes,omitempty"`

	/*
	   Name of the user pool.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block for verification message templates. Detailed below.
	VerificationMessageTemplate types.Cognito_UserPoolVerificationMessageTemplate `json:"verificationMessageTemplate,omitempty" yaml:"verificationMessageTemplate,omitempty"`

	// Configuration block for the schema attributes of a user pool. Detailed below. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Attributes can be added, but not modified or removed. Maximum of 50 attributes.
	Schemas []types.Cognito_UserPoolSchema `json:"schemas,omitempty" yaml:"schemas,omitempty"`

	// Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the `taint` command.
	SmsConfiguration types.Cognito_UserPoolSmsConfiguration `json:"smsConfiguration,omitempty" yaml:"smsConfiguration,omitempty"`

	// Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
	SoftwareTokenMfaConfiguration types.Cognito_UserPoolSoftwareTokenMfaConfiguration `json:"softwareTokenMfaConfiguration,omitempty" yaml:"softwareTokenMfaConfiguration,omitempty"`

	// Map of tags to assign to the User Pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Attributes supported as an alias for this user pool. Valid values: `phone_number`, `email`, or `preferred_username`. Conflicts with `username_attributes`.
	AliasAttributes []string `json:"aliasAttributes,omitempty" yaml:"aliasAttributes,omitempty"`

	// When active, DeletionProtection prevents accidental deletion of your user pool. Before you can delete a user pool that you have protected against deletion, you must deactivate this feature. Valid values are `ACTIVE` and `INACTIVE`, Default value is `INACTIVE`.
	DeletionProtection string `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// String representing the email verification subject. Conflicts with `verification_message_template` configuration block `email_subject` argument.
	EmailVerificationSubject string `json:"emailVerificationSubject,omitempty" yaml:"emailVerificationSubject,omitempty"`

	// Configuration block for the AWS Lambda triggers associated with the user pool. Detailed below.
	LambdaConfig types.Cognito_UserPoolLambdaConfig `json:"lambdaConfig,omitempty" yaml:"lambdaConfig,omitempty"`

	// Whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `alias_attributes`.
	UsernameAttributes []string `json:"usernameAttributes,omitempty" yaml:"usernameAttributes,omitempty"`

	// String representing the SMS authentication message. The Message must contain the `{####}` placeholder, which will be replaced with the code.
	SmsAuthenticationMessage string `json:"smsAuthenticationMessage,omitempty" yaml:"smsAuthenticationMessage,omitempty"`

	// Configuration block for user pool add-ons to enable user pool advanced security mode features. Detailed below.
	UserPoolAddOns types.Cognito_UserPoolUserPoolAddOns `json:"userPoolAddOns,omitempty" yaml:"userPoolAddOns,omitempty"`

	// Configuration block for username configuration. Detailed below.
	UsernameConfiguration types.Cognito_UserPoolUsernameConfiguration `json:"usernameConfiguration,omitempty" yaml:"usernameConfiguration,omitempty"`

	// Configuration block for the user pool's device tracking. Detailed below.
	DeviceConfiguration types.Cognito_UserPoolDeviceConfiguration `json:"deviceConfiguration,omitempty" yaml:"deviceConfiguration,omitempty"`

	// String representing the email verification message. Conflicts with `verification_message_template` configuration block `email_message` argument.
	EmailVerificationMessage string `json:"emailVerificationMessage,omitempty" yaml:"emailVerificationMessage,omitempty"`

	// Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values are `OFF` (MFA Tokens are not required), `ON` (MFA is required for all users to sign in; requires at least one of `sms_configuration` or `software_token_mfa_configuration` to be configured), or `OPTIONAL` (MFA Will be required only for individual users who have MFA Enabled; requires at least one of `sms_configuration` or `software_token_mfa_configuration` to be configured).
	MfaConfiguration string `json:"mfaConfiguration,omitempty" yaml:"mfaConfiguration,omitempty"`

	// Configuration block for information about the user pool password policy. Detailed below.
	PasswordPolicy types.Cognito_UserPoolPasswordPolicy `json:"passwordPolicy,omitempty" yaml:"passwordPolicy,omitempty"`
}
