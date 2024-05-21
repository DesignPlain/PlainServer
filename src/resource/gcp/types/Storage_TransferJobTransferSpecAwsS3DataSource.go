package types

type Storage_TransferJobTransferSpecAwsS3DataSource struct {
	// AWS credentials block.
	AwsAccessKey Storage_TransferJobTransferSpecAwsS3DataSourceAwsAccessKey `json:"awsAccessKey,omitempty" yaml:"awsAccessKey,omitempty"`

	// Google Cloud Storage bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The Amazon Resource Name (ARN) of the role to support temporary credentials via 'AssumeRoleWithWebIdentity'. For more information about ARNs, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns). When a role ARN is provided, Transfer Service fetches temporary credentials for the session using a 'AssumeRoleWithWebIdentity' call for the provided role using the [GoogleServiceAccount][] for this project.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
