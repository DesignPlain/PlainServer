package rds

type InstanceAutomatedBackupsReplication struct {
	// The AWS KMS key identifier for encryption of the replicated automated backups. The KMS key ID is the Amazon Resource Name (ARN) for the KMS encryption key in the destination AWS Region, for example, `arn:aws:kms:us-east-1:123456789012:key/AKIAIOSFODNN7EXAMPLE`.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// A URL that contains a [Signature Version 4](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html) signed request for the [`StartDBInstanceAutomatedBackupsReplication`](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_StartDBInstanceAutomatedBackupsReplication.html) action to be called in the AWS Region of the source DB instance.
	PreSignedUrl string `json:"preSignedUrl,omitempty" yaml:"preSignedUrl,omitempty"`

	// The retention period for the replicated automated backups, defaults to `7`.
	RetentionPeriod int `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`

	// The Amazon Resource Name (ARN) of the source DB instance for the replicated automated backups, for example, `arn:aws:rds:us-west-2:123456789012:db:mydatabase`.
	SourceDbInstanceArn string `json:"sourceDbInstanceArn,omitempty" yaml:"sourceDbInstanceArn,omitempty"`
}
