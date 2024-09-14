package types

type Timestreamwrite_TableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation struct {
	// Configuration of an S3 location to write error reports for records rejected, asynchronously, during magnetic store writes. See S3 Configuration below for more details.
	S3Configuration Timestreamwrite_TableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`
}
