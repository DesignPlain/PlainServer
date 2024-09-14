package types

type Bcmdata_ExportExportRefreshCadence struct {
	// Frequency that data exports are updated. The export refreshes each time the source data updates, up to three times daily. Valid values `SYNCHRONOUS`.
	Frequency string `json:"frequency,omitempty" yaml:"frequency,omitempty"`
}
