package types

type Fsx_DataRepositoryAssociationS3AutoImportPolicy struct {
	// A list of file event types to automatically export to your linked S3 bucket or import from the linked S3 bucket. Valid values are `NEW`, `CHANGED`, `DELETED`. Max of 3.
	Events []string `json:"events,omitempty" yaml:"events,omitempty"`
}
