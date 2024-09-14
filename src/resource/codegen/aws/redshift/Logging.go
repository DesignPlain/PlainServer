package redshift

type Logging struct {
	// Prefix applied to the log file names.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`

	// Name of an existing S3 bucket where the log files are to be stored. Required when `log_destination_type` is `s3`. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions. For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	/*
	   Identifier of the source cluster.

	   The following arguments are optional:
	*/
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// Log destination type. Valid values are `s3` and `cloudwatch`.
	LogDestinationType string `json:"logDestinationType,omitempty" yaml:"logDestinationType,omitempty"`

	// Collection of exported log types. Required when `log_destination_type` is `cloudwatch`. Valid values are `connectionlog`, `useractivitylog`, and `userlog`.
	LogExports []string `json:"logExports,omitempty" yaml:"logExports,omitempty"`
}
