package types

type Timestreamwrite_getTableMagneticStoreWritePropertyMagneticStoreRejectedDataLocation struct {
	// Object containing the following attributes to describe the configuration of an s3 location to write error reports for records rejected.
	S3Configurations []Timestreamwrite_getTableMagneticStoreWritePropertyMagneticStoreRejectedDataLocationS3Configuration `json:"s3Configurations,omitempty" yaml:"s3Configurations,omitempty"`
}
