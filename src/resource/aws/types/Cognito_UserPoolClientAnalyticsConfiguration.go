package types

type Cognito_UserPoolClientAnalyticsConfiguration struct {
	// If set to `true`, Amazon Cognito will include user data in the events it publishes to Amazon Pinpoint analytics.
	UserDataShared bool `json:"userDataShared,omitempty" yaml:"userDataShared,omitempty"`

	// Application ARN for an Amazon Pinpoint application. Conflicts with `external_id` and `role_arn`.
	ApplicationArn string `json:"applicationArn,omitempty" yaml:"applicationArn,omitempty"`

	// Application ID for an Amazon Pinpoint application.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// ID for the Analytics Configuration. Conflicts with `application_arn`.
	ExternalId string `json:"externalId,omitempty" yaml:"externalId,omitempty"`

	// ARN of an IAM role that authorizes Amazon Cognito to publish events to Amazon Pinpoint analytics. Conflicts with `application_arn`.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
