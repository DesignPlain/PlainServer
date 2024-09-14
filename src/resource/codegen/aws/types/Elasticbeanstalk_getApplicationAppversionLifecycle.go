package types

type Elasticbeanstalk_getApplicationAppversionLifecycle struct {
	// Specifies whether delete a version's source bundle from S3 when the application version is deleted.
	DeleteSourceFromS3 bool `json:"deleteSourceFromS3,omitempty" yaml:"deleteSourceFromS3,omitempty"`

	// Number of days to retain an application version.
	MaxAgeInDays int `json:"maxAgeInDays,omitempty" yaml:"maxAgeInDays,omitempty"`

	// Maximum number of application versions to retain.
	MaxCount int `json:"maxCount,omitempty" yaml:"maxCount,omitempty"`

	// ARN of an IAM service role under which the application version is deleted.  Elastic Beanstalk must have permission to assume this role.
	ServiceRole string `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`
}
