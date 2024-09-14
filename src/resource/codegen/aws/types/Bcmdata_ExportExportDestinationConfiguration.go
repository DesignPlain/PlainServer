package types

type Bcmdata_ExportExportDestinationConfiguration struct {
	// Object that describes the destination of the data exports file. See the `s3_destination` argument reference below.
	S3Destinations []Bcmdata_ExportExportDestinationConfigurationS3Destination `json:"s3Destinations,omitempty" yaml:"s3Destinations,omitempty"`
}
