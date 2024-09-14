package types

type Bcmdata_ExportExport struct {
	// Description for this specific data export.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Destination configuration for this specific data export. See the `destination_configurations` argument reference below.
	DestinationConfigurations []Bcmdata_ExportExportDestinationConfiguration `json:"destinationConfigurations,omitempty" yaml:"destinationConfigurations,omitempty"`

	// Amazon Resource Name (ARN) for this export.
	ExportArn string `json:"exportArn,omitempty" yaml:"exportArn,omitempty"`

	// Name of this specific data export.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Cadence for Amazon Web Services to update the export in your S3 bucket. See the `refresh_cadence` argument reference below.
	RefreshCadences []Bcmdata_ExportExportRefreshCadence `json:"refreshCadences,omitempty" yaml:"refreshCadences,omitempty"`

	// Data query for this specific data export. See the `data_query` argument reference below.
	DataQueries []Bcmdata_ExportExportDataQuery `json:"dataQueries,omitempty" yaml:"dataQueries,omitempty"`
}
