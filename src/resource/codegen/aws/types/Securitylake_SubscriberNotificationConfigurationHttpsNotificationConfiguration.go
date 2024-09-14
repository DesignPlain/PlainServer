package types

type Securitylake_SubscriberNotificationConfigurationHttpsNotificationConfiguration struct {
	// The API key name for the notification subscription.
	AuthorizationApiKeyName string `json:"authorizationApiKeyName,omitempty" yaml:"authorizationApiKeyName,omitempty"`

	// The API key value for the notification subscription.
	AuthorizationApiKeyValue string `json:"authorizationApiKeyValue,omitempty" yaml:"authorizationApiKeyValue,omitempty"`

	/*
	   The subscription endpoint in Security Lake.
	   If you prefer notification with an HTTPS endpoint, populate this field.
	*/
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	/*
	   The HTTP method used for the notification subscription.
	   Valid values are `POST` and `PUT`.
	*/
	HttpMethod string `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`

	/*
	   The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you created.
	   For more information about ARNs and how to use them in policies, see Managing data access and AWS Managed Policies in the Amazon Security Lake User Guide.
	*/
	TargetRoleArn string `json:"targetRoleArn,omitempty" yaml:"targetRoleArn,omitempty"`
}
