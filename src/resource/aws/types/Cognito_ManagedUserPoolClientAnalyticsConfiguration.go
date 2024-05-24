package types

type Cognito_ManagedUserPoolClientAnalyticsConfiguration struct {
	// Unique identifier for an Amazon Pinpoint application.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// ID for the Analytics Configuration and conflicts with `application_arn`.
	ExternalId string `json:"externalId,omitempty" yaml:"externalId,omitempty"`

	// ARN of an IAM role that authorizes Amazon Cognito to publish events to Amazon Pinpoint analytics. It conflicts with `application_arn`.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// If `user_data_shared` is set to `true`, Amazon Cognito will include user data in the events it publishes to Amazon Pinpoint analytics.
	UserDataShared bool `json:"userDataShared,omitempty" yaml:"userDataShared,omitempty"`

	// Application ARN for an Amazon Pinpoint application. It conflicts with `external_id` and `role_arn`.
	ApplicationArn string `json:"applicationArn,omitempty" yaml:"applicationArn,omitempty"`
}
