package types

type Glue_CatalogTableOpenTableFormatInputIcebergInput struct {
	// A required metadata operation. Can only be set to CREATE.
	MetadataOperation string `json:"metadataOperation,omitempty" yaml:"metadataOperation,omitempty"`

	// The table version for the Iceberg table. Defaults to 2.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
