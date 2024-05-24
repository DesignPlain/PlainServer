package guardduty

type PublishingDestination struct {
	// The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
	DestinationArn string `json:"destinationArn,omitempty" yaml:"destinationArn,omitempty"`

	/*
	   Currently there is only "S3" available as destination type which is also the default value

	   > --Note:-- In case of missing permissions (S3 Bucket Policy _or_ KMS Key permissions) the resource will fail to create. If the permissions are changed after resource creation, this can be asked from the AWS API via the "DescribePublishingDestination" call (https://docs.aws.amazon.com/cli/latest/reference/guardduty/describe-publishing-destination.html).
	*/
	DestinationType string `json:"destinationType,omitempty" yaml:"destinationType,omitempty"`

	// The detector ID of the GuardDuty.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`

	// The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
