package cognito

type User struct {
	// A map that contains user attributes and attribute values to be set for the user.
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	// A list of mediums to the welcome message will be sent through. Allowed values are `EMAIL` and `SMS`. If it's provided, make sure you have also specified `email` attribute for the `EMAIL` medium and `phone_number` for the `SMS`. More than one value can be specified. Amazon Cognito does not store the `desired_delivery_mediums` value. Defaults to `["SMS"]`.
	DesiredDeliveryMediums []string `json:"desiredDeliveryMediums,omitempty" yaml:"desiredDeliveryMediums,omitempty"`

	// The user's temporary password. Conflicts with `password`.
	TemporaryPassword string `json:"temporaryPassword,omitempty" yaml:"temporaryPassword,omitempty"`

	/*
	   The username for the user. Must be unique within the user pool. Must be a UTF-8 string between 1 and 128 characters. After the user is created, the username cannot be changed.

	   The following arguments are optional:
	*/
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the `client_metadata` value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ClientMetadata map[string]string `json:"clientMetadata,omitempty" yaml:"clientMetadata,omitempty"`

	// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the `enabled` value. The behavior can be changed with `message_action` argument. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// If this parameter is set to True and the `phone_number` or `email` address specified in the `attributes` parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the `force_alias_creation` value. Defaults to `false`.
	ForceAliasCreation bool `json:"forceAliasCreation,omitempty" yaml:"forceAliasCreation,omitempty"`

	// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to `SUPPRESS` to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the `message_action` value.
	MessageAction string `json:"messageAction,omitempty" yaml:"messageAction,omitempty"`

	// The user's permanent password. This password must conform to the password policy specified by user pool the user belongs to. The welcome message always contains only `temporary_password` value. You can suppress sending the welcome message with the `message_action` argument. Amazon Cognito does not store the `password` value. Conflicts with `temporary_password`.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// The user pool ID for the user pool where the user will be created.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`

	/*
	   The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the `validation_data` value. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).

	   > --NOTE:-- Clearing `password` or `temporary_password` does not reset user's password in Cognito.
	*/
	ValidationData map[string]string `json:"validationData,omitempty" yaml:"validationData,omitempty"`
}
