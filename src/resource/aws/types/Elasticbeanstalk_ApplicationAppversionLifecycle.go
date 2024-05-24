package types

type Elasticbeanstalk_ApplicationAppversionLifecycle struct {
	// Set to `true` to delete a version's source bundle from S3 when the application version is deleted.
	DeleteSourceFromS3 bool `json:"deleteSourceFromS3,omitempty" yaml:"deleteSourceFromS3,omitempty"`

	// The number of days to retain an application version ('max_age_in_days' and 'max_count' cannot be enabled simultaneously.).
	MaxAgeInDays int `json:"maxAgeInDays,omitempty" yaml:"maxAgeInDays,omitempty"`

	// The maximum number of application versions to retain ('max_age_in_days' and 'max_count' cannot be enabled simultaneously.).
	MaxCount int `json:"maxCount,omitempty" yaml:"maxCount,omitempty"`

	// The ARN of an IAM service role under which the application version is deleted.  Elastic Beanstalk must have permission to assume this role.
	ServiceRole string `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`
}
