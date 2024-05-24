package types

type S3_InventoryDestinationBucket struct {
	// Prefix that is prepended to all inventory results.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// ID of the account that owns the destination bucket. Recommended to be set to prevent problems if the destination bucket ownership changes.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Amazon S3 bucket ARN of the destination.
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`

	// Contains the type of server-side encryption to use to encrypt the inventory (documented below).
	Encryption S3_InventoryDestinationBucketEncryption `json:"encryption,omitempty" yaml:"encryption,omitempty"`

	// Specifies the output format of the inventory results. Can be `CSV`, [`ORC`](https://orc.apache.org/) or [`Parquet`](https://parquet.apache.org/).
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}
