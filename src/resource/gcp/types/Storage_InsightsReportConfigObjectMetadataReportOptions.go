package types

type Storage_InsightsReportConfigObjectMetadataReportOptions struct {
	// The metadata fields included in an inventory report.
	MetadataFields []string `json:"metadataFields,omitempty" yaml:"metadataFields,omitempty"`

	/*
	   Options for where the inventory reports are stored.
	   Structure is documented below.
	*/
	StorageDestinationOptions Storage_InsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions `json:"storageDestinationOptions,omitempty" yaml:"storageDestinationOptions,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	StorageFilters Storage_InsightsReportConfigObjectMetadataReportOptionsStorageFilters `json:"storageFilters,omitempty" yaml:"storageFilters,omitempty"`
}
