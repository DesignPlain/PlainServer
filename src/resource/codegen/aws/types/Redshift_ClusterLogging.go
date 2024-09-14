package types

type Redshift_ClusterLogging struct {
	// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	// The log destination type. An enum with possible values of `s3` and `cloudwatch`.
	LogDestinationType string `json:"logDestinationType,omitempty" yaml:"logDestinationType,omitempty"`

	// The collection of exported log types. Log types include the connection log, user log and user activity log. Required when `log_destination_type` is `cloudwatch`. Valid log types are `connectionlog`, `userlog`, and `useractivitylog`.
	LogExports []string `json:"logExports,omitempty" yaml:"logExports,omitempty"`

	// The prefix applied to the log file names.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`

	/*
	   The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
	   For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
	*/
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
