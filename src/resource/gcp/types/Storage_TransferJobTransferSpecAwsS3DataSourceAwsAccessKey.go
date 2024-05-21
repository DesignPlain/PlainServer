package types

type Storage_TransferJobTransferSpecAwsS3DataSourceAwsAccessKey struct {
	// AWS Secret Access Key.
	SecretAccessKey string `json:"secretAccessKey,omitempty" yaml:"secretAccessKey,omitempty"`

	// AWS Key ID.
	AccessKeyId string `json:"accessKeyId,omitempty" yaml:"accessKeyId,omitempty"`
}
