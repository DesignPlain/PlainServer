package types

type Cognito_UserPoolSmsConfiguration struct {
	// External ID used in IAM role trust relationships. For more information about using external IDs, see [How to Use an External ID When Granting Access to Your AWS Resources to a Third Party](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html).
	ExternalId string `json:"externalId,omitempty" yaml:"externalId,omitempty"`

	// ARN of the Amazon SNS caller. This is usually the IAM role that you've given Cognito permission to assume.
	SnsCallerArn string `json:"snsCallerArn,omitempty" yaml:"snsCallerArn,omitempty"`

	// The AWS Region to use with Amazon SNS integration. You can choose the same Region as your user pool, or a supported Legacy Amazon SNS alternate Region. Amazon Cognito resources in the Asia Pacific (Seoul) AWS Region must use your Amazon SNS configuration in the Asia Pacific (Tokyo) Region. For more information, see [SMS message settings for Amazon Cognito user pools](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html).
	SnsRegion string `json:"snsRegion,omitempty" yaml:"snsRegion,omitempty"`
}
