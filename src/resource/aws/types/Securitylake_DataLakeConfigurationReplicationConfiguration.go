package types

type Securitylake_DataLakeConfigurationReplicationConfiguration struct {
	// Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. Amazon S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different AWS Regions or within the same Region as the source bucket.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`
}
