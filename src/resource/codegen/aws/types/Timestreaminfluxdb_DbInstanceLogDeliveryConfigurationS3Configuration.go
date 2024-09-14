package types

type Timestreaminfluxdb_DbInstanceLogDeliveryConfigurationS3Configuration struct {
	// Name of the S3 bucket to deliver logs to.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	/*
	   Indicates whether log delivery to the S3 bucket is enabled.

	   --Note--: Only three arguments do updates in-place: `db_parameter_group_identifier`, `log_delivery_configuration`, and `tags`. Changes to any other argument after a DB instance has been deployed will cause destruction and re-creation of the DB instance. Additionally, when `db_parameter_group_identifier` is added to a DB instance or modified, the DB instance will be updated in-place but if `db_parameter_group_identifier` is removed from a DB instance, the DB instance will be destroyed and re-created.
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
