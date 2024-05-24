package types

type S3_AnalyticsConfigurationStorageClassAnalysisDataExport struct {
	// Specifies the destination for the exported analytics data (documented below).
	Destination S3_AnalyticsConfigurationStorageClassAnalysisDataExportDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Schema version of exported analytics data. Allowed values: `V_1`. Default value: `V_1`.
	OutputSchemaVersion string `json:"outputSchemaVersion,omitempty" yaml:"outputSchemaVersion,omitempty"`
}
