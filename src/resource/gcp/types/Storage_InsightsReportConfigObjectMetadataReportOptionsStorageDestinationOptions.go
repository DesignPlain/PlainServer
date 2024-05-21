package types

type Storage_InsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions struct {
	// The destination bucket that stores the generated inventory reports.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The path within the destination bucket to store generated inventory reports.
	DestinationPath string `json:"destinationPath,omitempty" yaml:"destinationPath,omitempty"`
}
