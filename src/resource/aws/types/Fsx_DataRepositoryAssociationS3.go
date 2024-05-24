package types

type Fsx_DataRepositoryAssociationS3 struct {
	// Specifies the type of updated objects that will be automatically exported from your file system to the linked S3 bucket. See the `events` configuration block.
	AutoExportPolicy Fsx_DataRepositoryAssociationS3AutoExportPolicy `json:"autoExportPolicy,omitempty" yaml:"autoExportPolicy,omitempty"`

	// Specifies the type of updated objects that will be automatically imported from the linked S3 bucket to your file system. See the `events` configuration block.
	AutoImportPolicy Fsx_DataRepositoryAssociationS3AutoImportPolicy `json:"autoImportPolicy,omitempty" yaml:"autoImportPolicy,omitempty"`
}
